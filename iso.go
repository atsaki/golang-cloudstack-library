package cloudstack

type Iso struct {
	Account               NullString `json:"account"`
	Accountid             ID         `json:"accountid"`
	Bootable              NullBool   `json:"bootable"`
	Checksum              NullString `json:"checksum"`
	Created               NullString `json:"created"`
	CrossZones            NullBool   `json:"crosszones"`
	Details               NullString `json:"details"`
	Displaytext           NullString `json:"displaytext"`
	Domain                NullString `json:"domain"`
	Domainid              ID         `json:"domainid"`
	Format                NullString `json:"format"`
	Hostid                ID         `json:"hostid"`
	Hostname              NullString `json:"hostname"`
	Hypervisor            NullString `json:"hypervisor"`
	Id                    ID         `json:"id"`
	Isdynamicallyscalable NullBool   `json:"isdynamicallyscalable"`
	Isextractable         NullBool   `json:"isextractable"`
	Isfeatured            NullBool   `json:"isfeatured"`
	Ispublic              NullBool   `json:"ispublic"`
	Isready               NullBool   `json:"isready"`
	Name                  NullString `json:"name"`
	Ostypeid              ID         `json:"ostypeid"`
	Ostypename            NullString `json:"ostypename"`
	Passwordenabled       NullBool   `json:"passwordenabled"`
	Project               NullString `json:"project"`
	Projectid             ID         `json:"projectid"`
	Removed               NullString `json:"removed"`
	Size                  NullInt64  `json:"size"`
	Sourcetemplateid      ID         `json:"sourcetemplateid"`
	Sshkeyenabled         NullBool   `json:"sshkeyenabled"`
	Status                NullString `json:"status"`
	Tag                   []Tag      `json:"tags"`
	Templatetag           NullString `json:"templatetag"`
	Templatetype          NullString `json:"templatetype"`
	Zoneid                ID         `json:"zoneid"`
	Zonename              NullString `json:"zonename"`
}
