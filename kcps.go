//
// Copyright 2016, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gokcps

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
)

// UnlimitedResourceID is a special ID to define an unlimited resource
const UnlimitedResourceID = "-1"

var idRegex = regexp.MustCompile(`^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|-1)$`)

// For KDDI
// var requestMux = &sync.Mutex{}

// IsID return true if the passed ID is either a UUID or a UnlimitedResourceID
func IsID(id string) bool {
	return idRegex.MatchString(id)
}

// OptionFunc can be passed to the courtesy helper functions to set additional parameters
type OptionFunc func(*KCPSClient, interface{}) error

type CSError struct {
	ErrorCode   int    `json:"errorcode"`
	CSErrorCode int    `json:"cserrorcode"`
	ErrorText   string `json:"errortext"`
}

func (e *CSError) Error() error {
	return fmt.Errorf("CloudStack API error %d (CSExceptionErrorCode: %d): %s", e.ErrorCode, e.CSErrorCode, e.ErrorText)
}

type KCPSClient struct {
	HTTPGETOnly bool // If `true` only use HTTP GET calls

	client  *http.Client // The http client for communicating
	baseURL string       // The base URL of the API
	apiKey  string       // Api key
	secret  string       // Secret key
	async   bool         // Wait for async calls to finish
	timeout int64        // Max waiting timeout in seconds for async jobs to finish; defaults to 300 seconds

	Asyncjob       *AsyncjobService
	Event          *EventService
	Firewall       *FirewallService
	GuestOS        *GuestOSService
	Host           *HostService
	ISO            *ISOService
	LoadBalancer   *LoadBalancerService
	NatPortForward *NatPortForwardService
	Nic            *NicService
	Snapshot       *SnapshotService
	Template       *TemplateService
	AccountDomain  *AccountDomainService
	VirtualMachine *VirtualMachineService
	Volume         *VolumeService
	Tags           *TagsService
}

// Creates a new client for communicating with CloudStack
func newClient(apiurl string, apikey string, secret string, async bool, verifyssl bool) *KCPSClient {
	cs := &KCPSClient{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !verifyssl}, // If verifyssl is true, skipping the verify should be false and vice versa
			},
			Timeout: time.Duration(60 * time.Second),
		},
		baseURL: apiurl,
		apiKey:  apikey,
		secret:  secret,
		async:   async,
		timeout: 300,
	}
	cs.Asyncjob = NewAsyncjobService(cs)
	cs.Event = NewEventService(cs)
	cs.Firewall = NewFirewallService(cs)
	cs.GuestOS = NewGuestOSService(cs)
	cs.Host = NewHostService(cs)
	cs.ISO = NewISOService(cs)
	cs.LoadBalancer = NewLoadBalancerService(cs)
	cs.NatPortForward = NewNatPortForwardService(cs)
	cs.Nic = NewNicService(cs)
	cs.Snapshot = NewSnapshotService(cs)
	cs.Template = NewTemplateService(cs)
	cs.AccountDomain = NewAccountDomainService(cs)
	cs.VirtualMachine = NewVirtualMachineService(cs)
	cs.Volume = NewVolumeService(cs)
	return cs
}

// Default non-async client. So for async calls you need to implement and check the async job result yourself. When using
// HTTPS with a self-signed certificate to connect to your CloudStack API, you would probably want to set 'verifyssl' to
// false so the call ignores the SSL errors/warnings.
func NewClient(apiurl string, apikey string, secret string, verifyssl bool) *KCPSClient {
	cs := newClient(apiurl, apikey, secret, false, verifyssl)
	return cs
}

// For sync API calls this client behaves exactly the same as a standard client call, but for async API calls
// this client will wait until the async job is finished or until the configured AsyncTimeout is reached. When the async
// job finishes successfully it will return actual object received from the API and nil, but when the timout is
// reached it will return the initial object containing the async job ID for the running job and a warning.
func NewAsyncClient(apiurl string, apikey string, secret string, verifyssl bool) *KCPSClient {
	cs := newClient(apiurl, apikey, secret, true, verifyssl)
	return cs
}

// When using the async client an api call will wait for the async call to finish before returning. The default is to poll for 300 seconds
// seconds, to check if the async job is finished.
func (cs *KCPSClient) AsyncTimeout(timeoutInSeconds int64) {
	cs.timeout = timeoutInSeconds
}

