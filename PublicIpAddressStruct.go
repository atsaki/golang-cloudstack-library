package cloudstack

type PublicIpAddress struct {
	ResourceBase
	// the account the public IP address is associated with
	Account NullString `json:"account"`
	// date the public IP address was acquired
	Allocated NullString `json:"allocated"`
	// the ID of the Network associated with the IP address
	AssociatedNetworkId ID `json:"associatednetworkid"`
	// the name of the Network associated with the IP address
	AssociatedNetworkName NullString `json:"associatednetworkname"`
	// the domain the public IP address is associated with
	Domain NullString `json:"domain"`
	// the domain ID the public IP address is associated with
	DomainId ID `json:"domainid"`
	// is public ip for display to the regular user
	ForDisplay NullBool `json:"fordisplay"`
	// the virtual network for the IP address
	ForVirtualNetwork NullBool `json:"forvirtualnetwork"`
	// public IP address id
	Id ID `json:"id"`
	// public IP address
	IpAddress NullString `json:"ipaddress"`
	// is public IP portable across the zones
	IsPortable NullBool `json:"isportable"`
	// true if the IP address is a source nat address, false otherwise
	IsSourceNat NullBool `json:"issourcenat"`
	// true if this ip is for static nat, false otherwise
	IsStaticNat NullBool `json:"isstaticnat"`
	// true if this ip is system ip (was allocated as a part of deployVm or
	// createLbRule)
	IsSystem NullBool `json:"issystem"`
	// the ID of the Network where ip belongs to
	NetworkId ID `json:"networkid"`
	// the physical network this belongs to
	PhysicalNetworkId ID `json:"physicalnetworkid"`
	// the project name of the address
	Project NullString `json:"project"`
	// the project id of the ipaddress
	ProjectId ID `json:"projectid"`
	// purpose of the IP address. In Acton this value is not null for Ips with
	// isSystem=true, and can have either StaticNat or LB value
	Purpose NullString `json:"purpose"`
	// State of the ip address. Can be: Allocatin, Allocated and Releasing
	State NullString `json:"state"`
	// the list of resource tags associated with ip address
	Tags []Tag `json:"tags"`
	// virutal machine display name the ip address is assigned to (not null only for
	// static nat Ip)
	VirtualMachineDisplayName NullString `json:"virtualmachinedisplayname"`
	// virutal machine id the ip address is assigned to (not null only for static
	// nat Ip)
	VirtualMachineId ID `json:"virtualmachineid"`
	// virutal machine name the ip address is assigned to (not null only for static
	// nat Ip)
	VirtualMachineName NullString `json:"virtualmachinename"`
	// the ID of the VLAN associated with the IP address. This parameter is visible
	// to ROOT admins only
	VlanId ID `json:"vlanid"`
	// the VLAN associated with the IP address
	VlanName NullString `json:"vlanname"`
	// virutal machine (dnat) ip address (not null only for static nat Ip)
	VmIpAddress NullString `json:"vmipaddress"`
	// VPC the ip belongs to
	VpcId ID `json:"vpcid"`
	// the ID of the zone the public IP address belongs to
	ZoneId ID `json:"zoneid"`
	// the name of the zone the public IP address belongs to
	ZoneName NullString `json:"zonename"`
}
