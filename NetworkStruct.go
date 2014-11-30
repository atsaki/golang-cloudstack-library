package cloudstack

type Network struct {
	ResourceBase
	// the owner of the network
	Account NullString `json:"account"`
	// ACL Id associated with the VPC network
	AclId ID `json:"aclid"`
	// acl type - access type to the network
	AclType NullString `json:"acltype"`
	// Broadcast domain type of the network
	BroadcastDomainType NullString `json:"broadcastdomaintype"`
	// broadcast uri of the network. This parameter is visible to ROOT admins only
	BroadcastUri NullString `json:"broadcasturi"`
	// list networks available for vm deployment
	CanUseForDeploy NullBool `json:"canusefordeploy"`
	// Cloudstack managed address space, all CloudStack managed VMs get IP address
	// from CIDR
	Cidr NullString `json:"cidr"`
	// an optional field, whether to the display the network to the end user or not.
	DisplayNetwork NullBool `json:"displaynetwork"`
	// the displaytext of the network
	DisplayText NullString `json:"displaytext"`
	// the first DNS for the network
	Dns1 NullString `json:"dns1"`
	// the second DNS for the network
	Dns2 NullString `json:"dns2"`
	// the domain name of the network owner
	Domain NullString `json:"domain"`
	// the domain id of the network owner
	DomainId ID `json:"domainid"`
	// the network's gateway
	Gateway NullString `json:"gateway"`
	// the id of the network
	Id ID `json:"id"`
	// the cidr of IPv6 network
	Ip6Cidr NullString `json:"ip6cidr"`
	// the gateway of IPv6 network
	Ip6Gateway NullString `json:"ip6gateway"`
	// true if network is default, false otherwise
	IsDefault NullBool `json:"isdefault"`
	// list networks that are persistent
	IsPersistent NullBool `json:"ispersistent"`
	// true if network is system, false otherwise
	IsSystem NullBool `json:"issystem"`
	// the name of the network
	Name NullString `json:"name"`
	// the network's netmask
	Netmask NullString `json:"netmask"`
	// the network CIDR of the guest network configured with IP reservation. It is
	// the summation of CIDR and RESERVED_IP_RANGE
	NetworkCidr NullString `json:"networkcidr"`
	// the network domain
	NetworkDomain NullString `json:"networkdomain"`
	// availability of the network offering the network is created from
	NetworkOfferingAvailability NullString `json:"networkofferingavailability"`
	// true if network offering is ip conserve mode enabled
	NetworkOfferingconserveMode NullBool `json:"networkofferingconservemode"`
	// display text of the network offering the network is created from
	NetworkOfferingDisplayText NullString `json:"networkofferingdisplaytext"`
	// network offering id the network is created from
	NetworkOfferingId ID `json:"networkofferingid"`
	// name of the network offering the network is created from
	NetworkOfferingName NullString `json:"networkofferingname"`
	// the physical network id
	PhysicalNetworkId ID `json:"physicalnetworkid"`
	// the project name of the address
	Project NullString `json:"project"`
	// the project id of the ipaddress
	ProjectId ID `json:"projectid"`
	// related to what other network configuration
	Related NullString `json:"related"`
	// the network's IP range not to be used by CloudStack guest VMs and can be used
	// for non CloudStack purposes
	ReservedIpRange NullString `json:"reservediprange"`
	// true network requires restart
	ReStartRequired NullBool `json:"restartrequired"`
	// the list of services
	Service []NetworkService `json:"service"`
	// true if network supports specifying ip ranges, false otherwise
	SpecifyIpRanges NullBool `json:"specifyipranges"`
	// state of the network
	State NullString `json:"state"`
	// true if network can span multiple zones
	StrechedL2Subnet NullBool `json:"strechedl2subnet"`
	// true if users from subdomains can access the domain level network
	SubDomainAccess NullBool `json:"subdomainaccess"`
	// the list of resource tags associated with network
	Tags []Tag `json:"tags"`
	// the traffic type of the network
	TrafficType NullString `json:"traffictype"`
	// the type of the network
	Type NullString `json:"type"`
	// The vlan of the network. This parameter is visible to ROOT admins only
	Vlan NullString `json:"vlan"`
	// VPC the network belongs to
	VpcId ID `json:"vpcid"`
	// zone id of the network
	ZoneId ID `json:"zoneid"`
	// the name of the zone the network belongs to
	ZoneName NullString `json:"zonename"`
	// If a network is enabled for 'streched l2 subnet' then represents zones on
	// which network currently spans
	ZonesNetworkspans []NullString `json:"zonesnetworkspans"`
}
