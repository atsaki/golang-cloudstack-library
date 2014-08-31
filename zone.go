package cloudstack

type Zone struct {
	Allocationstate       NullString `json:"allocationstate"`
	Capacity              []Capacity `json:"capacity"`
	Description           NullString `json:"description"`
	Dhcpprovider          NullString `json:"dhcpprovider"`
	Displaytext           NullString `json:"displaytext"`
	Dns1                  NullString `json:"dns1"`
	Dns2                  NullString `json:"dns2"`
	Domain                NullString `json:"domain"`
	Domainid              ID         `json:"domainid"`
	Domainname            NullString `json:"domainname"`
	Guestcidraddress      NullString `json:"guestcidraddress"`
	Id                    ID         `json:"id"`
	Internaldns1          NullString `json:"internaldns1"`
	Internaldns2          NullString `json:"internaldns2"`
	Ip6dns1               NullString `json:"ip6dns1"`
	Ip6dns2               NullString `json:"ip6dns2"`
	Localstorageenabled   NullBool   `json:"localstorageenabled"`
	Name                  NullString `json:"name"`
	Networktype           NullString `json:"networktype"`
	Resourcedetails       NullString `json:"resourcedetails"`
	Securitygroupsenabled NullBool   `json:"securitygroupsenabled"`
	Tag                   []Tag      `json:"tags"`
	Vlan                  NullString `json:"vlan"`
	Zonetoken             NullString `json:"zonetoken"`
}
