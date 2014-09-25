package cloudstack

type Firewallrule struct {
	Cidrlist    NullString `json:"cidrlist"`
	Endport     NullString `json:"endport"`
	Fordisplay  NullBool   `json:"fordisplay"`
	Icmpcode    NullInt64  `json:"icmpcode"`
	Icmptype    NullInt64  `json:"icmptype"`
	Id          ID         `json:"id"`
	Ipaddress   NullString `json:"ipaddress"`
	Ipaddressid ID         `json:"ipaddressid"`
	Networkid   ID         `json:"networkid"`
	Protocol    NullString `json:"protocol"`
	Startport   NullString `json:"startport"`
	State       NullString `json:"state"`
	Tag         []Tag      `json:"tags"`
}