var AsyncTimeoutErr = errors.New("Timeout while waiting for async job to finish")

// A helper function that you can use to get the result of a running async job. If the job is not finished within the configured
// timeout, the async job returns a AsyncTimeoutErr.
func (cs *KCPSClient) GetAsyncJobResult(jobid string, timeout int64) (json.RawMessage, error) {
	var timer time.Duration
	currentTime := time.Now().Unix()

	for {
		p := cs.Asyncjob.NewQueryAsyncJobResultParams(jobid)
		r, err := cs.Asyncjob.QueryAsyncJobResult(p)
		if err != nil {
			return nil, err
		}

		// Status 1 means the job is finished successfully
		if r.Jobstatus == 1 {
			return r.Jobresult, nil
		}

		// When the status is 2, the job has failed
		if r.Jobstatus == 2 {
			if r.Jobresulttype == "text" {
				return nil, fmt.Errorf(string(r.Jobresult))
			} else {
				return nil, fmt.Errorf("Undefined error: %s", string(r.Jobresult))
			}
		}

		if time.Now().Unix()-currentTime > timeout {
			return nil, AsyncTimeoutErr
		}

		// Add an (extremely simple) exponential backoff like feature to prevent
		// flooding the CloudStack API
		if timer < 15 {
			timer++
		}

		time.Sleep(timer * time.Second)
	}
}

// A helper function that you can use to get the result of a running async job. If the job is not finished within the configured
// timeout, the async job returns a AsyncTimeoutErr.
func (cs *KCPSClient) GetExAsyncJobResult(jobid string, timeout int64) (json.RawMessage, error) {
	var timer time.Duration
	currentTime := time.Now().Unix()

	for {
		p := cs.Asyncjob.NewQueryExAsyncJobResultParams(jobid)
		r, err := cs.Asyncjob.QueryExAsyncJobResult(p)
		if err != nil {
			return nil, err
		}

		// Status 1 means the job is finished successfully
		if r.Jobstatus == 1 {
			return r.Jobresult, nil
		}

		// When the status is 2, the job has failed
		if r.Jobstatus == 2 {
			if r.Jobresulttype == "text" {
				return nil, fmt.Errorf(string(r.Jobresult))
			} else {
				return nil, fmt.Errorf("Undefined error: %s", string(r.Jobresult))
			}
		}

		if time.Now().Unix()-currentTime > timeout {
			return nil, AsyncTimeoutErr
		}

		// Add an (extremely simple) exponential backoff like feature to prevent
		// flooding the CloudStack API
		if timer < 15 {
			timer++
		}

		time.Sleep(timer * time.Second)
	}
}

// Execute the request against a CS API. Will return the raw JSON data returned by the API and nil if
// no error occured. If the API returns an error the result will be nil and the HTTP error code and CS
// error details. If a processing (code) error occurs the result will be nil and the generated error
func (cs *KCPSClient) newRequest(api string, params url.Values) (json.RawMessage, error) {

	maxret := 5
	retry := 0
	var (
		err     error
		message json.RawMessage
	)

	for {
		retry++
		if maxret < retry {
			err = fmt.Errorf("api request max retry")
			break
		}

		m, e := cs.oneRequest(api, params)
		if e != nil && strings.HasSuffix(e.Error(), "connection reset by peer") {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			continue
		}
		message = m
		err = e
		break
	}

	return message, err
}

func (cs *KCPSClient) oneRequest(api string, params url.Values) (json.RawMessage, error) {
	params.Set("apiKey", cs.apiKey)
	params.Set("command", api)
	params.Set("response", "json")

	// Generate signature for API call
	// * Serialize parameters, URL encoding only values and sort them by key, done by encodeValues
	// * Convert the entire argument string to lowercase
	// * Replace all instances of '+' to '%20'
	// * Calculate HMAC SHA1 of argument string with CloudStack secret
	// * URL encode the string and convert to base64
	s := encodeValues(params)
	s2 := strings.ToLower(s)
	s3 := strings.Replace(s2, "+", "%20", -1)
	mac := hmac.New(sha1.New, []byte(cs.secret))
	mac.Write([]byte(s3))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	var err error
	var resp *http.Response
	if !cs.HTTPGETOnly && (api == "deployVirtualMachine" || api == "updateVirtualMachine") {
		// The deployVirtualMachine API should be called using a POST call
		// so we don't have to worry about the userdata size

		// Add the unescaped signature to the POST params
		params.Set("signature", signature)

		// Make a POST call
		resp, err = cs.client.PostForm(cs.baseURL, params)
	} else {
		// Create the final URL before we issue the request
		url := cs.baseURL + "?" + s + "&signature=" + url.QueryEscape(signature)

		// Make a GET call
		//		requestMux.Lock()
		resp, err = cs.client.Get(url)
		// For KDDI
		//  time.Sleep(2 * time.Second)
		//	requestMux.Unlock()
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Need to get the raw value to make the result play nice
	b, err = getRawValue(b)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var e CSError
		if err := json.Unmarshal(b, &e); err != nil {
			return nil, err
		}
		return nil, e.Error()
	}
	return b, nil
}

