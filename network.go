package cloudstack

type Network struct {
	Account                     NullString       `json:"account"`
	Aclid                       ID               `json:"aclid"`
	Acltype                     NullString       `json:"acltype"`
	Broadcastdomaintype         NullString       `json:"broadcastdomaintype"`
	Broadcasturi                NullString       `json:"broadcasturi"`
	Canusefordeploy             NullBool         `json:"canusefordeploy"`
	Cidr                        NullString       `json:"cidr"`
	Displaynetwork              NullBool         `json:"displaynetwork"`
	Displaytext                 NullString       `json:"displaytext"`
	Dns1                        NullString       `json:"dns1"`
	Dns2                        NullString       `json:"dns2"`
	Domain                      NullString       `json:"domain"`
	Domainid                    ID               `json:"domainid"`
	Gateway                     NullString       `json:"gateway"`
	Id                          ID               `json:"id"`
	Ip6cidr                     NullString       `json:"ip6cidr"`
	Ip6gateway                  NullString       `json:"ip6gateway"`
	Isdefault                   NullBool         `json:"isdefault"`
	Ispersistent                NullBool         `json:"ispersistent"`
	Issystem                    NullBool         `json:"issystem"`
	Name                        NullString       `json:"name"`
	Netmask                     NullString       `json:"netmask"`
	Networkcidr                 NullString       `json:"networkcidr"`
	Networkdomain               NullString       `json:"networkdomain"`
	Networkofferingavailability NullString       `json:"networkofferingavailability"`
	Networkofferingconservemode NullBool         `json:"networkofferingconservemode"`
	Networkofferingdisplaytext  NullString       `json:"networkofferingdisplaytext"`
	Networkofferingid           ID               `json:"networkofferingid"`
	Networkofferingname         NullString       `json:"networkofferingname"`
	Physicalnetworkid           ID               `json:"physicalnetworkid"`
	Project                     NullString       `json:"project"`
	Projectid                   ID               `json:"projectid"`
	Related                     NullString       `json:"related"`
	Reservediprange             NullString       `json:"reservediprange"`
	Restartrequired             NullBool         `json:"restartrequired"`
	Service                     []Networkservice `json:"service"`
	Specifyipranges             NullBool         `json:"specifyipranges"`
	State                       NullString       `json:"state"`
	Strechedl2subnet            NullBool         `json:"strechedl2subnet"`
	Subdomainaccess             NullBool         `json:"subdomainaccess"`
	Tag                         []Tag            `json:"tags"`
	Traffictype                 NullString       `json:"traffictype"`
	Type                        NullString       `json:"type"`
	Vlan                        NullString       `json:"vlan"`
	Vpcid                       ID               `json:"vpcid"`
	Zoneid                      ID               `json:"zoneid"`
	Zonename                    NullString       `json:"zonename"`
	Zonesnetworkspans           []NullString     `json:"zonesnetworkspans"`
}

type Networkservice struct {
	Capability []Networkservicecapability `json:"capability"`
	Name       NullString                 `json:"name"`
	Provider   []Networkserviceprovider   `json:"provider"`
}

type Networkservicecapability struct {
	Canchooseservicecapability NullBool   `json:"canchooseservicecapability"`
	Name                       NullString `json:"name"`
	Value                      NullString `json:"value"`
}

type Networkserviceprovider struct {
	Canenableindividualservice   NullBool     `json:"canenableindividualservice"`
	Destinationphysicalnetworkid ID           `json:"destinationphysicalnetworkid"`
	Id                           ID           `json:"id"`
	Name                         NullString   `json:"name"`
	Physicalnetworkid            ID           `json:"physicalnetworkid"`
	Servicelist                  []NullString `json:"servicelist"`
	State                        NullString   `json:"state"`
}
