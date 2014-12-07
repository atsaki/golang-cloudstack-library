package cloudstack

type LoadBalancerRule struct {
	ResourceBase
	// the account of the load balancer rule
	Account NullString `json:"account"`
	// the load balancer algorithm (source, roundrobin, leastconn)
	Algorithm NullString `json:"algorithm"`
	// the cidr list to forward traffic from
	CidrList NullString `json:"cidrlist"`
	// the description of the load balancer
	Description NullString `json:"description"`
	// the domain of the load balancer rule
	Domain NullString `json:"domain"`
	// the domain ID of the load balancer rule
	DomainId ID `json:"domainid"`
	// is rule for display to the regular user
	ForDisplay NullBool `json:"fordisplay"`
	// the load balancer rule ID
	Id ID `json:"id"`
	// the name of the load balancer
	Name NullString `json:"name"`
	// the id of the guest network the lb rule belongs to
	NetworkId ID `json:"networkid"`
	// the private port
	PrivatePort NullString `json:"privateport"`
	// the project name of the load balancer
	Project NullString `json:"project"`
	// the project id of the load balancer
	ProjectId ID `json:"projectid"`
	// the protocol of the loadbalanacer rule
	Protocol NullString `json:"protocol"`
	// the public ip address
	PublicIp NullString `json:"publicip"`
	// the public ip address id
	PublicIpId ID `json:"publicipid"`
	// the public port
	PublicPort NullString `json:"publicport"`
	// the state of the rule
	State NullString `json:"state"`
	// the list of resource tags associated with load balancer
	Tags []Tag `json:"tags"`
	// the id of the zone the rule belongs to
	ZoneId ID `json:"zoneid"`
}
