package cloudstack

type FirewallRule struct {
	ResourceBase
	// the cidr list to forward traffic from
	CidrList NullString `json:"cidrlist"`
	// the ending port of firewall rule's port range
	EndPort NullString `json:"endport"`
	// is rule for display to the regular user
	ForDisplay NullBool `json:"fordisplay"`
	// error code for this icmp message
	IcmpCode NullNumber `json:"icmpcode"`
	// type of the icmp message being sent
	IcmpType NullNumber `json:"icmptype"`
	// the ID of the firewall rule
	Id ID `json:"id"`
	// the public ip address for the firewall rule
	IpAddress NullString `json:"ipaddress"`
	// the public ip address id for the firewall rule
	IpAddressId ID `json:"ipaddressid"`
	// the network id of the firewall rule
	NetworkId ID `json:"networkid"`
	// the protocol of the firewall rule
	Protocol NullString `json:"protocol"`
	// the starting port of firewall rule's port range
	StartPort NullString `json:"startport"`
	// the state of the rule
	State NullString `json:"state"`
	// the list of resource tags associated with the rule
	Tags []Tag `json:"tags"`
}
