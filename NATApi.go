package cloudstack


// DeleteIpForwardingRule represents the paramter of DeleteIpForwardingRule
type DeleteIpForwardingRuleParameter struct {
	// the id of the forwarding rule
	Id ID
}

func NewDeleteIpForwardingRuleParameter(id string) (p *DeleteIpForwardingRuleParameter) {
	p = new(DeleteIpForwardingRuleParameter)
	p.Id.Set(id)
	return p
}

// Deletes an ip forwarding rule
func (c *Client) DeleteIpForwardingRule(p *DeleteIpForwardingRuleParameter) (*Result, error) {
	obj, err := c.Request("deleteIpForwardingRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// EnableStaticNat represents the paramter of EnableStaticNat
type EnableStaticNatParameter struct {
	// the public IP address id for which static nat feature is being enabled
	IpAddressId ID
	// The network of the vm the static nat will be enabled for. Required when
	// public Ip address is not associated with any Guest network yet (VPC case)
	NetworkId ID
	// the ID of the virtual machine for enabling static nat feature
	VirtualMachineId ID
	// VM guest nic Secondary ip address for the port forwarding rule
	VmGuestIp NullString
}

func NewEnableStaticNatParameter(ipaddressid string, virtualmachineid string) (p *EnableStaticNatParameter) {
	p = new(EnableStaticNatParameter)
	p.IpAddressId.Set(ipaddressid)
	p.VirtualMachineId.Set(virtualmachineid)
	return p
}

// Enables static nat for given ip address
func (c *Client) EnableStaticNat(p *EnableStaticNatParameter) (*Result, error) {
	obj, err := c.Request("enableStaticNat", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// CreateIpForwardingRule represents the paramter of CreateIpForwardingRule
type CreateIpForwardingRuleParameter struct {
	// the cidr list to forward traffic from
	CidrList []string
	// the end port for the rule
	EndPort NullNumber
	// the public IP address id of the forwarding rule, already associated via
	// associateIp
	IpAddressId ID
	// if true, firewall rule for source/end pubic port is automatically created; if
	// false - firewall rule has to be created explicitely. Has value true by
	// default
	OpenFirewall NullBool
	// the protocol for the rule. Valid values are TCP or UDP.
	Protocol NullString
	// the start port for the rule
	StartPort NullNumber
}

func NewCreateIpForwardingRuleParameter(ipaddressid string, protocol string, startport int) (p *CreateIpForwardingRuleParameter) {
	p = new(CreateIpForwardingRuleParameter)
	p.IpAddressId.Set(ipaddressid)
	p.Protocol.Set(protocol)
	p.StartPort.Set(startport)
	return p
}

// Creates an ip forwarding rule
func (c *Client) CreateIpForwardingRule(p *CreateIpForwardingRuleParameter) (*IpForwardingRule, error) {
	obj, err := c.Request("createIpForwardingRule", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*IpForwardingRule), err
}


// DisableStaticNat represents the paramter of DisableStaticNat
type DisableStaticNatParameter struct {
	// the public IP address id for which static nat feature is being disableed
	IpAddressId ID
}

func NewDisableStaticNatParameter(ipaddressid string) (p *DisableStaticNatParameter) {
	p = new(DisableStaticNatParameter)
	p.IpAddressId.Set(ipaddressid)
	return p
}

// Disables static rule for given ip address
func (c *Client) DisableStaticNat(p *DisableStaticNatParameter) (*Result, error) {
	obj, err := c.Request("disableStaticNat", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// ListIpForwardingRules represents the paramter of ListIpForwardingRules
type ListIpForwardingRulesParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// list only resources belonging to the domain specified
	DomainId ID
	// Lists rule with the specified ID.
	Id ID
	// list the rule belonging to this public ip address
	IpAddressId ID
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll  NullBool
	Page     NullNumber
	PageSize NullNumber
	// list objects by project
	ProjectId ID
	// Lists all rules applied to the specified Vm.
	VirtualMachineId ID
}

func NewListIpForwardingRulesParameter() (p *ListIpForwardingRulesParameter) {
	p = new(ListIpForwardingRulesParameter)
	return p
}

// List the ip forwarding rules
func (c *Client) ListIpForwardingRules(p *ListIpForwardingRulesParameter) ([]*IpForwardingRule, error) {
	obj, err := c.Request("listIpForwardingRules", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*IpForwardingRule), err
}

