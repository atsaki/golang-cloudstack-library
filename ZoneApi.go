package cloudstack


// DeleteZone represents the paramter of DeleteZone
type DeleteZoneParameter struct {
	// the ID of the Zone
	Id ID
}

func NewDeleteZoneParameter(id string) (p *DeleteZoneParameter) {
	p = new(DeleteZoneParameter)
	p.Id.Set(id)
	return p
}

// Deletes a Zone.
func (c *Client) DeleteZone(p *DeleteZoneParameter) (*Result, error) {
	obj, err := c.Request("deleteZone", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// UpdateZone represents the paramter of UpdateZone
type UpdateZoneParameter struct {
	// Allocation state of this cluster for allocation of new resources
	AllocationState NullString
	// the details for the Zone
	Details map[string]string
	// the dhcp Provider for the Zone
	DhcpProvider NullString
	// the first DNS for the Zone
	Dns1 NullString
	// the second DNS for the Zone
	Dns2 NullString
	// the dns search order list
	Dnssearchorder []string
	// Network domain name for the networks in the zone; empty string will update
	// domain with NULL value
	Domain NullString
	// the guest CIDR address for the Zone
	GuestCidrAddress NullString
	// the ID of the Zone
	Id ID
	// the first internal DNS for the Zone
	InternalDns1 NullString
	// the second internal DNS for the Zone
	InternalDns2 NullString
	// the first DNS for IPv6 network in the Zone
	Ip6Dns1 NullString
	// the second DNS for IPv6 network in the Zone
	Ip6Dns2 NullString
	// updates a private zone to public if set, but not vice-versa
	IsPublic NullBool
	// true if local storage offering enabled, false otherwise
	LocalStorageEnabled NullBool
	// the name of the Zone
	Name NullString
}

func NewUpdateZoneParameter(id string) (p *UpdateZoneParameter) {
	p = new(UpdateZoneParameter)
	p.Id.Set(id)
	return p
}

// Updates a Zone.
func (c *Client) UpdateZone(p *UpdateZoneParameter) (*Zone, error) {
	obj, err := c.Request("updateZone", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Zone), err
}


// CreateZone represents the paramter of CreateZone
type CreateZoneParameter struct {
	// Allocation state of this Zone for allocation of new resources
	AllocationState NullString
	// the first DNS for the Zone
	Dns1 NullString
	// the second DNS for the Zone
	Dns2 NullString
	// Network domain name for the networks in the zone
	Domain NullString
	// the ID of the containing domain, null for public zones
	DomainId ID
	// the guest CIDR address for the Zone
	GuestCidrAddress NullString
	// the first internal DNS for the Zone
	InternalDns1 NullString
	// the second internal DNS for the Zone
	InternalDns2 NullString
	// the first DNS for IPv6 network in the Zone
	Ip6Dns1 NullString
	// the second DNS for IPv6 network in the Zone
	Ip6Dns2 NullString
	// true if local storage offering enabled, false otherwise
	LocalStorageEnabled NullBool
	// the name of the Zone
	Name NullString
	// network type of the zone, can be Basic or Advanced
	NetworkType NullString
	// true if network is security group enabled, false otherwise
	SecurityGroupEnabled NullBool
}

func NewCreateZoneParameter(dns1 string, internaldns1 string, name string, networktype string) (p *CreateZoneParameter) {
	p = new(CreateZoneParameter)
	p.Dns1.Set(dns1)
	p.InternalDns1.Set(internaldns1)
	p.Name.Set(name)
	p.NetworkType.Set(networktype)
	return p
}

// Creates a Zone.
func (c *Client) CreateZone(p *CreateZoneParameter) (*Zone, error) {
	obj, err := c.Request("createZone", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Zone), err
}


// ListZones represents the paramter of ListZones
type ListZonesParameter struct {
	// true if you want to retrieve all available Zones. False if you only want to
	// return the Zones from which you have at least one VM. Default is false.
	Available NullBool
	// the ID of the domain associated with the zone
	DomainId ID
	// the ID of the zone
	Id ID
	// List by keyword
	Keyword NullString
	// the name of the zone
	Name NullString
	// the network type of the zone that the virtual machine belongs to
	NetworkType NullString
	Page        NullNumber
	PageSize    NullNumber
	// flag to display the capacity of the zones
	ShowCapacities NullBool
	// List zones by resource tags (key/value pairs)
	Tags map[string]string
}

func NewListZonesParameter() (p *ListZonesParameter) {
	p = new(ListZonesParameter)
	return p
}

// Lists zones
func (c *Client) ListZones(p *ListZonesParameter) ([]*Zone, error) {
	obj, err := c.Request("listZones", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Zone), err
}

