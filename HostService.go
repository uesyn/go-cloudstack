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
	"time"
)

type ListPremiumHostsParams struct {
	p map[string]interface{}
}

func (p *ListPremiumHostsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListPremiumHostsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListHostsParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewListPremiumHostsParams() *ListPremiumHostsParams {
	p := &ListPremiumHostsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists hosts.
func (s *HostService) ListPremiumHosts(p *ListPremiumHostsParams) (*ListPremiumHostsResponse, error) {
	resp, err := s.cs.newRequest("listPremiumHosts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = cnvCorrectPremiumHostJson(resp)
	if err != nil {
		return nil, err
	}

	var r ListPremiumHostsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListPremiumHostsResponse struct {
	Count        int            `json:"count"`
	PremiumHosts []*PremiumHost `json:"host"`
}

type PremiumHost struct {
	Cpuallocated          string `json:"cpuallocated,omitempty"`
	Cpunumber             int64  `json:"cpunumber,omitempty"`
	Cpuspeed              int64  `json:"cpuspeed,omitempty"`
	Cpuused               string `json:"cpuused,omitempty"`
	Hypervisor            string `json:"hypervisor,omitempty"`
	Memoryallocated       int64  `json:"memoryallocated,omitempty"`
	Memorytotal           int64  `json:"memorytotal,omitempty"`
	Memoryused            int64  `json:"memoryused,omitempty"`
	Name                  string `json:"name,omitempty"`
	Zoneid                string `json:"zoneid,omitempty"`
	Resourcestate         string `json:"resourcestate,omitempty"`
	State                 string `json:"state,omitempty"`
	Zonename              string `json:"zonename,omitempty"`
	DistributionGroupname string `json:"distributionGroupName,omitempty"`
}

type ListDistributionGroupsParams struct {
	p map[string]interface{}
}

func (p *ListDistributionGroupsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	return u
}

func (s *HostService) NewListDistributionGroupsParams() *ListDistributionGroupsParams {
	p := &ListDistributionGroupsParams{}
	p.p = make(map[string]interface{})
	return p
}

func (s *HostService) ListDistributionGroups(p *ListDistributionGroupsParams) (*ListDistributionGroupsResponse, error) {
	resp, err := s.cs.newRequest("listDistributionGroups", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListDistributionGroupsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListDistributionGroupsResponse struct {
	Count              int                  `json:"count"`
	DistributionGroups []*DistributionGroup `json:"distributiongroup"`
}

type DistributionGroup struct {
	Name string `json:"name"`
}

type ListPremiumVirtualMachinesParams struct {
	p map[string]interface{}
}

func (p *ListPremiumVirtualMachinesParams) toURLValues() url.Values {
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListPremiumVirtualMachinesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListPremiumVirtualMachinesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListPremiumVirtualMachinesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListPremiumVirtualMachinesParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListPremiumVirtualMachinesParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
	return
}

func (p *ListPremiumVirtualMachinesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

func (s *HostService) NewListPremiumVirtualMachines() *ListPremiumVirtualMachinesParams {
	p := &ListPremiumVirtualMachinesParams{}
	p.p = make(map[string]interface{})
	return p
}

func (s *HostService) ListPremiumVirtualMachines(p *ListPremiumVirtualMachinesParams) (*ListPremiumVirtualMachinesResponse, error) {

	var r ListPremiumVirtualMachinesResponse

	resp, _ := s.cs.newRequest("listPremiumVirtualMachines", p.toURLValues())
	var i interface{}
	if err := json.Unmarshal(resp, &i); err != nil {
		fmt.Println(err)
	}

	v := i.(map[string]interface{})["virtualmachine"].([]interface{})
	for _, m := range v {
		vmtmp := m.(map[string]interface{})
		rdint, err := strconv.Atoi(vmtmp["rootdeviceid"].(string))
		if err != nil {
			return nil, err
		}
		vmtmp["rootdeviceid"] = int64(rdint)

		utime := int64(vmtmp["created"].(float64)) / 1000
		t := time.Unix(utime, 0)
		vmtmp["created"] = t.Format("2006-01-02T15:04:05+09:00")
		fmt.Println("")
	}

	b, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(json.RawMessage(b), &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListPremiumVirtualMachinesResponse struct {
	Count int `json:"count"`
	// VirtualMachine struct is written in VirtualMachineService.go
	VirtualMachines []*VirtualMachine `json:"virtualmachine"`
}

type AddPremiumHostParams struct {
	p map[string]interface{}
}

func (p *AddPremiumHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["number"]; found {

		u.Set("number", strconv.Itoa(v.(int)))
	}
	if v, found := p.p["distributiongroup"]; found {
		u.Set("distributiongroup", v.(string))
	}
	return u
}

func (p *AddPremiumHostParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

func (p *AddPremiumHostParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *AddPremiumHostParams) SetNumber(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["number"] = v
	return
}

func (p *AddPremiumHostParams) SetDistributiongroup(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["distributiongroup"] = v
	return
}

// You should always use this function to get a new AddPremiumHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewAddPremiumHostParams(hypervisor, zoneid string, number int) *AddPremiumHostParams {
	p := &AddPremiumHostParams{}
	p.p = make(map[string]interface{})
	p.p["hypervisor"] = hypervisor
	p.p["zoneid"] = zoneid
	p.p["number"] = number
	return p
}

// Adds a new host.
func (s *HostService) AddPremiumHost(p *AddPremiumHostParams) (*AddPremiumHostResponse, error) {
	resp, err := s.cs.newRequest("addPremiumHosts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = cnvCorrectPremiumHostJson(resp)
	if err != nil {
		return nil, err
	}

	var r AddPremiumHostResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddPremiumHostResponse struct {
	Count        int            `json:"count"`
	PremiumHosts []*PremiumHost `json:"host"`
}

type RemovePremiumHostParams struct {
	p map[string]interface{}
}

func (p *RemovePremiumHostParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *RemovePremiumHostParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

// You should always use this function to get a new RemovePremiumHostParams instance,
// as then you are sure you have configured all required params
func (s *HostService) NewRemovePremiumHostParams(name string) *RemovePremiumHostParams {
	p := &RemovePremiumHostParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	return p
}

// Deletes a host.
func (s *HostService) RemovePremiumHost(p *RemovePremiumHostParams) (*RemovePremiumHostResponse, error) {
	resp, err := s.cs.newRequest("removePremiumHost", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemovePremiumHostResponse

	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type RemovePremiumHostResponse struct {
	Success bool `json:"success,omitempty"`
}

func cnvCorrectPremiumHostJson(message json.RawMessage) (json.RawMessage, error) {
	var i interface{}
	if err := json.Unmarshal(message, &i); err != nil {
		return nil, err
	}
	v := i.(map[string]interface{})["host"].([]interface{})
	for _, m := range v {
		premhost := m.(map[string]interface{})

		cpunum, err := strconv.Atoi(premhost["cpunumber"].(string))
		if err != nil {
			return nil, err
		}
		premhost["cpunumber"] = int64(cpunum)

		cpuspeed, err := strconv.Atoi(premhost["cpuspeed"].(string))
		if err != nil {
			return nil, err
		}
		premhost["cpuspeed"] = int64(cpuspeed)

		memalloc, err := strconv.Atoi(premhost["memoryallocated"].(string))
		if err != nil {
			return nil, err
		}
		premhost["memoryallocated"] = int64(memalloc)

		memtotal, err := strconv.Atoi(premhost["memorytotal"].(string))
		if err != nil {
			return nil, err
		}
		premhost["memorytotal"] = int64(memtotal)

		memused, err := strconv.Atoi(premhost["memoryused"].(string))
		if err != nil {
			return nil, err
		}
		premhost["memoryused"] = int64(memused)
	}

	b, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(b), nil
}
