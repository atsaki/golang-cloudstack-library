package cloudstack

type NetworkServiceProvider struct {
	ResourceBase
	// true if individual services can be enabled/disabled
	CanEnableIndividualService NullBool `json:"canenableindividualservice"`
	// the destination physical network
	DestinationPhysicalNetworkId ID `json:"destinationphysicalnetworkid"`
	// uuid of the network provider
	Id ID `json:"id"`
	// the provider name
	Name NullString `json:"name"`
	// the physical network this belongs to
	PhysicalNetworkId ID `json:"physicalnetworkid"`
	// services for this provider
	ServiceList []NullString `json:"servicelist"`
	// state of the network provider
	State NullString `json:"state"`
}
