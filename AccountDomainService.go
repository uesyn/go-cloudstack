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
	"fmt"
	"net/url"
	"strconv"
)

type ListUsersParams struct {
	p map[string]interface{}
}

func (p *ListUsersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["accounttype"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("accounttype", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *ListUsersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListUsersParams) SetAccounttype(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accounttype"] = v
	return
}

func (p *ListUsersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListUsersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListUsersParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListUsersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListUsersParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListUsersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListUsersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListUsersParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListUsersParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new ListUsersParams instance,
// as then you are sure you have configured all required params
func (s *AccountDomainService) NewListUsersParams() *ListUsersParams {
	p := &ListUsersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists user accounts
func (s *AccountDomainService) ListUsers(p *ListUsersParams) (*ListUsersResponse, error) {
	resp, err := s.cs.newRequest("listUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListUsersResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"user"`
}

type User struct {
	Account             string `json:"account,omitempty"`
	Accountid           string `json:"accountid,omitempty"`
	Accounttype         int    `json:"accounttype,omitempty"`
	Apikey              string `json:"apikey,omitempty"`
	Created             string `json:"created,omitempty"`
	Domain              string `json:"domain,omitempty"`
	Domainid            string `json:"domainid,omitempty"`
	Email               string `json:"email,omitempty"`
	Firstname           string `json:"firstname,omitempty"`
	Id                  string `json:"id,omitempty"`
	Iscallerchilddomain bool   `json:"iscallerchilddomain,omitempty"`
	Isdefault           bool   `json:"isdefault,omitempty"`
	Lastname            string `json:"lastname,omitempty"`
	Secretkey           string `json:"secretkey,omitempty"`
	State               string `json:"state,omitempty"`
	Timezone            string `json:"timezone,omitempty"`
	Username            string `json:"username,omitempty"`
}

type ListNetworksParams struct {
	p map[string]interface{}
}

func (p *ListNetworksParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListNetworksParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListNetworksParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListNetworksParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}
func (p *ListNetworksParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListNetworksParams instance,
// as then you are sure you have configured all required params
func (s *AccountDomainService) NewListNetworksParams() *ListNetworksParams {
	p := &ListNetworksParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all available networks.
func (s *AccountDomainService) ListNetworks(p *ListNetworksParams) (*ListNetworksResponse, error) {
	resp, err := s.cs.newRequest("listNetworks", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworksResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListNetworksResponse struct {
	Count    int        `json:"count"`
	Networks []*Network `json:"network"`
}

type Network struct {
	Account                     string `json:"account,omitempty"`
	Aclid                       string `json:"aclid,omitempty"`
	Acltype                     string `json:"acltype,omitempty"`
	Broadcastdomaintype         string `json:"broadcastdomaintype,omitempty"`
	Broadcasturi                string `json:"broadcasturi,omitempty"`
	Canusefordeploy             bool   `json:"canusefordeploy,omitempty"`
	Cidr                        string `json:"cidr,omitempty"`
	Displaynetwork              bool   `json:"displaynetwork,omitempty"`
	Displaytext                 string `json:"displaytext,omitempty"`
	Dns1                        string `json:"dns1,omitempty"`
	Dns2                        string `json:"dns2,omitempty"`
	Domain                      string `json:"domain,omitempty"`
	Domainid                    string `json:"domainid,omitempty"`
	Gateway                     string `json:"gateway,omitempty"`
	Id                          string `json:"id,omitempty"`
	Ip6cidr                     string `json:"ip6cidr,omitempty"`
	Ip6gateway                  string `json:"ip6gateway,omitempty"`
	Isdefault                   bool   `json:"isdefault,omitempty"`
	Ispersistent                bool   `json:"ispersistent,omitempty"`
	Issystem                    bool   `json:"issystem,omitempty"`
	Name                        string `json:"name,omitempty"`
	Netmask                     string `json:"netmask,omitempty"`
	Networkcidr                 string `json:"networkcidr,omitempty"`
	Networkdomain               string `json:"networkdomain,omitempty"`
	Networkofferingavailability string `json:"networkofferingavailability,omitempty"`
	Networkofferingconservemode bool   `json:"networkofferingconservemode,omitempty"`
	Networkofferingdisplaytext  string `json:"networkofferingdisplaytext,omitempty"`
	Networkofferingid           string `json:"networkofferingid,omitempty"`
	Networkofferingname         string `json:"networkofferingname,omitempty"`
	Physicalnetworkid           string `json:"physicalnetworkid,omitempty"`
	Project                     string `json:"project,omitempty"`
	Projectid                   string `json:"projectid,omitempty"`
	Related                     string `json:"related,omitempty"`
	Reservediprange             string `json:"reservediprange,omitempty"`
	Restartrequired             bool   `json:"restartrequired,omitempty"`
	Service                     []struct {
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
		} `json:"capability,omitempty"`
		Name     string `json:"name,omitempty"`
		Provider []struct {
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			State                        string   `json:"state,omitempty"`
		} `json:"provider,omitempty"`
	} `json:"service,omitempty"`
	Specifyipranges  bool   `json:"specifyipranges,omitempty"`
	State            string `json:"state,omitempty"`
	Strechedl2subnet bool   `json:"strechedl2subnet,omitempty"`
	Subdomainaccess  bool   `json:"subdomainaccess,omitempty"`
	Tags             []struct {
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
	Traffictype       string   `json:"traffictype,omitempty"`
	Type              string   `json:"type,omitempty"`
	Vlan              string   `json:"vlan,omitempty"`
	Vpcid             string   `json:"vpcid,omitempty"`
	Zoneid            string   `json:"zoneid,omitempty"`
	Zonename          string   `json:"zonename,omitempty"`
	Zonesnetworkspans []string `json:"zonesnetworkspans,omitempty"`
}

type ListServiceOfferingsParams struct {
	p map[string]interface{}
}

func (p *ListServiceOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["issystem"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("issystem", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["systemvmtype"]; found {
		u.Set("systemvmtype", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *ListServiceOfferingsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListServiceOfferingsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListServiceOfferingsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListServiceOfferingsParams) SetIssystem(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["issystem"] = v
	return
}

func (p *ListServiceOfferingsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListServiceOfferingsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListServiceOfferingsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListServiceOfferingsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListServiceOfferingsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListServiceOfferingsParams) SetSystemvmtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["systemvmtype"] = v
	return
}

func (p *ListServiceOfferingsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new ListServiceOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *AccountDomainService) NewListServiceOfferingsParams() *ListServiceOfferingsParams {
	p := &ListServiceOfferingsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all available service offerings.
func (s *AccountDomainService) ListServiceOfferings(p *ListServiceOfferingsParams) (*ListServiceOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listServiceOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListServiceOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListServiceOfferingsResponse struct {
	Count            int                `json:"count"`
	ServiceOfferings []*ServiceOffering `json:"serviceoffering"`
}

type ServiceOffering struct {
	Cpunumber                 int               `json:"cpunumber,omitempty"`
	Cpuspeed                  int               `json:"cpuspeed,omitempty"`
	Created                   string            `json:"created,omitempty"`
	Defaultuse                bool              `json:"defaultuse,omitempty"`
	Deploymentplanner         string            `json:"deploymentplanner,omitempty"`
	DiskBytesReadRate         int64             `json:"diskBytesReadRate,omitempty"`
	DiskBytesWriteRate        int64             `json:"diskBytesWriteRate,omitempty"`
	DiskIopsReadRate          int64             `json:"diskIopsReadRate,omitempty"`
	DiskIopsWriteRate         int64             `json:"diskIopsWriteRate,omitempty"`
	Displaytext               string            `json:"displaytext,omitempty"`
	Domain                    string            `json:"domain,omitempty"`
	Domainid                  string            `json:"domainid,omitempty"`
	Hosttags                  string            `json:"hosttags,omitempty"`
	Hypervisorsnapshotreserve int               `json:"hypervisorsnapshotreserve,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Iscustomized              bool              `json:"iscustomized,omitempty"`
	Iscustomizediops          bool              `json:"iscustomizediops,omitempty"`
	Issystem                  bool              `json:"issystem,omitempty"`
	Isvolatile                bool              `json:"isvolatile,omitempty"`
	Limitcpuuse               bool              `json:"limitcpuuse,omitempty"`
	Maxiops                   int64             `json:"maxiops,omitempty"`
	Memory                    int               `json:"memory,omitempty"`
	Miniops                   int64             `json:"miniops,omitempty"`
	Name                      string            `json:"name,omitempty"`
	Networkrate               int               `json:"networkrate,omitempty"`
	Offerha                   bool              `json:"offerha,omitempty"`
	Provisioningtype          string            `json:"provisioningtype,omitempty"`
	Serviceofferingdetails    map[string]string `json:"serviceofferingdetails,omitempty"`
	Storagetype               string            `json:"storagetype,omitempty"`
	Systemvmtype              string            `json:"systemvmtype,omitempty"`
	Tags                      string            `json:"tags,omitempty"`
}

type ListDiskOfferingsParams struct {
	p map[string]interface{}
}

func (p *ListDiskOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListDiskOfferingsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListDiskOfferingsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListDiskOfferingsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListDiskOfferingsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListDiskOfferingsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListDiskOfferingsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListDiskOfferingsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListDiskOfferingsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListDiskOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *AccountDomainService) NewListDiskOfferingsParams() *ListDiskOfferingsParams {
	p := &ListDiskOfferingsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all available disk offerings.
func (s *AccountDomainService) ListDiskOfferings(p *ListDiskOfferingsParams) (*ListDiskOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listDiskOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDiskOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDiskOfferingsResponse struct {
	Count         int             `json:"count"`
	DiskOfferings []*DiskOffering `json:"diskoffering"`
}

type DiskOffering struct {
	CacheMode                 string `json:"cacheMode,omitempty"`
	Created                   string `json:"created,omitempty"`
	DiskBytesReadRate         int64  `json:"diskBytesReadRate,omitempty"`
	DiskBytesWriteRate        int64  `json:"diskBytesWriteRate,omitempty"`
	DiskIopsReadRate          int64  `json:"diskIopsReadRate,omitempty"`
	DiskIopsWriteRate         int64  `json:"diskIopsWriteRate,omitempty"`
	Disksize                  int64  `json:"disksize,omitempty"`
	Displayoffering           bool   `json:"displayoffering,omitempty"`
	Displaytext               string `json:"displaytext,omitempty"`
	Domain                    string `json:"domain,omitempty"`
	Domainid                  string `json:"domainid,omitempty"`
	Hypervisorsnapshotreserve int    `json:"hypervisorsnapshotreserve,omitempty"`
	Id                        string `json:"id,omitempty"`
	Iscustomized              bool   `json:"iscustomized,omitempty"`
	Iscustomizediops          bool   `json:"iscustomizediops,omitempty"`
	Maxiops                   int64  `json:"maxiops,omitempty"`
	Miniops                   int64  `json:"miniops,omitempty"`
	Name                      string `json:"name,omitempty"`
	Provisioningtype          string `json:"provisioningtype,omitempty"`
	Storagetype               string `json:"storagetype,omitempty"`
	Tags                      string `json:"tags,omitempty"`
}

type ListZonesParams struct {
	p map[string]interface{}
}

func (p *ListZonesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["available"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("available", vv)
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networktype"]; found {
		u.Set("networktype", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["showcapacities"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("showcapacities", vv)
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	return u
}

func (p *ListZonesParams) SetAvailable(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["available"] = v
	return
}

func (p *ListZonesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListZonesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListZonesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListZonesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListZonesParams) SetNetworktype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networktype"] = v
	return
}

func (p *ListZonesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListZonesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListZonesParams) SetShowcapacities(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["showcapacities"] = v
	return
}

func (p *ListZonesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

// You should always use this function to get a new ListZonesParams instance,
// as then you are sure you have configured all required params
func (s *AccountDomainService) NewListZonesParams() *ListZonesParams {
	p := &ListZonesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists zones
func (s *AccountDomainService) ListZones(p *ListZonesParams) (*ListZonesResponse, error) {
	resp, err := s.cs.newRequest("listZones", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListZonesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListZonesResponse struct {
	Count int     `json:"count"`
	Zones []*Zone `json:"zone"`
}

type Zone struct {
	Allocationstate string `json:"allocationstate,omitempty"`
	Capacity        []struct {
		Capacitytotal int64  `json:"capacitytotal,omitempty"`
		Capacityused  int64  `json:"capacityused,omitempty"`
		Clusterid     string `json:"clusterid,omitempty"`
		Clustername   string `json:"clustername,omitempty"`
		Percentused   string `json:"percentused,omitempty"`
		Podid         string `json:"podid,omitempty"`
		Podname       string `json:"podname,omitempty"`
		Type          int    `json:"type,omitempty"`
		Zoneid        string `json:"zoneid,omitempty"`
		Zonename      string `json:"zonename,omitempty"`
	} `json:"capacity,omitempty"`
	Description           string            `json:"description,omitempty"`
	Dhcpprovider          string            `json:"dhcpprovider,omitempty"`
	Displaytext           string            `json:"displaytext,omitempty"`
	Dns1                  string            `json:"dns1,omitempty"`
	Dns2                  string            `json:"dns2,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Domainname            string            `json:"domainname,omitempty"`
	Guestcidraddress      string            `json:"guestcidraddress,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Internaldns1          string            `json:"internaldns1,omitempty"`
	Internaldns2          string            `json:"internaldns2,omitempty"`
	Ip6dns1               string            `json:"ip6dns1,omitempty"`
	Ip6dns2               string            `json:"ip6dns2,omitempty"`
	Localstorageenabled   bool              `json:"localstorageenabled,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Networktype           string            `json:"networktype,omitempty"`
	Resourcedetails       map[string]string `json:"resourcedetails,omitempty"`
	Securitygroupsenabled bool              `json:"securitygroupsenabled,omitempty"`
	Tags                  []struct {
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
	Zonetoken string `json:"zonetoken,omitempty"`
}
