package cloudstack

type DedicatedHost struct {
	ResourceBase
	// the Account ID of the host
	AccountId ID `json:"accountid"`
	// the Dedication Affinity Group ID of the host
	AffinityGroupId ID `json:"affinitygroupid"`
	// the domain ID of the host
	DomainId ID `json:"domainid"`
	// the ID of the host
	HostId ID `json:"hostid"`
	// the name of the host
	HostName NullString `json:"hostname"`
	// the ID of the dedicated resource
	Id ID `json:"id"`
}
