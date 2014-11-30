package cloudstack

type PortForwardingRule struct {
	ResourceBase
	// the cidr list to forward traffic from
	CidrList NullString `json:"cidrlist"`
	// is firewall for display to the regular user
	ForDisplay NullBool `json:"fordisplay"`
	// the ID of the port forwarding rule
	Id ID `json:"id"`
	// the public ip address for the port forwarding rule
	IpAddress NullString `json:"ipaddress"`
	// the public ip address id for the port forwarding rule
	IpAddressId ID `json:"ipaddressid"`
	// the id of the guest network the port forwarding rule belongs to
	NetworkId ID `json:"networkid"`
	// the ending port of port forwarding rule's private port range
	PrivateEndPort NullString `json:"privateendport"`
	// the starting port of port forwarding rule's private port range
	PrivatePort NullString `json:"privateport"`
	// the protocol of the port forwarding rule
	Protocol NullString `json:"protocol"`
	// the ending port of port forwarding rule's private port range
	PublicEndPort NullString `json:"publicendport"`
	// the starting port of port forwarding rule's public port range
	PublicPort NullString `json:"publicport"`
	// the state of the rule
	State NullString `json:"state"`
	// the list of resource tags associated with the rule
	Tags []Tag `json:"tags"`
	// the VM display name for the port forwarding rule
	VirtualMachineDisplayName NullString `json:"virtualmachinedisplayname"`
	// the VM ID for the port forwarding rule
	VirtualMachineId ID `json:"virtualmachineid"`
	// the VM name for the port forwarding rule
	VirtualMachineName NullString `json:"virtualmachinename"`
	// the vm ip address for the port forwarding rule
	VmGuestIp NullString `json:"vmguestip"`
}
