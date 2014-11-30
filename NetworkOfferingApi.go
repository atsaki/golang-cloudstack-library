package cloudstack


// ListNetworkOfferings represents the paramter of ListNetworkOfferings
type ListNetworkOfferingsParameter struct {
	// the availability of network offering. Default value is Required
	Availability NullString
	// list network offerings by display text
	DisplayText NullString
	// the network offering can be used only for network creation inside the VPC
	ForVpc NullBool
	// list network offerings by guest type: Shared or Isolated
	GuestIpType NullString
	// list network offerings by id
	Id ID
	// true if need to list only default network offerings. Default value is false
	IsDefault NullBool
	// true if offering has tags specified
	IsTagged NullBool
	// List by keyword
	Keyword NullString
	// list network offerings by name
	Name NullString
	// the ID of the network. Pass this in if you want to see the available network
	// offering that a network can be changed to.
	NetworkId ID
	Page      NullNumber
	PageSize  NullNumber
	// true if need to list only netwok offerings where source nat is supported,
	// false otherwise
	Sourcenatsupported NullBool
	// true if need to list only network offerings which support specifying ip
	// ranges
	SpecifyIpRanges NullBool
	// the tags for the network offering.
	SpecifyVlan NullBool
	// list network offerings by state
	State NullString
	// list network offerings supporting certain services
	SupportedServices []string
	// list network offerings by tags
	Tags NullString
	// list by traffic type
	TrafficType NullString
	// list netowrk offerings available for network creation in specific zone
	ZoneId ID
}

func NewListNetworkOfferingsParameter() (p *ListNetworkOfferingsParameter) {
	p = new(ListNetworkOfferingsParameter)
	return p
}

// Lists all available network offerings.
func (c *Client) ListNetworkOfferings(p *ListNetworkOfferingsParameter) ([]*NetworkOffering, error) {
	obj, err := c.Request("listNetworkOfferings", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*NetworkOffering), err
}


// DeleteNetworkOffering represents the paramter of DeleteNetworkOffering
type DeleteNetworkOfferingParameter struct {
	// the ID of the network offering
	Id ID
}

func NewDeleteNetworkOfferingParameter(id string) (p *DeleteNetworkOfferingParameter) {
	p = new(DeleteNetworkOfferingParameter)
	p.Id.Set(id)
	return p
}

// Deletes a network offering.
func (c *Client) DeleteNetworkOffering(p *DeleteNetworkOfferingParameter) (*Result, error) {
	obj, err := c.Request("deleteNetworkOffering", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// CreateNetworkOffering represents the paramter of CreateNetworkOffering
type CreateNetworkOfferingParameter struct {
	// the availability of network offering. Default value is Optional
	Availability NullString
	// true if the network offering is IP conserve mode enabled
	ConserveMode NullBool
	// Network offering details in key/value pairs. Supported keys are
	// internallbprovider/publiclbprovider with service provider as a value
	Details map[string]string
	// the display text of the network offering
	DisplayText NullString
	// true if guest network default egress policy is allow; false if default egress
	// policy is deny
	EgressDefaultPolicy NullBool
	// guest type of the network offering: Shared or Isolated
	GuestIpType NullString
	// true if network offering supports persistent networks; defaulted to false if
	// not specified
	IsPersistent NullBool
	// if true keepalive will be turned on in the loadbalancer. At the time of
	// writing this has only an effect on haproxy; the mode http and httpclose
	// options are unset in the haproxy conf file.
	KeepaliveEnabled NullBool
	// maximum number of concurrent connections supported by the network offering
	MaxConnections NullNumber
	// the name of the network offering
	Name NullString
	// data transfer rate in megabits per second allowed
	NetworkRate NullNumber
	// desired service capabilities as part of network offering
	ServiceCapabilityList map[string]string
	// the service offering ID used by virtual router provider
	ServiceOfferingId ID
	// provider to service mapping. If not specified, the provider for the service
	// will be mapped to the default provider on the physical network
	ServiceProviderList map[string]string
	// true if network offering supports specifying ip ranges; defaulted to false if
	// not specified
	SpecifyIpRanges NullBool
	// true if network offering supports vlans
	SpecifyVlan NullBool
	// services supported by the network offering
	SupportedServices []string
	// the tags for the network offering.
	Tags NullString
	// the traffic type for the network offering. Supported type in current release
	// is GUEST only
	TrafficType NullString
}

func NewCreateNetworkOfferingParameter(displaytext string, guestiptype string, name string, supportedservices []string, traffictype string) (p *CreateNetworkOfferingParameter) {
	p = new(CreateNetworkOfferingParameter)
	p.DisplayText.Set(displaytext)
	p.GuestIpType.Set(guestiptype)
	p.Name.Set(name)
	p.SupportedServices = supportedservices
	p.TrafficType.Set(traffictype)
	return p
}

// Creates a network offering.
func (c *Client) CreateNetworkOffering(p *CreateNetworkOfferingParameter) (*NetworkOffering, error) {
	obj, err := c.Request("createNetworkOffering", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*NetworkOffering), err
}


// UpdateNetworkOffering represents the paramter of UpdateNetworkOffering
type UpdateNetworkOfferingParameter struct {
	// the availability of network offering. Default value is Required for Guest
	// Virtual network offering; Optional for Guest Direct network offering
	Availability NullString
	// the display text of the network offering
	DisplayText NullString
	// the id of the network offering
	Id ID
	// if true keepalive will be turned on in the loadbalancer. At the time of
	// writing this has only an effect on haproxy; the mode http and httpclose
	// options are unset in the haproxy conf file.
	KeepaliveEnabled NullBool
	// maximum number of concurrent connections supported by the network offering
	MaxConnections NullNumber
	// the name of the network offering
	Name NullString
	// sort key of the network offering, integer
	SortKey NullNumber
	// update state for the network offering
	State NullString
}

func NewUpdateNetworkOfferingParameter() (p *UpdateNetworkOfferingParameter) {
	p = new(UpdateNetworkOfferingParameter)
	return p
}

// Updates a network offering.
func (c *Client) UpdateNetworkOffering(p *UpdateNetworkOfferingParameter) (*NetworkOffering, error) {
	obj, err := c.Request("updateNetworkOffering", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*NetworkOffering), err
}

