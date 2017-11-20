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

type IptoNetworklistParams struct {
	Networkid string
	Ipv4      string
	Ipv6      string
}

func (p *IptoNetworklistParams) SetNetworkid(networkid string) {
	p.Networkid = networkid
}

func (p *IptoNetworklistParams) SetIpv4(ipv4 string) {
	p.Ipv4 = ipv4
}

func (p *IptoNetworklistParams) SetIpv6(ipv6 string) {
	p.Ipv6 = ipv6
}

func (*VirtualMachineService) NewIptoNetworklistParams(networkid string) IptoNetworklistParams {
	p := IptoNetworklistParams{}
	p.Networkid = networkid
	return p
}

type DeployValueVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *DeployValueVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["networkids"]; found {
		for i, vv := range v.([]IptoNetworklistParams) {
			u.Set(fmt.Sprintf("iptonetworklist[%d].networkid", i), vv.Networkid)
			if vv.Ipv4 != "" {
				u.Set(fmt.Sprintf("iptonetworklist[%d].ip", i), vv.Ipv4)
			}
			if vv.Ipv6 != "" {
				u.Set(fmt.Sprintf("iptonetworklist[%d].ipv6", i), vv.Ipv6)
			}
		}
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeployValueVirtualMachineParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
	return
}

func (p *DeployValueVirtualMachineParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *DeployValueVirtualMachineParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *DeployValueVirtualMachineParams) SetIptoNetworklist(v []IptoNetworklistParams) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkids"] = v
	return
}

func (p *DeployValueVirtualMachineParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

func (p *DeployValueVirtualMachineParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
	return
}

func (p *DeployValueVirtualMachineParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
	return
}

func (p *DeployValueVirtualMachineParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new DeployValueVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewDeployValueVirtualMachineParams(serviceofferingid string, templateid string, zoneid string, name string) *DeployValueVirtualMachineParams {
	p := &DeployValueVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["serviceofferingid"] = serviceofferingid
	p.p["templateid"] = templateid
	p.p["zoneid"] = zoneid
	p.p["hypervisor"] = "VMware"
	p.p["name"] = name
	return p
}

// Creates and automatically starts a virtual machine based on a service offering, disk offering, and template.
func (s *VirtualMachineService) DeployValueVirtualMachine(p *DeployValueVirtualMachineParams) (*DeployValueVirtualMachineResponse, error) {

	if p.p["networkid"] == nil {
		netid, err := getNetworkIdByName(s.cs, "PublicFrontSegment")
		if err != nil {
			return nil, err
		}
		nparams := s.cs.VirtualMachine.NewIptoNetworklistParams(netid)
		p.SetIptoNetworklist([]IptoNetworklistParams{nparams})
	}

	resp, err := s.cs.newRequest("deployValueVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeployValueVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetExAsyncJobResult(r.JobID, s.cs.timeout)
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

		// for kcps api response
		b, err = cnvCorrectVirtualMachineJson(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return &r, nil
}

type DeployValueVirtualMachineResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Account       string `json:"account,omitempty"`
	Affinitygroup []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
		Project           string   `json:"project,omitempty"`
		Projectid         string   `json:"projectid,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Created               string            `json:"created,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Diskioread            int64             `json:"diskioread,omitempty"`
	Diskiowrite           int64             `json:"diskiowrite,omitempty"`
	Diskkbsread           int64             `json:"diskkbsread,omitempty"`
	Diskkbswrite          int64             `json:"diskkbswrite,omitempty"`
	Diskofferingid        string            `json:"diskofferingid,omitempty"`
	Diskofferingname      string            `json:"diskofferingname,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Networkkbsread        int64             `json:"networkkbsread,omitempty"`
	Networkkbswrite       int64             `json:"networkkbswrite,omitempty"`
	Nic                   []struct {
		Broadcasturi string `json:"broadcasturi,omitempty"`
		Deviceid     string `json:"deviceid,omitempty"`
		Gateway      string `json:"gateway,omitempty"`
		Id           string `json:"id,omitempty"`
		Ip6address   string `json:"ip6address,omitempty"`
		Ip6cidr      string `json:"ip6cidr,omitempty"`
		Ip6gateway   string `json:"ip6gateway,omitempty"`
		Ipaddress    string `json:"ipaddress,omitempty"`
		Isdefault    bool   `json:"isdefault,omitempty"`
		Isolationuri string `json:"isolationuri,omitempty"`
		Macaddress   string `json:"macaddress,omitempty"`
		Netmask      string `json:"netmask,omitempty"`
		Networkid    string `json:"networkid,omitempty"`
		Networkname  string `json:"networkname,omitempty"`
		Secondaryip  []struct {
			Id        string `json:"id,omitempty"`
			Ipaddress string `json:"ipaddress,omitempty"`
		} `json:"secondaryip,omitempty"`
		Traffictype      string `json:"traffictype,omitempty"`
		Type             string `json:"type,omitempty"`
		Virtualmachineid string `json:"virtualmachineid,omitempty"`
	} `json:"nic,omitempty"`
	Ostypeid        int64  `json:"ostypeid,omitempty"`
	Password        string `json:"password,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Project         string `json:"project,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
	Publicip        string `json:"publicip,omitempty"`
	Publicipid      string `json:"publicipid,omitempty"`
	Rootdeviceid    int64  `json:"rootdeviceid,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Securitygroup   []struct {
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Egressrule  []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"ingressrule,omitempty"`
		Name      string `json:"name,omitempty"`
		Project   string `json:"project,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Tags      []struct {
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
		Virtualmachinecount int      `json:"virtualmachinecount,omitempty"`
		Virtualmachineids   []string `json:"virtualmachineids,omitempty"`
	} `json:"securitygroup,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	State               string `json:"state,omitempty"`
	Tags                []struct {
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
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Userid              string `json:"userid,omitempty"`
	Username            string `json:"username,omitempty"`
	Vgpu                string `json:"vgpu,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type DestroyVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *DestroyVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["expunge"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("expunge", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

// You should always use this function to get a new DestroyVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewDestroyVirtualMachineParams(id string) *DestroyVirtualMachineParams {
	p := &DestroyVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Destroys a virtual machine.
func (s *VirtualMachineService) DestroyVirtualMachine(p *DestroyVirtualMachineParams) (*DestroyVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("destroyVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DestroyVirtualMachineResponse
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

		var r DeployValueVirtualMachineResponse
		if err := json.Unmarshal(resp, &r); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DestroyVirtualMachineResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Account       string `json:"account,omitempty"`
	Affinitygroup []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
		Project           string   `json:"project,omitempty"`
		Projectid         string   `json:"projectid,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Created               string            `json:"created,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Diskioread            int64             `json:"diskioread,omitempty"`
	Diskiowrite           int64             `json:"diskiowrite,omitempty"`
	Diskkbsread           int64             `json:"diskkbsread,omitempty"`
	Diskkbswrite          int64             `json:"diskkbswrite,omitempty"`
	Diskofferingid        string            `json:"diskofferingid,omitempty"`
	Diskofferingname      string            `json:"diskofferingname,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Networkkbsread        int64             `json:"networkkbsread,omitempty"`
	Networkkbswrite       int64             `json:"networkkbswrite,omitempty"`
	Nic                   []struct {
		Broadcasturi string `json:"broadcasturi,omitempty"`
		Deviceid     string `json:"deviceid,omitempty"`
		Gateway      string `json:"gateway,omitempty"`
		Id           string `json:"id,omitempty"`
		Ip6address   string `json:"ip6address,omitempty"`
		Ip6cidr      string `json:"ip6cidr,omitempty"`
		Ip6gateway   string `json:"ip6gateway,omitempty"`
		Ipaddress    string `json:"ipaddress,omitempty"`
		Isdefault    bool   `json:"isdefault,omitempty"`
		Isolationuri string `json:"isolationuri,omitempty"`
		Macaddress   string `json:"macaddress,omitempty"`
		Netmask      string `json:"netmask,omitempty"`
		Networkid    string `json:"networkid,omitempty"`
		Networkname  string `json:"networkname,omitempty"`
		Secondaryip  []struct {
			Id        string `json:"id,omitempty"`
			Ipaddress string `json:"ipaddress,omitempty"`
		} `json:"secondaryip,omitempty"`
		Traffictype      string `json:"traffictype,omitempty"`
		Type             string `json:"type,omitempty"`
		Virtualmachineid string `json:"virtualmachineid,omitempty"`
	} `json:"nic,omitempty"`
	Ostypeid        int64  `json:"ostypeid,omitempty"`
	Password        string `json:"password,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Project         string `json:"project,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
	Publicip        string `json:"publicip,omitempty"`
	Publicipid      string `json:"publicipid,omitempty"`
	Rootdeviceid    int64  `json:"rootdeviceid,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Securitygroup   []struct {
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Egressrule  []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"ingressrule,omitempty"`
		Name      string `json:"name,omitempty"`
		Project   string `json:"project,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Tags      []struct {
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
		Virtualmachinecount int      `json:"virtualmachinecount,omitempty"`
		Virtualmachineids   []string `json:"virtualmachineids,omitempty"`
	} `json:"securitygroup,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	State               string `json:"state,omitempty"`
	Tags                []struct {
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
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Userid              string `json:"userid,omitempty"`
	Username            string `json:"username,omitempty"`
	Vgpu                string `json:"vgpu,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type RebootVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *RebootVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RebootVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new RebootVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewRebootVirtualMachineParams(id string) *RebootVirtualMachineParams {
	p := &RebootVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Reboots a virtual machine.
func (s *VirtualMachineService) RebootVirtualMachine(p *RebootVirtualMachineParams) (*RebootVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("rebootVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RebootVirtualMachineResponse
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

		// for kcps api response
		b, err = cnvCorrectVirtualMachineJson(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type RebootVirtualMachineResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Account       string `json:"account,omitempty"`
	Affinitygroup []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
		Project           string   `json:"project,omitempty"`
		Projectid         string   `json:"projectid,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Created               string            `json:"created,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Diskioread            int64             `json:"diskioread,omitempty"`
	Diskiowrite           int64             `json:"diskiowrite,omitempty"`
	Diskkbsread           int64             `json:"diskkbsread,omitempty"`
	Diskkbswrite          int64             `json:"diskkbswrite,omitempty"`
	Diskofferingid        string            `json:"diskofferingid,omitempty"`
	Diskofferingname      string            `json:"diskofferingname,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Networkkbsread        int64             `json:"networkkbsread,omitempty"`
	Networkkbswrite       int64             `json:"networkkbswrite,omitempty"`
	Nic                   []struct {
		Broadcasturi string `json:"broadcasturi,omitempty"`
		Deviceid     string `json:"deviceid,omitempty"`
		Gateway      string `json:"gateway,omitempty"`
		Id           string `json:"id,omitempty"`
		Ip6address   string `json:"ip6address,omitempty"`
		Ip6cidr      string `json:"ip6cidr,omitempty"`
		Ip6gateway   string `json:"ip6gateway,omitempty"`
		Ipaddress    string `json:"ipaddress,omitempty"`
		Isdefault    bool   `json:"isdefault,omitempty"`
		Isolationuri string `json:"isolationuri,omitempty"`
		Macaddress   string `json:"macaddress,omitempty"`
		Netmask      string `json:"netmask,omitempty"`
		Networkid    string `json:"networkid,omitempty"`
		Networkname  string `json:"networkname,omitempty"`
		Secondaryip  []struct {
			Id        string `json:"id,omitempty"`
			Ipaddress string `json:"ipaddress,omitempty"`
		} `json:"secondaryip,omitempty"`
		Traffictype      string `json:"traffictype,omitempty"`
		Type             string `json:"type,omitempty"`
		Virtualmachineid string `json:"virtualmachineid,omitempty"`
	} `json:"nic,omitempty"`
	Ostypeid        int64  `json:"ostypeid,omitempty"`
	Password        string `json:"password,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Project         string `json:"project,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
	Publicip        string `json:"publicip,omitempty"`
	Publicipid      string `json:"publicipid,omitempty"`
	Rootdeviceid    int64  `json:"rootdeviceid,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Securitygroup   []struct {
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Egressrule  []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"ingressrule,omitempty"`
		Name      string `json:"name,omitempty"`
		Project   string `json:"project,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Tags      []struct {
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
		Virtualmachinecount int      `json:"virtualmachinecount,omitempty"`
		Virtualmachineids   []string `json:"virtualmachineids,omitempty"`
	} `json:"securitygroup,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	State               string `json:"state,omitempty"`
	Tags                []struct {
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
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Userid              string `json:"userid,omitempty"`
	Username            string `json:"username,omitempty"`
	Vgpu                string `json:"vgpu,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type StartVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *StartVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StartVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new StartVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewStartVirtualMachineParams(id string) *StartVirtualMachineParams {
	p := &StartVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Starts a virtual machine.
func (s *VirtualMachineService) StartVirtualMachine(p *StartVirtualMachineParams) (*StartVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("startVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetExAsyncJobResult(r.JobID, s.cs.timeout)
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

		// for kcps api response
		b, err = cnvCorrectVirtualMachineJson(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type StartVirtualMachineResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Account       string `json:"account,omitempty"`
	Affinitygroup []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
		Project           string   `json:"project,omitempty"`
		Projectid         string   `json:"projectid,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Created               string            `json:"created,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Diskioread            int64             `json:"diskioread,omitempty"`
	Diskiowrite           int64             `json:"diskiowrite,omitempty"`
	Diskkbsread           int64             `json:"diskkbsread,omitempty"`
	Diskkbswrite          int64             `json:"diskkbswrite,omitempty"`
	Diskofferingid        string            `json:"diskofferingid,omitempty"`
	Diskofferingname      string            `json:"diskofferingname,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Networkkbsread        int64             `json:"networkkbsread,omitempty"`
	Networkkbswrite       int64             `json:"networkkbswrite,omitempty"`
	Nic                   []struct {
		Broadcasturi string `json:"broadcasturi,omitempty"`
		Deviceid     string `json:"deviceid,omitempty"`
		Gateway      string `json:"gateway,omitempty"`
		Id           string `json:"id,omitempty"`
		Ip6address   string `json:"ip6address,omitempty"`
		Ip6cidr      string `json:"ip6cidr,omitempty"`
		Ip6gateway   string `json:"ip6gateway,omitempty"`
		Ipaddress    string `json:"ipaddress,omitempty"`
		Isdefault    bool   `json:"isdefault,omitempty"`
		Isolationuri string `json:"isolationuri,omitempty"`
		Macaddress   string `json:"macaddress,omitempty"`
		Netmask      string `json:"netmask,omitempty"`
		Networkid    string `json:"networkid,omitempty"`
		Networkname  string `json:"networkname,omitempty"`
		Secondaryip  []struct {
			Id        string `json:"id,omitempty"`
			Ipaddress string `json:"ipaddress,omitempty"`
		} `json:"secondaryip,omitempty"`
		Traffictype      string `json:"traffictype,omitempty"`
		Type             string `json:"type,omitempty"`
		Virtualmachineid string `json:"virtualmachineid,omitempty"`
	} `json:"nic,omitempty"`
	Ostypeid        int64  `json:"ostypeid,omitempty"`
	Password        string `json:"password,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Project         string `json:"project,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
	Publicip        string `json:"publicip,omitempty"`
	Publicipid      string `json:"publicipid,omitempty"`
	Rootdeviceid    int64  `json:"rootdeviceid,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Securitygroup   []struct {
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Egressrule  []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"ingressrule,omitempty"`
		Name      string `json:"name,omitempty"`
		Project   string `json:"project,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Tags      []struct {
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
		Virtualmachinecount int      `json:"virtualmachinecount,omitempty"`
		Virtualmachineids   []string `json:"virtualmachineids,omitempty"`
	} `json:"securitygroup,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	State               string `json:"state,omitempty"`
	Tags                []struct {
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
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Userid              string `json:"userid,omitempty"`
	Username            string `json:"username,omitempty"`
	Vgpu                string `json:"vgpu,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type StopVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *StopVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StopVirtualMachineParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
	return
}

func (p *StopVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new StopVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewStopVirtualMachineParams(id string) *StopVirtualMachineParams {
	p := &StopVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Stops a virtual machine.
func (s *VirtualMachineService) StopVirtualMachine(p *StopVirtualMachineParams) (*StopVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("stopVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = cnvCorrectVirtualMachineJson(resp)
	if err != nil {
		return nil, err
	}

	var r StopVirtualMachineResponse
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

		// for kcps api response
		b, err = cnvCorrectVirtualMachineJson(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type StopVirtualMachineResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Account       string `json:"account,omitempty"`
	Affinitygroup []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
		Project           string   `json:"project,omitempty"`
		Projectid         string   `json:"projectid,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Created               string            `json:"created,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Diskioread            int64             `json:"diskioread,omitempty"`
	Diskiowrite           int64             `json:"diskiowrite,omitempty"`
	Diskkbsread           int64             `json:"diskkbsread,omitempty"`
	Diskkbswrite          int64             `json:"diskkbswrite,omitempty"`
	Diskofferingid        string            `json:"diskofferingid,omitempty"`
	Diskofferingname      string            `json:"diskofferingname,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Networkkbsread        int64             `json:"networkkbsread,omitempty"`
	Networkkbswrite       int64             `json:"networkkbswrite,omitempty"`
	Nic                   []struct {
		Broadcasturi string `json:"broadcasturi,omitempty"`
		Deviceid     string `json:"deviceid,omitempty"`
		Gateway      string `json:"gateway,omitempty"`
		Id           string `json:"id,omitempty"`
		Ip6address   string `json:"ip6address,omitempty"`
		Ip6cidr      string `json:"ip6cidr,omitempty"`
		Ip6gateway   string `json:"ip6gateway,omitempty"`
		Ipaddress    string `json:"ipaddress,omitempty"`
		Isdefault    bool   `json:"isdefault,omitempty"`
		Isolationuri string `json:"isolationuri,omitempty"`
		Macaddress   string `json:"macaddress,omitempty"`
		Netmask      string `json:"netmask,omitempty"`
		Networkid    string `json:"networkid,omitempty"`
		Networkname  string `json:"networkname,omitempty"`
		Secondaryip  []struct {
			Id        string `json:"id,omitempty"`
			Ipaddress string `json:"ipaddress,omitempty"`
		} `json:"secondaryip,omitempty"`
		Traffictype      string `json:"traffictype,omitempty"`
		Type             string `json:"type,omitempty"`
		Virtualmachineid string `json:"virtualmachineid,omitempty"`
	} `json:"nic,omitempty"`
	Ostypeid        int64  `json:"ostypeid,omitempty"`
	Password        string `json:"password,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Project         string `json:"project,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
	Publicip        string `json:"publicip,omitempty"`
	Publicipid      string `json:"publicipid,omitempty"`
	Rootdeviceid    int64  `json:"rootdeviceid,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Securitygroup   []struct {
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Egressrule  []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"ingressrule,omitempty"`
		Name      string `json:"name,omitempty"`
		Project   string `json:"project,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Tags      []struct {
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
		Virtualmachinecount int      `json:"virtualmachinecount,omitempty"`
		Virtualmachineids   []string `json:"virtualmachineids,omitempty"`
	} `json:"securitygroup,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	State               string `json:"state,omitempty"`
	Tags                []struct {
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
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Userid              string `json:"userid,omitempty"`
	Username            string `json:"username,omitempty"`
	Vgpu                string `json:"vgpu,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type ResetPasswordForVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *ResetPasswordForVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ResetPasswordForVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new ResetPasswordForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewResetPasswordForVirtualMachineParams(id string) *ResetPasswordForVirtualMachineParams {
	p := &ResetPasswordForVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Resets the password for virtual machine. The virtual machine must be in a "Stopped" state and the template must already support this feature for this command to take effect. [async]
func (s *VirtualMachineService) ResetPasswordForVirtualMachine(p *ResetPasswordForVirtualMachineParams) (*ResetPasswordForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("resetPasswordForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = cnvCorrectVirtualMachineJson(resp)
	if err != nil {
		return nil, err
	}

	var r ResetPasswordForVirtualMachineResponse
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

		// for kcps api response
		b, err = cnvCorrectVirtualMachineJson(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type ResetPasswordForVirtualMachineResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Account       string `json:"account,omitempty"`
	Affinitygroup []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
		Project           string   `json:"project,omitempty"`
		Projectid         string   `json:"projectid,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Created               string            `json:"created,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Diskioread            int64             `json:"diskioread,omitempty"`
	Diskiowrite           int64             `json:"diskiowrite,omitempty"`
	Diskkbsread           int64             `json:"diskkbsread,omitempty"`
	Diskkbswrite          int64             `json:"diskkbswrite,omitempty"`
	Diskofferingid        string            `json:"diskofferingid,omitempty"`
	Diskofferingname      string            `json:"diskofferingname,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Networkkbsread        int64             `json:"networkkbsread,omitempty"`
	Networkkbswrite       int64             `json:"networkkbswrite,omitempty"`
	Nic                   []struct {
		Broadcasturi string `json:"broadcasturi,omitempty"`
		Deviceid     string `json:"deviceid,omitempty"`
		Gateway      string `json:"gateway,omitempty"`
		Id           string `json:"id,omitempty"`
		Ip6address   string `json:"ip6address,omitempty"`
		Ip6cidr      string `json:"ip6cidr,omitempty"`
		Ip6gateway   string `json:"ip6gateway,omitempty"`
		Ipaddress    string `json:"ipaddress,omitempty"`
		Isdefault    bool   `json:"isdefault,omitempty"`
		Isolationuri string `json:"isolationuri,omitempty"`
		Macaddress   string `json:"macaddress,omitempty"`
		Netmask      string `json:"netmask,omitempty"`
		Networkid    string `json:"networkid,omitempty"`
		Networkname  string `json:"networkname,omitempty"`
		Secondaryip  []struct {
			Id        string `json:"id,omitempty"`
			Ipaddress string `json:"ipaddress,omitempty"`
		} `json:"secondaryip,omitempty"`
		Traffictype      string `json:"traffictype,omitempty"`
		Type             string `json:"type,omitempty"`
		Virtualmachineid string `json:"virtualmachineid,omitempty"`
	} `json:"nic,omitempty"`
	Ostypeid        int64  `json:"ostypeid,omitempty"`
	Password        string `json:"password,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Project         string `json:"project,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
	Publicip        string `json:"publicip,omitempty"`
	Publicipid      string `json:"publicipid,omitempty"`
	Rootdeviceid    int64  `json:"rootdeviceid,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Securitygroup   []struct {
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Egressrule  []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"ingressrule,omitempty"`
		Name      string `json:"name,omitempty"`
		Project   string `json:"project,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Tags      []struct {
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
		Virtualmachinecount int      `json:"virtualmachinecount,omitempty"`
		Virtualmachineids   []string `json:"virtualmachineids,omitempty"`
	} `json:"securitygroup,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	State               string `json:"state,omitempty"`
	Tags                []struct {
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
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Userid              string `json:"userid,omitempty"`
	Username            string `json:"username,omitempty"`
	Vgpu                string `json:"vgpu,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type ListVirtualMachinesParams struct {
	p map[string]interface{}
}

func (p *ListVirtualMachinesParams) toURLValues() url.Values {
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
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVirtualMachinesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}
func (p *ListVirtualMachinesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListVirtualMachinesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListVirtualMachinesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}
func (p *ListVirtualMachinesParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListVirtualMachinesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListVirtualMachinesParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
	return
}

func (p *ListVirtualMachinesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListVirtualMachinesParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewListVirtualMachinesParams() *ListVirtualMachinesParams {
	p := &ListVirtualMachinesParams{}
	p.p = make(map[string]interface{})
	return p
}

// List the virtual machines owned by the account.
func (s *VirtualMachineService) ListVirtualMachines(p *ListVirtualMachinesParams) (*ListVirtualMachinesResponse, error) {
	resp, err := s.cs.newRequest("listVirtualMachines", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVirtualMachinesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVirtualMachinesResponse struct {
	Count           int               `json:"count"`
	VirtualMachines []*VirtualMachine `json:"virtualmachine"`
}

type VirtualMachine struct {
	Account       string `json:"account,omitempty"`
	Affinitygroup []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
		Project           string   `json:"project,omitempty"`
		Projectid         string   `json:"projectid,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Created               string            `json:"created,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Diskioread            int64             `json:"diskioread,omitempty"`
	Diskiowrite           int64             `json:"diskiowrite,omitempty"`
	Diskkbsread           int64             `json:"diskkbsread,omitempty"`
	Diskkbswrite          int64             `json:"diskkbswrite,omitempty"`
	Diskofferingid        string            `json:"diskofferingid,omitempty"`
	Diskofferingname      string            `json:"diskofferingname,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Networkkbsread        int64             `json:"networkkbsread,omitempty"`
	Networkkbswrite       int64             `json:"networkkbswrite,omitempty"`
	Nic                   []struct {
		Broadcasturi string `json:"broadcasturi,omitempty"`
		Deviceid     string `json:"deviceid,omitempty"`
		Gateway      string `json:"gateway,omitempty"`
		Id           string `json:"id,omitempty"`
		Ip6address   string `json:"ip6address,omitempty"`
		Ip6cidr      string `json:"ip6cidr,omitempty"`
		Ip6gateway   string `json:"ip6gateway,omitempty"`
		Ipaddress    string `json:"ipaddress,omitempty"`
		Isdefault    bool   `json:"isdefault,omitempty"`
		Isolationuri string `json:"isolationuri,omitempty"`
		Macaddress   string `json:"macaddress,omitempty"`
		Netmask      string `json:"netmask,omitempty"`
		Networkid    string `json:"networkid,omitempty"`
		Networkname  string `json:"networkname,omitempty"`
		Secondaryip  []struct {
			Id        string `json:"id,omitempty"`
			Ipaddress string `json:"ipaddress,omitempty"`
		} `json:"secondaryip,omitempty"`
		Traffictype      string `json:"traffictype,omitempty"`
		Type             string `json:"type,omitempty"`
		Virtualmachineid string `json:"virtualmachineid,omitempty"`
	} `json:"nic,omitempty"`
	Ostypeid        int64  `json:"ostypeid,omitempty"`
	Password        string `json:"password,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Project         string `json:"project,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
	Publicip        string `json:"publicip,omitempty"`
	Publicipid      string `json:"publicipid,omitempty"`
	Rootdeviceid    int64  `json:"rootdeviceid,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Securitygroup   []struct {
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Egressrule  []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"ingressrule,omitempty"`
		Name      string `json:"name,omitempty"`
		Project   string `json:"project,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Tags      []struct {
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
		Virtualmachinecount int      `json:"virtualmachinecount,omitempty"`
		Virtualmachineids   []string `json:"virtualmachineids,omitempty"`
	} `json:"securitygroup,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	State               string `json:"state,omitempty"`
	Tags                []struct {
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
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Userid              string `json:"userid,omitempty"`
	Username            string `json:"username,omitempty"`
	Vgpu                string `json:"vgpu,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type ChangeServiceForVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *ChangeServiceForVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (p *ChangeServiceForVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ChangeServiceForVirtualMachineParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

// You should always use this function to get a new ChangeServiceForVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewChangeServiceForVirtualMachineParams(id string, serviceofferingid string) *ChangeServiceForVirtualMachineParams {
	p := &ChangeServiceForVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Changes the service offering for a virtual machine. The virtual machine must be in a "Stopped" state for this command to take effect.
func (s *VirtualMachineService) ChangeServiceForVirtualMachine(p *ChangeServiceForVirtualMachineParams) (*ChangeServiceForVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("changeServiceForVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = cnvCorrectVirtualMachineJson(resp)
	if err != nil {
		return nil, err
	}

	var r ChangeServiceForVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ChangeServiceForVirtualMachineResponse struct {
	Account       string `json:"account,omitempty"`
	Affinitygroup []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
		Project           string   `json:"project,omitempty"`
		Projectid         string   `json:"projectid,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Cpuused               int               `json:"cpuused,omitempty"`
	Created               string            `json:"created,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Diskioread            int64             `json:"diskioread,omitempty"`
	Diskiowrite           int64             `json:"diskiowrite,omitempty"`
	Diskkbsread           int64             `json:"diskkbsread,omitempty"`
	Diskkbswrite          int64             `json:"diskkbswrite,omitempty"`
	Diskofferingid        string            `json:"diskofferingid,omitempty"`
	Diskofferingname      string            `json:"diskofferingname,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Networkkbsread        int64             `json:"networkkbsread,omitempty"`
	Networkkbswrite       int64             `json:"networkkbswrite,omitempty"`
	Nic                   []struct {
		Broadcasturi string `json:"broadcasturi,omitempty"`
		Deviceid     string `json:"deviceid,omitempty"`
		Gateway      string `json:"gateway,omitempty"`
		Id           string `json:"id,omitempty"`
		Ip6address   string `json:"ip6address,omitempty"`
		Ip6cidr      string `json:"ip6cidr,omitempty"`
		Ip6gateway   string `json:"ip6gateway,omitempty"`
		Ipaddress    string `json:"ipaddress,omitempty"`
		Isdefault    bool   `json:"isdefault,omitempty"`
		Isolationuri string `json:"isolationuri,omitempty"`
		Macaddress   string `json:"macaddress,omitempty"`
		Netmask      string `json:"netmask,omitempty"`
		Networkid    string `json:"networkid,omitempty"`
		Networkname  string `json:"networkname,omitempty"`
		Secondaryip  []struct {
			Id        string `json:"id,omitempty"`
			Ipaddress string `json:"ipaddress,omitempty"`
		} `json:"secondaryip,omitempty"`
		Traffictype      string `json:"traffictype,omitempty"`
		Type             string `json:"type,omitempty"`
		Virtualmachineid string `json:"virtualmachineid,omitempty"`
	} `json:"nic,omitempty"`
	Ostypeid        int64  `json:"ostypeid,omitempty"`
	Password        string `json:"password,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Project         string `json:"project,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
	Publicip        string `json:"publicip,omitempty"`
	Publicipid      string `json:"publicipid,omitempty"`
	Rootdeviceid    int64  `json:"rootdeviceid,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Securitygroup   []struct {
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Egressrule  []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"ingressrule,omitempty"`
		Name      string `json:"name,omitempty"`
		Project   string `json:"project,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Tags      []struct {
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
		Virtualmachinecount int      `json:"virtualmachinecount,omitempty"`
		Virtualmachineids   []string `json:"virtualmachineids,omitempty"`
	} `json:"securitygroup,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	State               string `json:"state,omitempty"`
	Tags                []struct {
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
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Userid              string `json:"userid,omitempty"`
	Username            string `json:"username,omitempty"`
	Vgpu                string `json:"vgpu,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type ScaleVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *ScaleVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (p *ScaleVirtualMachineParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ScaleVirtualMachineParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

// You should always use this function to get a new ScaleVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewScaleVirtualMachineParams(id string, serviceofferingid string) *ScaleVirtualMachineParams {
	p := &ScaleVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Scales the virtual machine to a new service offering.
func (s *VirtualMachineService) ScaleVirtualMachine(p *ScaleVirtualMachineParams) (*ScaleVirtualMachineResponse, error) {
	resp, err := s.cs.newRequest("scaleVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	resp, err = cnvCorrectVirtualMachineJson(resp)
	if err != nil {
		return nil, err
	}

	var r ScaleVirtualMachineResponse
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

type ScaleVirtualMachineResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

func cnvCorrectVirtualMachineJson(message json.RawMessage) (json.RawMessage, error) {
	var m interface{}
	if err := json.Unmarshal(message, &m); err != nil {
		return nil, err
	}

	vm := m.(map[string]interface{})

	if cpunum, ok := vm["cpunumber"].(string); ok {
		cpunumstr, err := strconv.Atoi(cpunum)
		if err != nil {
			return nil, err
		}
		vm["cpunumber"] = int(cpunumstr)
	}

	if _, ok := vm["cpuspeed"].(string); ok {
		cpuspeed, err := strconv.Atoi(vm["cpuspeed"].(string))
		if err != nil {
			return nil, err
		}
		vm["cpuspeed"] = int(cpuspeed)
	}

	if _, ok := vm["created"].(float64); ok {
		utime := int64(vm["created"].(float64)) / 1000
		t := time.Unix(utime, 0)
		vm["created"] = t.Format("2006-01-02T15:04:05+09:00")
	}

	if _, ok := vm["rootdeviceid"].(string); ok {
		rootdeviceid, err := strconv.ParseInt(vm["rootdeviceid"].(string), 10, 64)
		if err != nil {
			return nil, err
		}
		vm["rootdeviceid"] = int64(rootdeviceid)
	}

	b, err := json.Marshal(vm)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(b), nil
}

// For KCPS premium host
type DeployPremiumVirtualMachineParams struct {
	p map[string]interface{}
}

func (p *DeployPremiumVirtualMachineParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["hostname"]; found {
		u.Set("hostname", v.(string))
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["networkids"]; found {
		for i, vv := range v.([]IptoNetworklistParams) {
			u.Set(fmt.Sprintf("iptonetworklist[%d].networkid", i), vv.Networkid)
			if vv.Ipv4 != "" {
				u.Set(fmt.Sprintf("iptonetworklist[%d].ip", i), vv.Ipv4)
			}
			if vv.Ipv6 != "" {
				u.Set(fmt.Sprintf("iptonetworklist[%d].ipv6", i), vv.Ipv6)
			}
		}
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	if v, found := p.p["templateid"]; found {
		u.Set("templateid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeployPremiumVirtualMachineParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
	return
}

func (p *DeployPremiumVirtualMachineParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
	return
}

func (p *DeployPremiumVirtualMachineParams) SetHostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostname"] = v
	return
}

func (p *DeployPremiumVirtualMachineParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *DeployPremiumVirtualMachineParams) SetNetworkids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkids"] = v
	return
}

func (p *DeployPremiumVirtualMachineParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

func (p *DeployPremiumVirtualMachineParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
	return
}

func (p *DeployPremiumVirtualMachineParams) SetTemplateid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["templateid"] = v
	return
}

func (p *DeployPremiumVirtualMachineParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

func (p *DeployPremiumVirtualMachineParams) SetIptoNetworklist(v []IptoNetworklistParams) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkids"] = v
	return
}

// You should always use this function to get a new DeployPremiumVirtualMachineParams instance,
// as then you are sure you have configured all required params
func (s *VirtualMachineService) NewDeployPremiumVirtualMachineParams(serviceofferingid string, templateid string, zoneid string, name string, hostname string) *DeployPremiumVirtualMachineParams {
	p := &DeployPremiumVirtualMachineParams{}
	p.p = make(map[string]interface{})
	p.p["serviceofferingid"] = serviceofferingid
	p.p["templateid"] = templateid
	p.p["zoneid"] = zoneid
	p.p["hostname"] = hostname
	p.p["name"] = name
	p.p["hypervisor"] = "VMware"
	return p
}

// Creates and automatically starts a virtual machine based on a service offering, disk offering, and template.
func (s *VirtualMachineService) DeployPremiumVirtualMachine(p *DeployPremiumVirtualMachineParams) (*DeployPremiumVirtualMachineResponse, error) {
	if p.p["networkid"] == nil {
		netid, err := getNetworkIdByName(s.cs, "PublicFrontSegment")
		if err != nil {
			return nil, err
		}
		nparams := s.cs.VirtualMachine.NewIptoNetworklistParams(netid)
		p.SetIptoNetworklist([]IptoNetworklistParams{nparams})
	}
	resp, err := s.cs.newRequest("deployPremiumVirtualMachine", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeployPremiumVirtualMachineResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetExAsyncJobResult(r.JobID, s.cs.timeout)
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

		// for kcps api response
		b, err = cnvCorrectVirtualMachineJson(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return &r, nil
}

type DeployPremiumVirtualMachineResponse struct {
	JobID         string `json:"jobid,omitempty"`
	Account       string `json:"account,omitempty"`
	Affinitygroup []struct {
		Account           string   `json:"account,omitempty"`
		Description       string   `json:"description,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Id                string   `json:"id,omitempty"`
		Name              string   `json:"name,omitempty"`
		Project           string   `json:"project,omitempty"`
		Projectid         string   `json:"projectid,omitempty"`
		Type              string   `json:"type,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Cpunumber             int               `json:"cpunumber,omitempty"`
	Cpuspeed              int               `json:"cpuspeed,omitempty"`
	Cpuused               string            `json:"cpuused,omitempty"`
	Created               string            `json:"created,omitempty"`
	Details               map[string]string `json:"details,omitempty"`
	Diskioread            int64             `json:"diskioread,omitempty"`
	Diskiowrite           int64             `json:"diskiowrite,omitempty"`
	Diskkbsread           int64             `json:"diskkbsread,omitempty"`
	Diskkbswrite          int64             `json:"diskkbswrite,omitempty"`
	Diskofferingid        string            `json:"diskofferingid,omitempty"`
	Diskofferingname      string            `json:"diskofferingname,omitempty"`
	Displayname           string            `json:"displayname,omitempty"`
	Displayvm             bool              `json:"displayvm,omitempty"`
	Domain                string            `json:"domain,omitempty"`
	Domainid              string            `json:"domainid,omitempty"`
	Forvirtualnetwork     bool              `json:"forvirtualnetwork,omitempty"`
	Group                 string            `json:"group,omitempty"`
	Groupid               string            `json:"groupid,omitempty"`
	Guestosid             string            `json:"guestosid,omitempty"`
	Haenable              bool              `json:"haenable,omitempty"`
	Hostid                string            `json:"hostid,omitempty"`
	Hostname              string            `json:"hostname,omitempty"`
	Hypervisor            string            `json:"hypervisor,omitempty"`
	Id                    string            `json:"id,omitempty"`
	Instancename          string            `json:"instancename,omitempty"`
	Isdynamicallyscalable bool              `json:"isdynamicallyscalable,omitempty"`
	Isodisplaytext        string            `json:"isodisplaytext,omitempty"`
	Isoid                 string            `json:"isoid,omitempty"`
	Isoname               string            `json:"isoname,omitempty"`
	Keypair               string            `json:"keypair,omitempty"`
	Memory                int               `json:"memory,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Networkkbsread        int64             `json:"networkkbsread,omitempty"`
	Networkkbswrite       int64             `json:"networkkbswrite,omitempty"`
	Nic                   []struct {
		Broadcasturi string `json:"broadcasturi,omitempty"`
		Deviceid     string `json:"deviceid,omitempty"`
		Gateway      string `json:"gateway,omitempty"`
		Id           string `json:"id,omitempty"`
		Ip6address   string `json:"ip6address,omitempty"`
		Ip6cidr      string `json:"ip6cidr,omitempty"`
		Ip6gateway   string `json:"ip6gateway,omitempty"`
		Ipaddress    string `json:"ipaddress,omitempty"`
		Isdefault    bool   `json:"isdefault,omitempty"`
		Isolationuri string `json:"isolationuri,omitempty"`
		Macaddress   string `json:"macaddress,omitempty"`
		Netmask      string `json:"netmask,omitempty"`
		Networkid    string `json:"networkid,omitempty"`
		Networkname  string `json:"networkname,omitempty"`
		Secondaryip  []struct {
			Id        string `json:"id,omitempty"`
			Ipaddress string `json:"ipaddress,omitempty"`
		} `json:"secondaryip,omitempty"`
		Traffictype      string `json:"traffictype,omitempty"`
		Type             string `json:"type,omitempty"`
		Virtualmachineid string `json:"virtualmachineid,omitempty"`
	} `json:"nic,omitempty"`
	Ostypeid        int64  `json:"ostypeid,omitempty"`
	Password        string `json:"password,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Project         string `json:"project,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
	Publicip        string `json:"publicip,omitempty"`
	Publicipid      string `json:"publicipid,omitempty"`
	Rootdeviceid    int64  `json:"rootdeviceid,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Securitygroup   []struct {
		Account     string `json:"account,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Egressrule  []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"egressrule,omitempty"`
		Id          string `json:"id,omitempty"`
		Ingressrule []struct {
			Account           string `json:"account,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Tags              []struct {
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
		} `json:"ingressrule,omitempty"`
		Name      string `json:"name,omitempty"`
		Project   string `json:"project,omitempty"`
		Projectid string `json:"projectid,omitempty"`
		Tags      []struct {
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
		Virtualmachinecount int      `json:"virtualmachinecount,omitempty"`
		Virtualmachineids   []string `json:"virtualmachineids,omitempty"`
	} `json:"securitygroup,omitempty"`
	Serviceofferingid   string `json:"serviceofferingid,omitempty"`
	Serviceofferingname string `json:"serviceofferingname,omitempty"`
	Servicestate        string `json:"servicestate,omitempty"`
	State               string `json:"state,omitempty"`
	Tags                []struct {
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
	Templatedisplaytext string `json:"templatedisplaytext,omitempty"`
	Templateid          string `json:"templateid,omitempty"`
	Templatename        string `json:"templatename,omitempty"`
	Userid              string `json:"userid,omitempty"`
	Username            string `json:"username,omitempty"`
	Vgpu                string `json:"vgpu,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

func getNetworkIdByName(cs *KCPSClient, networkname string) (string, error) {
	params := cs.AccountDomain.NewListNetworksParams()
	params.SetKeyword(networkname)
	resp, err := cs.AccountDomain.ListNetworks(params)
	if err != nil {
		return "", err
	}

	if resp.Count > 1 {
		return "", fmt.Errorf("Multiple network found. Network Name:", networkname)
	}

	if resp.Count == 0 {
		return "", fmt.Errorf("Network not found. Network Name:", networkname)
	}
	return resp.Networks[0].Id, nil
}
