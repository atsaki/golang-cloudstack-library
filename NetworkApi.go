package cloudstack


// CreateNetwork represents the paramter of CreateNetwork
type CreateNetworkParameter struct {
	// account who will own the network
	Account NullString
	// Network ACL Id associated for the network
	AclId ID
	// Access control type; supported values are account and domain. In 3.0 all
	// shared networks should have aclType=Domain, and all Isolated networks -
	// Account. Account means that only the account owner can use the network,
	// domain - all accouns in the domain can use the network
	AclType NullString
	// an optional field, whether to the display the network to the end user or not.
	DisplayNetwork NullBool
	// the display text of the network
	DisplayText NullString
	// domain ID of the account owning a network
	DomainId ID
	// the ending IP address in the network IP range. If not specified, will be
	// defaulted to startIP
	EndIp NullString
	// the ending IPv6 address in the IPv6 network range
	EndIpv6 NullString
	// the gateway of the network. Required for Shared networks and Isolated
	// networks when it belongs to VPC
	Gateway NullString
	// the CIDR of IPv6 network, must be at least /64
	Ip6Cidr NullString
	// the gateway of the IPv6 network. Required for Shared networks and Isolated
	// networks when it belongs to VPC
	Ip6Gateway NullString
	// the isolated private vlan for this network
	IsolatedPvlan NullString
	// the name of the network
	Name NullString
	// the netmask of the network. Required for Shared networks and Isolated
	// networks when it belongs to VPC
	Netmask NullString
	// network domain
	NetworkDomain NullString
	// the network offering id
	NetworkOfferingId ID
	// the Physical Network ID the network belongs to
	PhysicalNetworkId ID
	// an optional project for the ssh key
	ProjectId ID
	// the beginning IP address in the network IP range
	StartIp NullString
	// the beginning IPv6 address in the IPv6 network range
	StartIpv6 NullString
	// Defines whether to allow subdomains to use networks dedicated to their parent
	// domain(s). Should be used with aclType=Domain, defaulted to
	// allow.subdomain.network.access global config if not specified
	SubDomainAccess NullBool
	// the ID or VID of the network
	Vlan NullString
	// the VPC network belongs to
	VpcId ID
	// the Zone ID for the network
	ZoneId ID
}

func NewCreateNetworkParameter(displaytext string, name string, networkofferingid string, zoneid string) (p *CreateNetworkParameter) {
	p = new(CreateNetworkParameter)
	p.DisplayText.Set(displaytext)
	p.Name.Set(name)
	p.NetworkOfferingId.Set(networkofferingid)
	p.ZoneId.Set(zoneid)
	return p
}

// Creates a network
func (c *Client) CreateNetwork(p *CreateNetworkParameter) (*Network, error) {
	obj, err := c.Request("createNetwork", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Network), err
}


// ListNetworks represents the paramter of ListNetworks
type ListNetworksParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// list networks by ACL (access control list) type. Supported values are Account
	// and Domain
	AclType NullString
	// list networks available for vm deployment
	CanUseForDeploy NullBool
	// list resources by display flag; only ROOT admin is eligible to pass this
	// parameter
	DisplayNetwork NullBool
	// list only resources belonging to the domain specified
	DomainId ID
	// the network belongs to vpc
	ForVpc NullBool
	// list networks by id
	Id ID
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// true if network is system, false otherwise
	IsSystem NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll  NullBool
	Page     NullNumber
	PageSize NullNumber
	// list networks by physical network id
	PhysicalNetworkId ID
	// list objects by project
	ProjectId ID
	// list networks by restartRequired
	ReStartRequired NullBool
	// true if need to list only networks which support specifying ip ranges
	SpecifyIpRanges NullBool
	// list networks supporting certain services
	SupportedServices []string
	// List resources by tags (key/value pairs)
	Tags map[string]string
	// type of the traffic
	TrafficType NullString
	// the type of the network. Supported values are: Isolated and Shared
	Type NullString
	// List networks by VPC
	VpcId ID
	// the Zone ID of the network
	ZoneId ID
}

func NewListNetworksParameter() (p *ListNetworksParameter) {
	p = new(ListNetworksParameter)
	return p
}

// Lists all available networks.
func (c *Client) ListNetworks(p *ListNetworksParameter) ([]*Network, error) {
	obj, err := c.Request("listNetworks", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Network), err
}


// RestartNetwork represents the paramter of RestartNetwork
type RestartNetworkParameter struct {
	// If cleanup old network elements
	Cleanup NullBool
	// The id of the network to restart.
	Id ID
}

func NewRestartNetworkParameter(id string) (p *RestartNetworkParameter) {
	p = new(RestartNetworkParameter)
	p.Id.Set(id)
	return p
}

// Restarts the network; includes 1) restarting network elements - virtual
// routers, dhcp servers 2) reapplying all public ips 3) reapplying
// loadBalancing/portForwarding rules
func (c *Client) RestartNetwork(p *RestartNetworkParameter) (*Network, error) {
	obj, err := c.Request("restartNetwork", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Network), err
}


// DeleteNetwork represents the paramter of DeleteNetwork
type DeleteNetworkParameter struct {
	// Force delete a network. Network will be marked as 'Destroy' even when
	// commands to shutdown and cleanup to the backend fails.
	Forced NullBool
	// the ID of the network
	Id ID
}

func NewDeleteNetworkParameter(id string) (p *DeleteNetworkParameter) {
	p = new(DeleteNetworkParameter)
	p.Id.Set(id)
	return p
}

// Deletes a network
func (c *Client) DeleteNetwork(p *DeleteNetworkParameter) (*Result, error) {
	obj, err := c.Request("deleteNetwork", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// UpdateNetwork represents the paramter of UpdateNetwork
type UpdateNetworkParameter struct {
	// Force update even if cidr type is different
	changeCidr NullBool
	// an optional field, in case you want to set a custom id to the resource.
	// Allowed to Root Admins only
	CustomId ID
	// an optional field, whether to the display the network to the end user or not.
	DisplayNetwork NullBool
	// the new display text for the network
	DisplayText NullString
	// CIDR for Guest VMs,Cloudstack allocates IPs to Guest VMs only from this CIDR
	GuestVmCidr NullString
	// the ID of the network
	Id ID
	// the new name for the network
	Name NullString
	// network domain
	NetworkDomain NullString
	// network offering ID
	NetworkOfferingId ID
}

func NewUpdateNetworkParameter(id string) (p *UpdateNetworkParameter) {
	p = new(UpdateNetworkParameter)
	p.Id.Set(id)
	return p
}

// Updates a network
func (c *Client) UpdateNetwork(p *UpdateNetworkParameter) (*Network, error) {
	obj, err := c.Request("updateNetwork", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Network), err
}

