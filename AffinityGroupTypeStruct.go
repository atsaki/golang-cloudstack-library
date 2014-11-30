package cloudstack

type AffinityGroupType struct {
	ResourceBase
	// the type of the affinity group
	Type NullString `json:"type"`
}
