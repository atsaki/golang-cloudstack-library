package cloudstack


// DeleteEgressFirewallRule represents the paramter of DeleteEgressFirewallRule
type DeleteEgressFirewallRuleParameter struct {
	// the ID of the firewall rule
	Id ID
}

func NewDeleteEgressFirewallRuleParameter(id string) (p *DeleteEgressFirewallRuleParameter) {
	p = new(DeleteEgressFirewallRuleParameter)
	p.Id.Set(id)
	return p
}

// Deletes an ggress firewall rule
func (c *Client) DeleteEgressFirewallRule(p *DeleteEgressFirewallRuleParameter) (*Result, error) {
	obj, err := c.Request("deleteEgressFirewallRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// DeleteFirewallRule represents the paramter of DeleteFirewallRule
type DeleteFirewallRuleParameter struct {
	// the ID of the firewall rule
	Id ID
}

func NewDeleteFirewallRuleParameter(id string) (p *DeleteFirewallRuleParameter) {
	p = new(DeleteFirewallRuleParameter)
	p.Id.Set(id)
	return p
}

// Deletes a firewall rule
func (c *Client) DeleteFirewallRule(p *DeleteFirewallRuleParameter) (*Result, error) {
	obj, err := c.Request("deleteFirewallRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// UpdatePortForwardingRule represents the paramter of UpdatePortForwardingRule
type UpdatePortForwardingRuleParameter struct {
	// an optional field, in case you want to set a custom id to the resource.
	// Allowed to Root Admins only
	CustomId ID
	// an optional field, whether to the display the rule to the end user or not
	ForDisplay NullBool
	// the ID of the port forwarding rule
	Id ID
	// the IP address id of the port forwarding rule
	IpAddressId ID
	// the private IP address of the port forwarding rule
	PrivateIp NullString
	// the private port of the port forwarding rule
	PrivatePort NullString
	// the protocol for the port fowarding rule. Valid values are TCP or UDP.
	Protocol NullString
	// the public port of the port forwarding rule
	PublicPort NullString
	// the ID of the virtual machine for the port forwarding rule
	VirtualMachineId ID
}

func NewUpdatePortForwardingRuleParameter(id string) (p *UpdatePortForwardingRuleParameter) {
	p = new(UpdatePortForwardingRuleParameter)
	p.Id.Set(id)
	return p
}

// Updates a port forwarding rule.  Only the private port and the virtual
// machine can be updated.
func (c *Client) UpdatePortForwardingRule(p *UpdatePortForwardingRuleParameter) (*PortForwardingRule, error) {
	obj, err := c.Request("updatePortForwardingRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*PortForwardingRule), err
}


// ListPortForwardingRules represents the paramter of ListPortForwardingRules
type ListPortForwardingRulesParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// list only resources belonging to the domain specified
	DomainId ID
	// list resources by display flag; only ROOT admin is eligible to pass this
	// parameter
	ForDisplay NullBool
	// Lists rule with the specified ID.
	Id ID
	// the id of IP address of the port forwarding services
	IpAddressId ID
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll NullBool
	// list port forwarding rules for ceratin network
	NetworkId ID
	Page      NullNumber
	PageSize  NullNumber
	// list objects by project
	ProjectId ID
	// List resources by tags (key/value pairs)
	Tags map[string]string
}

func NewListPortForwardingRulesParameter() (p *ListPortForwardingRulesParameter) {
	p = new(ListPortForwardingRulesParameter)
	return p
}

// Lists all port forwarding rules for an IP address.
func (c *Client) ListPortForwardingRules(p *ListPortForwardingRulesParameter) ([]*PortForwardingRule, error) {
	obj, err := c.Request("listPortForwardingRules", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*PortForwardingRule), err
}


// CreatePortForwardingRule represents the paramter of CreatePortForwardingRule
type CreatePortForwardingRuleParameter struct {
	// the cidr list to forward traffic from
	CidrList []string
	// an optional field, whether to the display the rule to the end user or not
	ForDisplay NullBool
	// the IP address id of the port forwarding rule
	IpAddressId ID
	// The network of the vm the Port Forwarding rule will be created for. Required
	// when public Ip address is not associated with any Guest network yet (VPC
	// case)
	NetworkId ID
	// if true, firewall rule for source/end pubic port is automatically created; if
	// false - firewall rule has to be created explicitely. If not specified 1)
	// defaulted to false when PF rule is being created for VPC guest network 2) in
	// all other cases defaulted to true
	OpenFirewall NullBool
	// the ending port of port forwarding rule's private port range
	PrivateEndPort NullNumber
	// the starting port of port forwarding rule's private port range
	PrivatePort NullNumber
	// the protocol for the port fowarding rule. Valid values are TCP or UDP.
	Protocol NullString
	// the ending port of port forwarding rule's private port range
	PublicEndPort NullNumber
	// the starting port of port forwarding rule's public port range
	PublicPort NullNumber
	// the ID of the virtual machine for the port forwarding rule
	VirtualMachineId ID
	// VM guest nic Secondary ip address for the port forwarding rule
	VmGuestIp NullString
}

func NewCreatePortForwardingRuleParameter(ipaddressid string, privateport int, protocol string, publicport int, virtualmachineid string) (p *CreatePortForwardingRuleParameter) {
	p = new(CreatePortForwardingRuleParameter)
	p.IpAddressId.Set(ipaddressid)
	p.PrivatePort.Set(privateport)
	p.Protocol.Set(protocol)
	p.PublicPort.Set(publicport)
	p.VirtualMachineId.Set(virtualmachineid)
	return p
}

// Creates a port forwarding rule
func (c *Client) CreatePortForwardingRule(p *CreatePortForwardingRuleParameter) (*PortForwardingRule, error) {
	obj, err := c.Request("createPortForwardingRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*PortForwardingRule), err
}


// CreateEgressFirewallRule represents the paramter of CreateEgressFirewallRule
type CreateEgressFirewallRuleParameter struct {
	// the cidr list to forward traffic from
	CidrList []string
	// the ending port of firewall rule
	EndPort NullNumber
	// an optional field, whether to the display the rule to the end user or not
	ForDisplay NullBool
	// error code for this icmp message
	IcmpCode NullNumber
	// type of the icmp message being sent
	IcmpType NullNumber
	// the network id of the port forwarding rule
	NetworkId ID
	// the protocol for the firewall rule. Valid values are TCP/UDP/ICMP.
	Protocol NullString
	// the starting port of firewall rule
	StartPort NullNumber
	// type of firewallrule: system/user
	Type NullString
}

func NewCreateEgressFirewallRuleParameter(networkid string, protocol string) (p *CreateEgressFirewallRuleParameter) {
	p = new(CreateEgressFirewallRuleParameter)
	p.NetworkId.Set(networkid)
	p.Protocol.Set(protocol)
	return p
}

// Creates a egress firewall rule for a given network
func (c *Client) CreateEgressFirewallRule(p *CreateEgressFirewallRuleParameter) (*FirewallRule, error) {
	obj, err := c.Request("createEgressFirewallRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*FirewallRule), err
}


// UpdateEgressFirewallRule represents the paramter of UpdateEgressFirewallRule
type UpdateEgressFirewallRuleParameter struct {
	// an optional field, in case you want to set a custom id to the resource.
	// Allowed to Root Admins only
	CustomId ID
	// an optional field, whether to the display the rule to the end user or not
	ForDisplay NullBool
	// the ID of the egress firewall rule
	Id ID
}

func NewUpdateEgressFirewallRuleParameter(id string) (p *UpdateEgressFirewallRuleParameter) {
	p = new(UpdateEgressFirewallRuleParameter)
	p.Id.Set(id)
	return p
}

// Updates egress firewall rule
func (c *Client) UpdateEgressFirewallRule(p *UpdateEgressFirewallRuleParameter) (*FirewallRule, error) {
	obj, err := c.Request("updateEgressFirewallRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*FirewallRule), err
}


// DeletePortForwardingRule represents the paramter of DeletePortForwardingRule
type DeletePortForwardingRuleParameter struct {
	// the ID of the port forwarding rule
	Id ID
}

func NewDeletePortForwardingRuleParameter(id string) (p *DeletePortForwardingRuleParameter) {
	p = new(DeletePortForwardingRuleParameter)
	p.Id.Set(id)
	return p
}

// Deletes a port forwarding rule
func (c *Client) DeletePortForwardingRule(p *DeletePortForwardingRuleParameter) (*Result, error) {
	obj, err := c.Request("deletePortForwardingRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// ListFirewallRules represents the paramter of ListFirewallRules
type ListFirewallRulesParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// list only resources belonging to the domain specified
	DomainId ID
	// list resources by display flag; only ROOT admin is eligible to pass this
	// parameter
	ForDisplay NullBool
	// Lists rule with the specified ID.
	Id ID
	// the id of IP address of the firwall services
	IpAddressId ID
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll NullBool
	// list firewall rules for ceratin network
	NetworkId ID
	Page      NullNumber
	PageSize  NullNumber
	// list objects by project
	ProjectId ID
	// List resources by tags (key/value pairs)
	Tags map[string]string
}

func NewListFirewallRulesParameter() (p *ListFirewallRulesParameter) {
	p = new(ListFirewallRulesParameter)
	return p
}

// Lists all firewall rules for an IP address.
func (c *Client) ListFirewallRules(p *ListFirewallRulesParameter) ([]*FirewallRule, error) {
	obj, err := c.Request("listFirewallRules", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*FirewallRule), err
}


// UpdateFirewallRule represents the paramter of UpdateFirewallRule
type UpdateFirewallRuleParameter struct {
	// an optional field, in case you want to set a custom id to the resource.
	// Allowed to Root Admins only
	CustomId ID
	// an optional field, whether to the display the rule to the end user or not
	ForDisplay NullBool
	// the ID of the firewall rule
	Id ID
}

func NewUpdateFirewallRuleParameter(id string) (p *UpdateFirewallRuleParameter) {
	p = new(UpdateFirewallRuleParameter)
	p.Id.Set(id)
	return p
}

// Updates firewall rule
func (c *Client) UpdateFirewallRule(p *UpdateFirewallRuleParameter) (*FirewallRule, error) {
	obj, err := c.Request("updateFirewallRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*FirewallRule), err
}


// CreateFirewallRule represents the paramter of CreateFirewallRule
type CreateFirewallRuleParameter struct {
	// the cidr list to forward traffic from
	CidrList []string
	// the ending port of firewall rule
	EndPort NullNumber
	// an optional field, whether to the display the rule to the end user or not
	ForDisplay NullBool
	// error code for this icmp message
	IcmpCode NullNumber
	// type of the icmp message being sent
	IcmpType NullNumber
	// the IP address id of the port forwarding rule
	IpAddressId ID
	// the protocol for the firewall rule. Valid values are TCP/UDP/ICMP.
	Protocol NullString
	// the starting port of firewall rule
	StartPort NullNumber
	// type of firewallrule: system/user
	Type NullString
}

func NewCreateFirewallRuleParameter(ipaddressid string, protocol string) (p *CreateFirewallRuleParameter) {
	p = new(CreateFirewallRuleParameter)
	p.IpAddressId.Set(ipaddressid)
	p.Protocol.Set(protocol)
	return p
}

// Creates a firewall rule for a given ip address
func (c *Client) CreateFirewallRule(p *CreateFirewallRuleParameter) (*FirewallRule, error) {
	obj, err := c.Request("createFirewallRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*FirewallRule), err
}


// ListEgressFirewallRules represents the paramter of ListEgressFirewallRules
type ListEgressFirewallRulesParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// list only resources belonging to the domain specified
	DomainId ID
	// list resources by display flag; only ROOT admin is eligible to pass this
	// parameter
	ForDisplay NullBool
	// Lists rule with the specified ID.
	Id ID
	// the id of IP address of the firwall services
	IpAddressId ID
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll NullBool
	// the id network network for the egress firwall services
	NetworkId ID
	Page      NullNumber
	PageSize  NullNumber
	// list objects by project
	ProjectId ID
	// List resources by tags (key/value pairs)
	Tags map[string]string
}

func NewListEgressFirewallRulesParameter() (p *ListEgressFirewallRulesParameter) {
	p = new(ListEgressFirewallRulesParameter)
	return p
}

// Lists all egress firewall rules for network id.
func (c *Client) ListEgressFirewallRules(p *ListEgressFirewallRulesParameter) ([]*FirewallRule, error) {
	obj, err := c.Request("listEgressFirewallRules", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*FirewallRule), err
}

