package cloudstack

type Portforwardingrule struct {
	Cidrlist                  NullString `json:"cidrlist"`
	Fordisplay                NullBool   `json:"fordisplay"`
	Id                        ID         `json:"id"`
	Ipaddress                 NullString `json:"ipaddress"`
	Ipaddressid               ID         `json:"ipaddressid"`
	Networkid                 ID         `json:"networkid"`
	Privateendport            NullString `json:"privateendport"`
	Privateport               NullString `json:"privateport"`
	Protocol                  NullString `json:"protocol"`
	Publicendport             NullString `json:"publicendport"`
	Publicport                NullString `json:"publicport"`
	State                     NullString `json:"state"`
	Tag                       []Tag      `json:"tags"`
	Virtualmachinedisplayname NullString `json:"virtualmachinedisplayname"`
	Virtualmachineid          ID         `json:"virtualmachineid"`
	Virtualmachinename        NullString `json:"virtualmachinename"`
	Vmguestip                 NullString `json:"vmguestip"`
}
