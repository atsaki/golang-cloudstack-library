package cloudstack

type NetworkService struct {
	// the list of capabilities
	Capability []NetworkServiceCapability `json:"capability"`
	// the service name
	Name NullString `json:"name"`
	// the service provider name
	Provider []NetworkServiceProvider `json:"provider"`
}
