package cloudstack

type Affinitygroup struct {
	Account           NullString `json:"account"`
	Description       NullString `json:"description"`
	Domain            NullString `json:"domain"`
	Domainid          ID         `json:"domainid"`
	Id                ID         `json:"id"`
	Name              NullString `json:"name"`
	Type              NullString `json:"type"`
	VirtualmachineIds []ID       `json:"virtualmachineids"`
}
