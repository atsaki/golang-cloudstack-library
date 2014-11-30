package cloudstack

type NetworkOffering struct {
	ResourceBase
	// availability of the network offering
	Availability NullString `json:"availability"`
	// true if network offering is ip conserve mode enabled
	ConserveMode NullBool `json:"conservemode"`
	// the date this network offering was created
	Created NullString `json:"created"`
	// additional key/value details tied with network offering
	Details NullString `json:"details"`
	// an alternate display text of the network offering.
	DisplayText NullString `json:"displaytext"`
	// true if guest network default egress policy is allow; false if default egress
	// policy is deny
	EgressDefaultPolicy NullBool `json:"egressdefaultpolicy"`
	// true if network offering can be used by VPC networks only
	ForVpc NullBool `json:"forvpc"`
	// guest type of the network offering, can be Shared or Isolated
	GuestIpType NullString `json:"guestiptype"`
	// the id of the network offering
	Id ID `json:"id"`
	// true if network offering is default, false otherwise
	IsDefault NullBool `json:"isdefault"`
	// true if network offering supports persistent networks, false otherwise
	IsPersistent NullBool `json:"ispersistent"`
	// maximum number of concurrents connections to be handled by lb
	MaxConnections NullNumber `json:"maxconnections"`
	// the name of the network offering
	Name NullString `json:"name"`
	// data transfer rate in megabits per second allowed.
	NetworkRate NullNumber `json:"networkrate"`
	// the list of supported services
	Service []NetworkService `json:"service"`
	// the ID of the service offering used by virtual router provider
	ServiceOfferingId ID `json:"serviceofferingid"`
	// true if network offering supports specifying ip ranges, false otherwise
	SpecifyIpRanges NullBool `json:"specifyipranges"`
	// true if network offering supports vlans, false otherwise
	SpecifyVlan NullBool `json:"specifyvlan"`
	// state of the network offering. Can be Disabled/Enabled/Inactive
	State NullString `json:"state"`
	// true if network offering supports network that span multiple zones
	SupportsStrechedL2Subnet NullBool `json:"supportsstrechedl2subnet"`
	// the tags for the network offering
	Tags NullString `json:"tags"`
	// the traffic type for the network offering, supported types are Public,
	// Management, Control, Guest, Vlan or Storage.
	TrafficType NullString `json:"traffictype"`
}
