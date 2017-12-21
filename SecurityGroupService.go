package gokcps

type Egressrule struct {
	Account           string `json:"account,omitempty"`
	Cidr              string `json:"cidr,omitempty"`
	Endport           int    `json:"endport,omitempty"`
	Icmpcode          int    `json:"icmpcode,omitempty"`
	Icmptype          int    `json:"icmptype,omitempty"`
	Protocol          string `json:"protocol,omitempty"`
	Ruleid            string `json:"ruleid,omitempty"`
	Securitygroupname string `json:"securitygroupname,omitempty"`
	Startport         int    `json:"startport,omitempty"`
	Tags              []Tag  `json:"tags,omitempty"`
}

type Ingressrule struct {
	Account           string `json:"account,omitempty"`
	Cidr              string `json:"cidr,omitempty"`
	Endport           int    `json:"endport,omitempty"`
	Icmpcode          int    `json:"icmpcode,omitempty"`
	Icmptype          int    `json:"icmptype,omitempty"`
	Protocol          string `json:"protocol,omitempty"`
	Ruleid            string `json:"ruleid,omitempty"`
	Securitygroupname string `json:"securitygroupname,omitempty"`
	Startport         int    `json:"startport,omitempty"`
	Tags              []Tag  `json:"tags,omitempty"`
}

type Securitygroup struct {
	Account             string        `json:"account,omitempty"`
	Description         string        `json:"description,omitempty"`
	Domain              string        `json:"domain,omitempty"`
	Domainid            string        `json:"domainid,omitempty"`
	Egressrule          []Egressrule  `json:"egressrule,omitempty"`
	Id                  string        `json:"id,omitempty"`
	Ingressrule         []Ingressrule `json:"ingressrule,omitempty"`
	Name                string        `json:"name,omitempty"`
	Project             string        `json:"project,omitempty"`
	Projectid           string        `json:"projectid,omitempty"`
	Tags                []Tag         `json:"tags,omitempty"`
	Virtualmachinecount int           `json:"virtualmachinecount,omitempty"`
	Virtualmachineids   []string      `json:"virtualmachineids,omitempty"`
}
