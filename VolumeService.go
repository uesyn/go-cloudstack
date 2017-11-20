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
	"strings"
)

type AttachVolumeParams struct {
	p map[string]interface{}
}

func (p *AttachVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *AttachVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *AttachVolumeParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new AttachVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewAttachVolumeParams(id string, virtualmachineid string) *AttachVolumeParams {
	p := &AttachVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Attaches a disk volume to a virtual machine.
func (s *VolumeService) AttachVolume(p *AttachVolumeParams) (*AttachVolumeResponse, error) {
	resp, err := s.cs.newRequest("attachVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AttachVolumeResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type AttachVolumeResponse struct {
	JobID                      string `json:"jobid,omitempty"`
	Account                    string `json:"account,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Chaininfo                  string `json:"chaininfo,omitempty"`
	Created                    string `json:"created,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Deviceid                   int64  `json:"deviceid,omitempty"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate,omitempty"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate,omitempty"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate,omitempty"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Id                         string `json:"id,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Isodisplaytext             string `json:"isodisplaytext,omitempty"`
	Isoid                      string `json:"isoid,omitempty"`
	Isoname                    string `json:"isoname,omitempty"`
	Maxiops                    int64  `json:"maxiops,omitempty"`
	Miniops                    int64  `json:"miniops,omitempty"`
	Name                       string `json:"name,omitempty"`
	Path                       string `json:"path,omitempty"`
	Project                    string `json:"project,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Provisioningtype           string `json:"provisioningtype,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Size                       int64  `json:"size,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	State                      string `json:"state,omitempty"`
	Status                     string `json:"status,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Tags                       []struct {
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
	Type                string `json:"type,omitempty"`
	Virtualmachineid    string `json:"virtualmachineid,omitempty"`
	Vmdisplayname       string `json:"vmdisplayname,omitempty"`
	Vmname              string `json:"vmname,omitempty"`
	Vmstate             string `json:"vmstate,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type DetachVolumeParams struct {
	p map[string]interface{}
}

func (p *DetachVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DetachVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DetachVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewDetachVolumeParams() *DetachVolumeParams {
	p := &DetachVolumeParams{}
	p.p = make(map[string]interface{})
	return p
}

// Detaches a disk volume from a virtual machine.
func (s *VolumeService) DetachVolume(p *DetachVolumeParams) (*DetachVolumeResponse, error) {
	resp, err := s.cs.newRequest("detachVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DetachVolumeResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type DetachVolumeResponse struct {
	JobID                      string `json:"jobid,omitempty"`
	Account                    string `json:"account,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Chaininfo                  string `json:"chaininfo,omitempty"`
	Created                    string `json:"created,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Deviceid                   int64  `json:"deviceid,omitempty"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate,omitempty"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate,omitempty"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate,omitempty"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Id                         string `json:"id,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Isodisplaytext             string `json:"isodisplaytext,omitempty"`
	Isoid                      string `json:"isoid,omitempty"`
	Isoname                    string `json:"isoname,omitempty"`
	Maxiops                    int64  `json:"maxiops,omitempty"`
	Miniops                    int64  `json:"miniops,omitempty"`
	Name                       string `json:"name,omitempty"`
	Path                       string `json:"path,omitempty"`
	Project                    string `json:"project,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Provisioningtype           string `json:"provisioningtype,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Size                       int64  `json:"size,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	State                      string `json:"state,omitempty"`
	Status                     string `json:"status,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Tags                       []struct {
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
	Type                string `json:"type,omitempty"`
	Virtualmachineid    string `json:"virtualmachineid,omitempty"`
	Vmdisplayname       string `json:"vmdisplayname,omitempty"`
	Vmname              string `json:"vmname,omitempty"`
	Vmstate             string `json:"vmstate,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type CreateVolumeParams struct {
	p map[string]interface{}
}

func (p *CreateVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["diskofferingid"]; found {
		u.Set("diskofferingid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	if v, found := p.p["snapshotid"]; found {
		u.Set("snapshotid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateVolumeParams) SetDiskofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["diskofferingid"] = v
	return
}

func (p *CreateVolumeParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateVolumeParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
	return
}

func (p *CreateVolumeParams) SetSnapshotid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["snapshotid"] = v
	return
}

func (p *CreateVolumeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new CreateVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewCreateVolumeParams() *CreateVolumeParams {
	p := &CreateVolumeParams{}
	p.p = make(map[string]interface{})
	return p
}

// Creates a disk volume from a disk offering. This disk volume must still be attached to a virtual machine to make use of it.
func (s *VolumeService) CreateVolume(p *CreateVolumeParams) (*CreateVolumeResponse, error) {
	resp, err := s.cs.newRequest("createVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVolumeResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CreateVolumeResponse struct {
	JobID                      string `json:"jobid,omitempty"`
	Account                    string `json:"account,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Chaininfo                  string `json:"chaininfo,omitempty"`
	Created                    string `json:"created,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Deviceid                   int64  `json:"deviceid,omitempty"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate,omitempty"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate,omitempty"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate,omitempty"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Id                         string `json:"id,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Isodisplaytext             string `json:"isodisplaytext,omitempty"`
	Isoid                      string `json:"isoid,omitempty"`
	Isoname                    string `json:"isoname,omitempty"`
	Maxiops                    int64  `json:"maxiops,omitempty"`
	Miniops                    int64  `json:"miniops,omitempty"`
	Name                       string `json:"name,omitempty"`
	Path                       string `json:"path,omitempty"`
	Project                    string `json:"project,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Provisioningtype           string `json:"provisioningtype,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Size                       int64  `json:"size,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	State                      string `json:"state,omitempty"`
	Status                     string `json:"status,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Tags                       []struct {
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
	Type                string `json:"type,omitempty"`
	Virtualmachineid    string `json:"virtualmachineid,omitempty"`
	Vmdisplayname       string `json:"vmdisplayname,omitempty"`
	Vmname              string `json:"vmname,omitempty"`
	Vmstate             string `json:"vmstate,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type DeleteVolumeParams struct {
	p map[string]interface{}
}

func (p *DeleteVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewDeleteVolumeParams(id string) *DeleteVolumeParams {
	p := &DeleteVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a detached disk volume.
func (s *VolumeService) DeleteVolume(p *DeleteVolumeParams) (*DeleteVolumeResponse, error) {
	resp, err := s.cs.newRequest("deleteVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVolumeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteVolumeResponse struct {
	Displaytext string `json:"displaytext,omitempty"`
	Success     string `json:"success,omitempty"`
}

type ListVolumesParams struct {
	p map[string]interface{}
}

func (p *ListVolumesParams) toURLValues() url.Values {
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
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVolumesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}
func (p *ListVolumesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListVolumesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListVolumesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListVolumesParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
	return
}

func (p *ListVolumesParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

func (p *ListVolumesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListVolumesParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewListVolumesParams() *ListVolumesParams {
	p := &ListVolumesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumeID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListVolumesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVolumes(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.Volumes[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.Volumes {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumeByName(name string, opts ...OptionFunc) (*Volume, int, error) {
	id, count, err := s.GetVolumeID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetVolumeByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VolumeService) GetVolumeByID(id string, opts ...OptionFunc) (*Volume, int, error) {
	p := &ListVolumesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range opts {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListVolumes(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.Volumes[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for Volume UUID: %s!", id)
}

// Lists all volumes.
func (s *VolumeService) ListVolumes(p *ListVolumesParams) (*ListVolumesResponse, error) {
	resp, err := s.cs.newRequest("listVolumes", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVolumesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVolumesResponse struct {
	Count   int       `json:"count"`
	Volumes []*Volume `json:"volume"`
}

type Volume struct {
	Account                    string `json:"account,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Chaininfo                  string `json:"chaininfo,omitempty"`
	Created                    string `json:"created,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Deviceid                   int64  `json:"deviceid,omitempty"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate,omitempty"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate,omitempty"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate,omitempty"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Id                         string `json:"id,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Isodisplaytext             string `json:"isodisplaytext,omitempty"`
	Isoid                      string `json:"isoid,omitempty"`
	Isoname                    string `json:"isoname,omitempty"`
	Maxiops                    int64  `json:"maxiops,omitempty"`
	Miniops                    int64  `json:"miniops,omitempty"`
	Name                       string `json:"name,omitempty"`
	Path                       string `json:"path,omitempty"`
	Project                    string `json:"project,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Provisioningtype           string `json:"provisioningtype,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Size                       int64  `json:"size,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	State                      string `json:"state,omitempty"`
	Status                     string `json:"status,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Tags                       []struct {
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
	Type                string `json:"type,omitempty"`
	Virtualmachineid    string `json:"virtualmachineid,omitempty"`
	Vmdisplayname       string `json:"vmdisplayname,omitempty"`
	Vmname              string `json:"vmname,omitempty"`
	Vmstate             string `json:"vmstate,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}

type ResizeVolumeParams struct {
	p map[string]interface{}
}

func (p *ResizeVolumeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["size"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("size", vv)
	}
	return u
}

func (p *ResizeVolumeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ResizeVolumeParams) SetSize(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["size"] = v
	return
}

// You should always use this function to get a new ResizeVolumeParams instance,
// as then you are sure you have configured all required params
func (s *VolumeService) NewResizeVolumeParams(id string, size int64) *ResizeVolumeParams {
	p := &ResizeVolumeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["size"] = size
	return p
}

// Resizes a volume
func (s *VolumeService) ResizeVolume(p *ResizeVolumeParams) (*ResizeVolumeResponse, error) {
	resp, err := s.cs.newRequest("resizeVolume", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResizeVolumeResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type ResizeVolumeResponse struct {
	JobID                      string `json:"jobid,omitempty"`
	Account                    string `json:"account,omitempty"`
	Attached                   string `json:"attached,omitempty"`
	Chaininfo                  string `json:"chaininfo,omitempty"`
	Created                    string `json:"created,omitempty"`
	Destroyed                  bool   `json:"destroyed,omitempty"`
	Deviceid                   int64  `json:"deviceid,omitempty"`
	DiskBytesReadRate          int64  `json:"diskBytesReadRate,omitempty"`
	DiskBytesWriteRate         int64  `json:"diskBytesWriteRate,omitempty"`
	DiskIopsReadRate           int64  `json:"diskIopsReadRate,omitempty"`
	DiskIopsWriteRate          int64  `json:"diskIopsWriteRate,omitempty"`
	Diskofferingdisplaytext    string `json:"diskofferingdisplaytext,omitempty"`
	Diskofferingid             string `json:"diskofferingid,omitempty"`
	Diskofferingname           string `json:"diskofferingname,omitempty"`
	Displayvolume              bool   `json:"displayvolume,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	Domainid                   string `json:"domainid,omitempty"`
	Hypervisor                 string `json:"hypervisor,omitempty"`
	Id                         string `json:"id,omitempty"`
	Isextractable              bool   `json:"isextractable,omitempty"`
	Isodisplaytext             string `json:"isodisplaytext,omitempty"`
	Isoid                      string `json:"isoid,omitempty"`
	Isoname                    string `json:"isoname,omitempty"`
	Maxiops                    int64  `json:"maxiops,omitempty"`
	Miniops                    int64  `json:"miniops,omitempty"`
	Name                       string `json:"name,omitempty"`
	Path                       string `json:"path,omitempty"`
	Project                    string `json:"project,omitempty"`
	Projectid                  string `json:"projectid,omitempty"`
	Provisioningtype           string `json:"provisioningtype,omitempty"`
	Quiescevm                  bool   `json:"quiescevm,omitempty"`
	Serviceofferingdisplaytext string `json:"serviceofferingdisplaytext,omitempty"`
	Serviceofferingid          string `json:"serviceofferingid,omitempty"`
	Serviceofferingname        string `json:"serviceofferingname,omitempty"`
	Size                       int64  `json:"size,omitempty"`
	Snapshotid                 string `json:"snapshotid,omitempty"`
	State                      string `json:"state,omitempty"`
	Status                     string `json:"status,omitempty"`
	Storage                    string `json:"storage,omitempty"`
	Storageid                  string `json:"storageid,omitempty"`
	Storagetype                string `json:"storagetype,omitempty"`
	Tags                       []struct {
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
	Type                string `json:"type,omitempty"`
	Virtualmachineid    string `json:"virtualmachineid,omitempty"`
	Vmdisplayname       string `json:"vmdisplayname,omitempty"`
	Vmname              string `json:"vmname,omitempty"`
	Vmstate             string `json:"vmstate,omitempty"`
	Zoneid              string `json:"zoneid,omitempty"`
	Zonename            string `json:"zonename,omitempty"`
}
