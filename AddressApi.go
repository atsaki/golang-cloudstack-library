package cloudstack


// ListPublicIpAddresses represents the paramter of ListPublicIpAddresses
type ListPublicIpAddressesParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// limits search results to allocated public IP addresses
	AllocatedOnly NullBool
	// lists all public IP addresses associated to the network specified
	AssociatedNetworkId ID
	// list only resources belonging to the domain specified
	DomainId ID
	// list resources by display flag; only ROOT admin is eligible to pass this
	// parameter
	ForDisplay NullBool
	// list only ips used for load balancing
	ForLoadBalancing NullBool
	// the virtual network for the IP address
	ForVirtualNetwork NullBool
	// lists ip address by id
	Id ID
	// lists the specified IP address
	IpAddress NullString
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// list only source nat ip addresses
	IsSourceNat NullBool
	// list only static nat ip addresses
	IsStaticNat NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll  NullBool
	Page     NullNumber
	PageSize NullNumber
	// lists all public IP addresses by physical network id
	PhysicalNetworkId ID
	// list objects by project
	ProjectId ID
	// List resources by tags (key/value pairs)
	Tags map[string]string
	// lists all public IP addresses by VLAN ID
	VlanId ID
	// List ips belonging to the VPC
	VpcId ID
	// lists all public IP addresses by Zone ID
	ZoneId ID
}

func NewListPublicIpAddressesParameter() (p *ListPublicIpAddressesParameter) {
	p = new(ListPublicIpAddressesParameter)
	return p
}

// Lists all public ip addresses
func (c *Client) ListPublicIpAddresses(p *ListPublicIpAddressesParameter) ([]*PublicIpAddress, error) {
	obj, err := c.Request("listPublicIpAddresses", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*PublicIpAddress), err
}


// UpdateIpAddress represents the paramter of UpdateIpAddress
type UpdateIpAddressParameter struct {
	// an optional field, in case you want to set a custom id to the resource.
	// Allowed to Root Admins only
	CustomId ID
	// an optional field, whether to the display the ip to the end user or not
	ForDisplay NullBool
	// the id of the public ip address to update
	Id ID
}

func NewUpdateIpAddressParameter(id string) (p *UpdateIpAddressParameter) {
	p = new(UpdateIpAddressParameter)
	p.Id.Set(id)
	return p
}

// Updates an ip address
func (c *Client) UpdateIpAddress(p *UpdateIpAddressParameter) (*PublicIpAddress, error) {
	obj, err := c.Request("updateIpAddress", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*PublicIpAddress), err
}


// DisassociateIpAddress represents the paramter of DisassociateIpAddress
type DisassociateIpAddressParameter struct {
	// the id of the public ip address to disassociate
	Id ID
}

func NewDisassociateIpAddressParameter(id string) (p *DisassociateIpAddressParameter) {
	p = new(DisassociateIpAddressParameter)
	p.Id.Set(id)
	return p
}

// Disassociates an ip address from the account.
func (c *Client) DisassociateIpAddress(p *DisassociateIpAddressParameter) (*Result, error) {
	obj, err := c.Request("disassociateIpAddress", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// AssociateIpAddress represents the paramter of AssociateIpAddress
type AssociateIpAddressParameter struct {
	// the account to associate with this IP address
	Account NullString
	// the ID of the domain to associate with this IP address
	DomainId ID
	// an optional field, whether to the display the ip to the end user or not
	ForDisplay NullBool
	// should be set to true if public IP is required to be transferable across
	// zones, if not specified defaults to false
	IsPortable NullBool
	// The network this ip address should be associated to.
	NetworkId ID
	// Deploy vm for the project
	ProjectId ID
	// region ID from where portable ip is to be associated.
	RegionId ID
	// the VPC you want the ip address to be associated with
	VpcId ID
	// the ID of the availability zone you want to acquire an public IP address from
	ZoneId ID
}

func NewAssociateIpAddressParameter() (p *AssociateIpAddressParameter) {
	p = new(AssociateIpAddressParameter)
	return p
}

// Acquires and associates a public IP to an account.
func (c *Client) AssociateIpAddress(p *AssociateIpAddressParameter) (*PublicIpAddress, error) {
	obj, err := c.Request("associateIpAddress", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*PublicIpAddress), err
}

