package cloudstack

type NetworkServiceCapability struct {
	// the capability name
	Name NullString `json:"name"`
	// the capability value
	Value NullString `json:"value"`
	// can this service capability value can be choosable while creatine network offerings
	CanChooseServiceCapability NullBool `json:"canchooseservicecapability"`
}
