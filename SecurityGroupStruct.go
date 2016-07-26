package cloudstack

type SecurityGroup struct {
	ResourceBase
	// the account owning the security group
	Account NullString `json:"account"`
	// the description of the security group
	Description NullString `json:"description"`
	// the domain name of the security group
	Domain NullString `json:"domain"`
	// the domain ID of the security group
	DomainId ID `json:"domainid"`
	// the list of egress rules associated with the security group
	EgressRule []SecurityGroupEgress `json:"egressrule"`
	// the ID of the security group
	Id ID `json:"id"`
	// the list of ingress rules associated with the security group
	IngressRule []SecurityGroupIngress `json:"ingressrule"`
	// the name of the security group
	Name NullString `json:"name"`
	// the project name of the group
	Project NullString `json:"project"`
	// the project id of the group
	ProjectId ID `json:"projectid"`
	// the list of resource tags associated with the rule
	Tags []Tag `json:"tags"`
}

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
	// the list of resource tags associated with the rule
	Tags []Tag `json:"tags"`
}

type SecurityGroupEgress struct {
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
	// the list of resource tags associated with the rule
	Tags []Tag `json:"tags"`
}
