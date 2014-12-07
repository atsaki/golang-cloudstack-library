package cloudstack


// ListLoadBalancerRuleInstances represents the paramter of ListLoadBalancerRuleInstances
type ListLoadBalancerRuleInstancesParameter struct {
	// true if listing all virtual machines currently applied to the load balancer
	// rule; default is true
	Applied NullBool
	// the ID of the load balancer rule
	Id ID
	// List by keyword
	Keyword NullString
	// true if lb rule vm ip information to be included; default is false
	LbVmIps  NullBool
	Page     NullNumber
	PageSize NullNumber
}

func NewListLoadBalancerRuleInstancesParameter(id string) (p *ListLoadBalancerRuleInstancesParameter) {
	p = new(ListLoadBalancerRuleInstancesParameter)
	p.Id.Set(id)
	return p
}

// List all virtual machine instances that are assigned to a load balancer rule.
func (c *Client) ListLoadBalancerRuleInstances(p *ListLoadBalancerRuleInstancesParameter) ([]*LoadBalancerRuleInstance, error) {
	obj, err := c.Request("listLoadBalancerRuleInstances", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*LoadBalancerRuleInstance), err
}


// UpdateLoadBalancerRule represents the paramter of UpdateLoadBalancerRule
type UpdateLoadBalancerRuleParameter struct {
	// load balancer algorithm (source, roundrobin, leastconn)
	Algorithm NullString
	// an optional field, in case you want to set a custom id to the resource.
	// Allowed to Root Admins only
	CustomId ID
	// the description of the load balancer rule
	Description NullString
	// an optional field, whether to the display the rule to the end user or not
	ForDisplay NullBool
	// the id of the load balancer rule to update
	Id ID
	// the name of the load balancer rule
	Name NullString
}

func NewUpdateLoadBalancerRuleParameter(id string) (p *UpdateLoadBalancerRuleParameter) {
	p = new(UpdateLoadBalancerRuleParameter)
	p.Id.Set(id)
	return p
}

