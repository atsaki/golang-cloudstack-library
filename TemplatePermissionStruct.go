package cloudstack

type TemplatePermission struct {
	ResourceBase
	// the list of accounts the template is available for
	Account []NullString `json:"account"`
	// the ID of the domain to which the template belongs
	DomainId ID `json:"domainid"`
	// the template ID
	Id ID `json:"id"`
	// true if this template is a public template, false otherwise
	IsPublic NullBool `json:"ispublic"`
	// the list of projects the template is available for
	ProjectIds []ID `json:"projectids"`
}
