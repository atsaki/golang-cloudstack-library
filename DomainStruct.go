package cloudstack

type Domain struct {
	ResourceBase
	// whether the domain has one or more sub-domains
	HasChild NullBool `json:"haschild"`
	// the ID of the domain
	Id ID `json:"id"`
	// the level of the domain
	Level NullNumber `json:"level"`
	// the name of the domain
	Name NullString `json:"name"`
	// the network domain
	NetworkDomain NullString `json:"networkdomain"`
	// the domain ID of the parent domain
	ParentDomainId ID `json:"parentdomainid"`
	// the domain name of the parent domain
	ParentDomainName NullString `json:"parentdomainname"`
	// the path of the domain
	Path NullString `json:"path"`
}
