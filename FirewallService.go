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
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

// Helper function for maintaining backwards compatibility
func convertFirewallServiceResponse(b []byte) ([]byte, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	if _, ok := raw["firewallrule"]; ok {
		return convertFirewallServiceListResponse(b)
	}

	for _, k := range []string{"endport", "startport"} {
		if sVal, ok := raw[k].(string); ok {
			iVal, err := strconv.Atoi(sVal)
			if err != nil {
				return nil, err
			}
			raw[k] = iVal
		}
	}

	return json.Marshal(raw)
}

// Helper function for maintaining backwards compatibility
func convertFirewallServiceListResponse(b []byte) ([]byte, error) {
	var rawList struct {
		Count         int                      `json:"count"`
		FirewallRules []map[string]interface{} `json:"firewallrule"`
	}

	if err := json.Unmarshal(b, &rawList); err != nil {
		return nil, err
	}

	for _, r := range rawList.FirewallRules {
		for _, k := range []string{"endport", "startport"} {
			if sVal, ok := r[k].(string); ok {
				iVal, err := strconv.Atoi(sVal)
				if err != nil {
					return nil, err
				}
				r[k] = iVal
			}
		}
	}

	return json.Marshal(rawList)
}

type ListFirewallRulesParams struct {
	p map[string]interface{}
}

func (p *ListFirewallRulesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	return u
}

func (p *ListFirewallRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListFirewallRulesParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
	return
}
func (p *ListFirewallRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}
func (p *ListFirewallRulesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

// You should always use this function to get a new ListFirewallRulesParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewListFirewallRulesParams() *ListFirewallRulesParams {
	p := &ListFirewallRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all firewall rules for an IP address.
func (s *FirewallService) ListFirewallRules(p *ListFirewallRulesParams) (*ListFirewallRulesResponse, error) {
	resp, err := s.cs.newRequest("listFirewallRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = convertFirewallServiceResponse(resp)
	if err != nil {
		return nil, err
	}

	var r ListFirewallRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListFirewallRulesResponse struct {
	Count         int             `json:"count"`
	FirewallRules []*FirewallRule `json:"firewallrule"`
}

type FirewallRule struct {
	Cidrlist    string `json:"cidrlist,omitempty"`
	Endport     int    `json:"endport,omitempty"`
	Fordisplay  bool   `json:"fordisplay,omitempty"`
	Icmpcode    int    `json:"icmpcode,omitempty"`
	Icmptype    int    `json:"icmptype,omitempty"`
	Id          string `json:"id,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Ipaddressid string `json:"ipaddressid,omitempty"`
	Networkid   string `json:"networkid,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	Startport   int    `json:"startport,omitempty"`
	State       string `json:"state,omitempty"`
	Tags        []struct {
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
}

type CreateFirewallRuleParams struct {
	p map[string]interface{}
}

func (p *CreateFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := p.p["icmpcode"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmpcode", vv)
	}
	if v, found := p.p["icmptype"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("icmptype", vv)
	}
	return u
}

func (p *CreateFirewallRuleParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
	return
}

func (p *CreateFirewallRuleParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
	return
}

func (p *CreateFirewallRuleParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
	return
}

func (p *CreateFirewallRuleParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
	return
}

func (p *CreateFirewallRuleParams) SetIcmpcode(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmpcode"] = v
	return
}

func (p *CreateFirewallRuleParams) SetIcmptype(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["icmptype"] = v
	return
}

func (p *CreateFirewallRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
	return
}

// You should always use this function to get a new CreateFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewCreateFirewallRuleParams(ipaddressid string, protocol string, ciderlist []string) *CreateFirewallRuleParams {
	p := &CreateFirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddressid"] = ipaddressid
	p.p["protocol"] = protocol
	p.p["cidrlist"] = ciderlist
	return p
}

// Creates a firewall rule for a given IP address
func (s *FirewallService) CreateFirewallRule(p *CreateFirewallRuleParams) (*CreateFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("createFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateFirewallRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateFirewallRuleResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Cidrlist    string `json:"cidrlist,omitempty"`
	Endport     int    `json:"endport,omitempty"`
	Fordisplay  bool   `json:"fordisplay,omitempty"`
	Icmpcode    int    `json:"icmpcode,omitempty"`
	Icmptype    int    `json:"icmptype,omitempty"`
	Id          string `json:"id,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Ipaddressid string `json:"ipaddressid,omitempty"`
	Networkid   string `json:"networkid,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	Startport   int    `json:"startport,omitempty"`
	State       string `json:"state,omitempty"`
	Tags        []struct {
		Account      string `json:"account,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
}

type DeleteFirewallRuleParams struct {
	p map[string]interface{}
}

func (p *DeleteFirewallRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteFirewallRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteFirewallRuleParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDeleteFirewallRuleParams(id string) *DeleteFirewallRuleParams {
	p := &DeleteFirewallRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a firewall rule
func (s *FirewallService) DeleteFirewallRule(p *DeleteFirewallRuleParams) (*DeleteFirewallRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteFirewallRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteFirewallRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = convertFirewallServiceResponse(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DeleteFirewallRuleResponse struct {
	JobID   string `json:"jobid,omitempty"`
	Success bool   `json:"success,omitempty"`
}

type EnableStaticNatParams struct {
	p map[string]interface{}
}

func (p *EnableStaticNatParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["vmguestip"]; found {
		u.Set("vmguestip", v.(string))
	}
	return u
}

func (p *EnableStaticNatParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
	return
}

func (p *EnableStaticNatParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *EnableStaticNatParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

func (p *EnableStaticNatParams) SetVmguestip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmguestip"] = v
	return
}

// You should always use this function to get a new EnableStaticNatParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewEnableStaticNatParams(ipaddressid string, virtualmachineid string) *EnableStaticNatParams {
	p := &EnableStaticNatParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddressid"] = ipaddressid
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Enables static NAT for given IP address
func (s *FirewallService) EnableStaticNat(p *EnableStaticNatParams) (*EnableStaticNatResponse, error) {
	resp, err := s.cs.newRequest("enableStaticNat", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableStaticNatResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type EnableStaticNatResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     string `json:"success,omitempty"`
}

type DisableStaticNatParams struct {
	p map[string]interface{}
}

func (p *DisableStaticNatParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	return u
}

func (p *DisableStaticNatParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
	return
}

// You should always use this function to get a new DisableStaticNatParams instance,
// as then you are sure you have configured all required params
func (s *FirewallService) NewDisableStaticNatParams(ipaddressid string) *DisableStaticNatParams {
	p := &DisableStaticNatParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddressid"] = ipaddressid
	return p
}

// Disables static rule for given IP address
func (s *FirewallService) DisableStaticNat(p *DisableStaticNatParams) (*DisableStaticNatResponse, error) {
	resp, err := s.cs.newRequest("disableStaticNat", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DisableStaticNatResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DisableStaticNatResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}
