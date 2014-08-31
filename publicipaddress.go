package cloudstack

type Publicipaddress struct {
	Account                   NullString `json:"account"`
	Allocated                 NullString `json:"allocated"`
	Associatednetworkid       ID         `json:"associatednetworkid"`
	Associatednetworkname     NullString `json:"associatednetworkname"`
	Domain                    NullString `json:"domain"`
	Domainid                  ID         `json:"domainid"`
	Fordisplay                NullBool   `json:"fordisplay"`
	Forvirtualnetwork         NullBool   `json:"forvirtualnetwork"`
	Id                        ID         `json:"id"`
	Ipaddress                 NullString `json:"ipaddress"`
	Isportable                NullBool   `json:"isportable"`
	Issourcenat               NullBool   `json:"issourcenat"`
	Isstaticnat               NullBool   `json:"isstaticnat"`
	Issystem                  NullBool   `json:"issystem"`
	Networkid                 ID         `json:"networkid"`
	Physicalnetworkid         ID         `json:"physicalnetworkid"`
	Project                   NullString `json:"project"`
	Projectid                 ID         `json:"projectid"`
	Purpose                   NullString `json:"purpose"`
	State                     NullString `json:"state"`
	Tag                       []Tag      `json:"tags"`
	Virtualmachinedisplayname NullString `json:"virtualmachinedisplayname"`
	Virtualmachineid          ID         `json:"virtualmachineid"`
	Virtualmachinename        NullString `json:"virtualmachinename"`
	Vlanid                    ID         `json:"vlanid"`
	Vlanname                  NullString `json:"vlanname"`
	Vmipaddress               NullString `json:"vmipaddress"`
	Vpcid                     ID         `json:"vpcid"`
	Zoneid                    ID         `json:"zoneid"`
	Zonename                  NullString `json:"zonename"`
}
