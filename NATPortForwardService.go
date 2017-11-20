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

type ListPortForwardingRulesParams struct {
	p map[string]interface{}
}

func (p *ListPortForwardingRulesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["fordisplay"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("fordisplay", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
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
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
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

func (p *ListPortForwardingRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetFordisplay(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["fordisplay"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListPortForwardingRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

// You should always use this function to get a new ListPortForwardingRulesParams instance,
// as then you are sure you have configured all required params
func (s *NatPortForwardService) NewListPortForwardingRulesParams() *ListPortForwardingRulesParams {
	p := &ListPortForwardingRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists all port forwarding rules for an IP address.
func (s *NatPortForwardService) ListPortForwardingRules(p *ListPortForwardingRulesParams) (*ListPortForwardingRulesResponse, error) {
	resp, err := s.cs.newRequest("listPortForwardingRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = convertFirewallServiceResponse(resp)
	if err != nil {
		return nil, err
	}

	var r ListPortForwardingRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPortForwardingRulesResponse struct {
	Count               int                   `json:"count"`
	PortForwardingRules []*PortForwardingRule `json:"portforwardingrule"`
}

type PortForwardingRule struct {
	Cidrlist       string `json:"cidrlist,omitempty"`
	Fordisplay     bool   `json:"fordisplay,omitempty"`
	Id             string `json:"id,omitempty"`
	Ipaddress      string `json:"ipaddress,omitempty"`
	Ipaddressid    string `json:"ipaddressid,omitempty"`
	Networkid      string `json:"networkid,omitempty"`
	Privateendport string `json:"privateendport,omitempty"`
	Privateport    string `json:"privateport,omitempty"`
	Protocol       string `json:"protocol,omitempty"`
	Publicendport  string `json:"publicendport,omitempty"`
	Publicport     string `json:"publicport,omitempty"`
	State          string `json:"state,omitempty"`
	Tags           []struct {
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
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname,omitempty"`
	Virtualmachineid          string `json:"virtualmachineid,omitempty"`
	Virtualmachinename        string `json:"virtualmachinename,omitempty"`
	Vmguestip                 string `json:"vmguestip,omitempty"`
}

type CreatePortForwardingRuleParams struct {
	p map[string]interface{}
}

func (p *CreatePortForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := p.p["privateport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateport", vv)
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["publicport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicport", vv)
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["privateendport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateendport", vv)
	}
	if v, found := p.p["publicendport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicendport", vv)
	}
	if v, found := p.p["vmguestip"]; found {
		u.Set("vmguestip", v.(string))
	}
	return u
}

func (p *CreatePortForwardingRuleParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetPrivateport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privateport"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetPublicport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicport"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetPrivateendport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privateendport"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetPublicendport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicendport"] = v
	return
}

func (p *CreatePortForwardingRuleParams) SetVmguestip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmguestip"] = v
	return
}

// You should always use this function to get a new CreatePortForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *NatPortForwardService) NewCreatePortForwardingRuleParams(ipaddressid string, privateport int, protocol string, publicport int, virtualmachineid string) *CreatePortForwardingRuleParams {
	p := &CreatePortForwardingRuleParams{}
	p.p = make(map[string]interface{})
	p.p["ipaddressid"] = ipaddressid
	p.p["privateport"] = privateport
	p.p["protocol"] = protocol
	p.p["publicport"] = publicport
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Creates a port forwarding rule
func (s *NatPortForwardService) CreatePortForwardingRule(p *CreatePortForwardingRuleParams) (*CreatePortForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("createPortForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreatePortForwardingRuleResponse
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

type CreatePortForwardingRuleResponse struct {
	JobID          string `json:"jobid,omitempty"`
	Cidrlist       string `json:"cidrlist,omitempty"`
	Fordisplay     bool   `json:"fordisplay,omitempty"`
	Id             string `json:"id,omitempty"`
	Ipaddress      string `json:"ipaddress,omitempty"`
	Ipaddressid    string `json:"ipaddressid,omitempty"`
	Networkid      string `json:"networkid,omitempty"`
	Privateendport string `json:"privateendport,omitempty"`
	Privateport    string `json:"privateport,omitempty"`
	Protocol       string `json:"protocol,omitempty"`
	Publicendport  string `json:"publicendport,omitempty"`
	Publicport     string `json:"publicport,omitempty"`
	State          string `json:"state,omitempty"`
	Tags           []struct {
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
	Virtualmachinedisplayname string `json:"virtualmachinedisplayname,omitempty"`
	Virtualmachineid          string `json:"virtualmachineid,omitempty"`
	Virtualmachinename        string `json:"virtualmachinename,omitempty"`
	Vmguestip                 string `json:"vmguestip,omitempty"`
}

type DeletePortForwardingRuleParams struct {
	p map[string]interface{}
}

func (p *DeletePortForwardingRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeletePortForwardingRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeletePortForwardingRuleParams instance,
// as then you are sure you have configured all required params
func (s *NatPortForwardService) NewDeletePortForwardingRuleParams(id string) *DeletePortForwardingRuleParams {
	p := &DeletePortForwardingRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a port forwarding rule
func (s *NatPortForwardService) DeletePortForwardingRule(p *DeletePortForwardingRuleParams) (*DeletePortForwardingRuleResponse, error) {
	resp, err := s.cs.newRequest("deletePortForwardingRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeletePortForwardingRuleResponse
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

type DeletePortForwardingRuleResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}