// Updates load balancer
func (c *Client) UpdateLoadBalancerRule(p *UpdateLoadBalancerRuleParameter) (*LoadBalancerRule, error) {
	obj, err := c.Request("updateLoadBalancerRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*LoadBalancerRule), err
}


// DeleteLoadBalancerRule represents the paramter of DeleteLoadBalancerRule
type DeleteLoadBalancerRuleParameter struct {
	// the ID of the load balancer rule
	Id ID
}

func NewDeleteLoadBalancerRuleParameter(id string) (p *DeleteLoadBalancerRuleParameter) {
	p = new(DeleteLoadBalancerRuleParameter)
	p.Id.Set(id)
	return p
}

// Deletes a load balancer rule.
func (c *Client) DeleteLoadBalancerRule(p *DeleteLoadBalancerRuleParameter) (*Result, error) {
	obj, err := c.Request("deleteLoadBalancerRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// ListLoadBalancerRules represents the paramter of ListLoadBalancerRules
type ListLoadBalancerRulesParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// list only resources belonging to the domain specified
	DomainId ID
	// list resources by display flag; only ROOT admin is eligible to pass this
	// parameter
	ForDisplay NullBool
	// the ID of the load balancer rule
	Id ID
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll NullBool
	// the name of the load balancer rule
	Name NullString
	// list by network id the rule belongs to
	NetworkId ID
	Page      NullNumber
	PageSize  NullNumber
	// list objects by project
	ProjectId ID
	// the public IP address id of the load balancer rule
	PublicIpId ID
	// List resources by tags (key/value pairs)
	Tags map[string]string
	// the ID of the virtual machine of the load balancer rule
	VirtualMachineId ID
	// the availability zone ID
	ZoneId ID
}

func NewListLoadBalancerRulesParameter() (p *ListLoadBalancerRulesParameter) {
	p = new(ListLoadBalancerRulesParameter)
	return p
}

// Lists load balancer rules.
func (c *Client) ListLoadBalancerRules(p *ListLoadBalancerRulesParameter) ([]*LoadBalancerRule, error) {
	obj, err := c.Request("listLoadBalancerRules", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*LoadBalancerRule), err
}


// AssignToLoadBalancerRule represents the paramter of AssignToLoadBalancerRule
type AssignToLoadBalancerRuleParameter struct {
	// the ID of the load balancer rule
	Id ID
	// the list of IDs of the virtual machine that are being assigned to the load
	// balancer rule(i.e. virtualMachineIds=1,2,3)
	VirtualMachineIds []string
	// VM ID and IP map, vmidipmap[0].vmid=1 vmidipmap[0].ip=10.1.1.75
	VmIdIpmap map[string]string
}

func NewAssignToLoadBalancerRuleParameter(id string) (p *AssignToLoadBalancerRuleParameter) {
	p = new(AssignToLoadBalancerRuleParameter)
	p.Id.Set(id)
	return p
}

// Assigns virtual machine or a list of virtual machines to a load balancer
// rule.
func (c *Client) AssignToLoadBalancerRule(p *AssignToLoadBalancerRuleParameter) (*Result, error) {
	obj, err := c.Request("assignToLoadBalancerRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// RemoveFromLoadBalancerRule represents the paramter of RemoveFromLoadBalancerRule
type RemoveFromLoadBalancerRuleParameter struct {
	// The ID of the load balancer rule
	Id ID
	// the list of IDs of the virtual machines that are being removed from the load
	// balancer rule (i.e. virtualMachineIds=1,2,3)
	VirtualMachineIds []string
	// VM ID and IP map, vmidipmap[0].vmid=1 vmidipmap[0].ip=10.1.1.75
	VmIdIpmap map[string]string
}

func NewRemoveFromLoadBalancerRuleParameter(id string) (p *RemoveFromLoadBalancerRuleParameter) {
	p = new(RemoveFromLoadBalancerRuleParameter)
	p.Id.Set(id)
	return p
}

// Removes a virtual machine or a list of virtual machines from a load balancer
// rule.
func (c *Client) RemoveFromLoadBalancerRule(p *RemoveFromLoadBalancerRuleParameter) (*Result, error) {
	obj, err := c.Request("removeFromLoadBalancerRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// CreateLoadBalancerRule represents the paramter of CreateLoadBalancerRule
type CreateLoadBalancerRuleParameter struct {
	// the account associated with the load balancer. Must be used with the domainId
	// parameter.
	Account NullString
	// load balancer algorithm (source, roundrobin, leastconn)
	Algorithm NullString
	// the cidr list to forward traffic from
	CidrList []string
	// the description of the load balancer rule
	Description NullString
	// the domain ID associated with the load balancer
	DomainId ID
	// an optional field, whether to the display the rule to the end user or not
	ForDisplay NullBool
	// name of the load balancer rule
	Name NullString
	// The guest network this rule will be created for. Required when public Ip
	// address is not associated with any Guest network yet (VPC case)
	NetworkId ID
	// if true, firewall rule for source/end pubic port is automatically created; if
	// false - firewall rule has to be created explicitely. If not specified 1)
	// defaulted to false when LB rule is being created for VPC guest network 2) in
	// all other cases defaulted to true
	OpenFirewall NullBool
	// the private port of the private ip address/virtual machine where the network
	// traffic will be load balanced to
	PrivatePort NullNumber
	// The protocol for the LB
	Protocol NullString
	// public ip address id from where the network traffic will be load balanced
	// from
	PublicIpId ID
	// the public port from where the network traffic will be load balanced from
	PublicPort NullNumber
	// zone where the load balancer is going to be created. This parameter is
	// required when LB service provider is ElasticLoadBalancerVm
	ZoneId ID
}

func NewCreateLoadBalancerRuleParameter(algorithm string, name string, privateport int, publicport int) (p *CreateLoadBalancerRuleParameter) {
	p = new(CreateLoadBalancerRuleParameter)
	p.Algorithm.Set(algorithm)
	p.Name.Set(name)
	p.PrivatePort.Set(privateport)
	p.PublicPort.Set(publicport)
	return p
}

// Creates a load balancer rule
func (c *Client) CreateLoadBalancerRule(p *CreateLoadBalancerRuleParameter) (*LoadBalancerRule, error) {
	obj, err := c.Request("createLoadBalancerRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*LoadBalancerRule), err
}

