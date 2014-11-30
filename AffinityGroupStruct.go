package cloudstack

type AffinityGroup struct {
	ResourceBase
	// the account owning the affinity group
	Account NullString `json:"account"`
	// the description of the affinity group
	Description NullString `json:"description"`
	// the domain name of the affinity group
	Domain NullString `json:"domain"`
	// the domain ID of the affinity group
	DomainId ID `json:"domainid"`
	// the ID of the affinity group
	Id ID `json:"id"`
	// the name of the affinity group
	Name NullString `json:"name"`
	// the type of the affinity group
	Type NullString `json:"type"`
	// virtual machine Ids associated with this affinity group
	VirtualMachineIds []ID `json:"virtualmachineids"`
}
