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
