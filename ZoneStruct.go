package cloudstack

type Zone struct {
	ResourceBase
	// the allocation state of the cluster
	AllocationState NullString `json:"allocationstate"`
	// the capacity of the Zone
	Capacity []Capacity `json:"capacity"`
	// Zone description
	Description NullString `json:"description"`
	// the dhcp Provider for the Zone
	DhcpProvider NullString `json:"dhcpprovider"`
	// the display text of the zone
	DisplayText NullString `json:"displaytext"`
	// the first DNS for the Zone
	Dns1 NullString `json:"dns1"`
	// the second DNS for the Zone
	Dns2 NullString `json:"dns2"`
	// Network domain name for the networks in the zone
	Domain NullString `json:"domain"`
	// the UUID of the containing domain, null for public zones
	DomainId ID `json:"domainid"`
	// the name of the containing domain, null for public zones
	DomainName NullString `json:"domainname"`
	// the guest CIDR address for the Zone
	GuestCidrAddress NullString `json:"guestcidraddress"`
	// Zone id
	Id ID `json:"id"`
	// the first internal DNS for the Zone
	InternalDns1 NullString `json:"internaldns1"`
	// the second internal DNS for the Zone
	InternalDns2 NullString `json:"internaldns2"`
	// the first IPv6 DNS for the Zone
	Ip6Dns1 NullString `json:"ip6dns1"`
	// the second IPv6 DNS for the Zone
	Ip6Dns2 NullString `json:"ip6dns2"`
	// true if local storage offering enabled, false otherwise
	LocalStorageEnabled NullBool `json:"localstorageenabled"`
	// Zone name
	Name NullString `json:"name"`
	// the network type of the zone; can be Basic or Advanced
	NetworkType NullString `json:"networktype"`
	// Meta data associated with the zone (key/value pairs)
	ResourceDetails NullString `json:"resourcedetails"`
	// true if security groups support is enabled, false otherwise
	SecurityGroupsEnabled NullBool `json:"securitygroupsenabled"`
	// the list of resource tags associated with zone.
	Tags []Tag `json:"tags"`
	// the vlan range of the zone
	Vlan NullString `json:"vlan"`
	// Zone Token
	Zonetoken NullString `json:"zonetoken"`
}
