package cloudstack

type Securitygroup struct {
	Account     NullString    `json:"account"`
	Description NullString    `json:"description"`
	Domain      NullString    `json:"domain"`
	Domainid    ID            `json:"domainid"`
	Egressrule  []Egressrule  `json:"egressrule"`
	Id          ID            `json:"id"`
	Ingressrule []Ingressrule `json:"ingressrule"`
	Name        NullString    `json:"name"`
	Project     NullString    `json:"project"`
	Projectid   ID            `json:"projectid"`
	Tag         []Tag         `json:"tags"`
}

type Egressrule struct {
	Account           NullString `json:"account"`
	Cidr              NullString `json:"cidr"`
	Endport           NullInt64  `json:"endport"`
	Icmpcode          NullInt64  `json:"icmpcode"`
	Icmptype          NullInt64  `json:"icmptype"`
	Protocol          NullString `json:"protocol"`
	Ruleid            ID         `json:"ruleid"`
	Securitygroupname NullString `json:"securitygroupname"`
	Startport         NullInt64  `json:"startport"`
}

type Ingressrule struct {
	Account           NullString `json:"account"`
	Cidr              NullString `json:"cidr"`
	Endport           NullInt64  `json:"endport"`
	Icmpcode          NullInt64  `json:"icmpcode"`
	Icmptype          NullInt64  `json:"icmptype"`
	Protocol          NullString `json:"protocol"`
	Ruleid            ID         `json:"ruleid"`
	Securitygroupname NullString `json:"securitygroupname"`
	Startport         NullInt64  `json:"startport"`
}
