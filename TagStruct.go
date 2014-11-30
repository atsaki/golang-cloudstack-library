package cloudstack

type Tag struct {
	ResourceBase
	// the account associated with the tag
	Account NullString `json:"account"`
	// customer associated with the tag
	Customer NullString `json:"customer"`
	// the domain associated with the tag
	Domain NullString `json:"domain"`
	// the ID of the domain associated with the tag
	DomainId ID `json:"domainid"`
	// tag key name
	Key NullString `json:"key"`
	// the project name where tag belongs to
	Project NullString `json:"project"`
	// the project id the tag belongs to
	ProjectId ID `json:"projectid"`
	// id of the resource
	ResourceId ID `json:"resourceid"`
	// resource type
	ResourceType NullString `json:"resourcetype"`
	// tag value
	Value NullString `json:"value"`
}