// Custom version of net/url Encode that only URL escapes values
// Unmodified portions here remain under BSD license of The Go Authors: https://go.googlesource.com/go/+/master/LICENSE
func encodeValues(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := k + "="
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}

// Generic function to get the first raw value from a response as json.RawMessage
func getRawValue(b json.RawMessage) (json.RawMessage, error) {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	for _, v := range m {
		return v, nil
	}
	return nil, fmt.Errorf("Unable to extract the raw value from:\n\n%s\n\n", string(b))
}

// VPCIDSetter is an interface that every type that can set a vpc ID must implement
type VPCIDSetter interface {
	SetVpcid(string)
}

// WithVPCID takes a vpc ID and sets the `vpcid` parameter
func WithVPCID(id string) OptionFunc {
	return func(cs *KCPSClient, p interface{}) error {
		vs, ok := p.(VPCIDSetter)

		if !ok || id == "" {
			return nil
		}

		vs.SetVpcid(id)

		return nil
	}
}

type AsyncjobService struct {
	cs *KCPSClient
}

func NewAsyncjobService(cs *KCPSClient) *AsyncjobService {
	return &AsyncjobService{cs: cs}
}

type EventService struct {
	cs *KCPSClient
}

func NewEventService(cs *KCPSClient) *EventService {
	return &EventService{cs: cs}
}

type FirewallService struct {
	cs *KCPSClient
}

func NewFirewallService(cs *KCPSClient) *FirewallService {
	return &FirewallService{cs: cs}
}

type GuestOSService struct {
	cs *KCPSClient
}

func NewGuestOSService(cs *KCPSClient) *GuestOSService {
	return &GuestOSService{cs: cs}
}

type HostService struct {
	cs *KCPSClient
}

func NewHostService(cs *KCPSClient) *HostService {
	return &HostService{cs: cs}
}

type ISOService struct {
	cs *KCPSClient
}

func NewISOService(cs *KCPSClient) *ISOService {
	return &ISOService{cs: cs}
}

type LoadBalancerService struct {
	cs *KCPSClient
}

func NewLoadBalancerService(cs *KCPSClient) *LoadBalancerService {
	return &LoadBalancerService{cs: cs}
}

type NatPortForwardService struct {
	cs *KCPSClient
}

func NewNatPortForwardService(cs *KCPSClient) *NatPortForwardService {
	return &NatPortForwardService{cs: cs}
}

type NicService struct {
	cs *KCPSClient
}

func NewNicService(cs *KCPSClient) *NicService {
	return &NicService{cs: cs}
}

type SnapshotService struct {
	cs *KCPSClient
}

func NewSnapshotService(cs *KCPSClient) *SnapshotService {
	return &SnapshotService{cs: cs}
}

type TemplateService struct {
	cs *KCPSClient
}

func NewTemplateService(cs *KCPSClient) *TemplateService {
	return &TemplateService{cs: cs}
}

type AccountDomainService struct {
	cs *KCPSClient
}

func NewAccountDomainService(cs *KCPSClient) *AccountDomainService {
	return &AccountDomainService{cs: cs}
}

type VirtualMachineService struct {
	cs *KCPSClient
}

func NewVirtualMachineService(cs *KCPSClient) *VirtualMachineService {
	return &VirtualMachineService{cs: cs}
}

type VolumeService struct {
	cs *KCPSClient
}

func NewVolumeService(cs *KCPSClient) *VolumeService {
	return &VolumeService{cs: cs}
}

type TagsService struct {
	cs *KCPSClient
}

func NewTagsService(cs *KCPSClient) *TagsService {
	return &TagsService{cs: cs}
}
