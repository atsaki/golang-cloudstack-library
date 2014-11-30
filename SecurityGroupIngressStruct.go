package cloudstack

type SecurityGroupIngress struct {
	ResourceBase
	// account owning the security group rule
	Account NullString `json:"account"`
	// the CIDR notation for the base IP address of the security group rule
	Cidr NullString `json:"cidr"`
	// the ending IP of the security group rule
	EndPort NullNumber `json:"endport"`
	// the code for the ICMP message response
	IcmpCode NullNumber `json:"icmpcode"`
	// the type of the ICMP message response
	IcmpType NullNumber `json:"icmptype"`
	// the protocol of the security group rule
	Protocol NullString `json:"protocol"`
	// the id of the security group rule
	RuleId ID `json:"ruleid"`
	// security group name
	SecurityGroupName NullString `json:"securitygroupname"`
	// the starting IP of the security group rule
	StartPort NullNumber `json:"startport"`
}
