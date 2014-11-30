package cloudstack

type Template struct {
	ResourceBase
	// the account name to which the template belongs
	Account NullString `json:"account"`
	// the account id to which the template belongs
	AccountId ID `json:"accountid"`
	// true if the ISO is bootable, false otherwise
	Bootable NullBool `json:"bootable"`
	// checksum of the template
	Checksum NullString `json:"checksum"`
	// the date this template was created
	Created NullString `json:"created"`
	// true if the template is managed across all Zones, false otherwise
	CrossZones NullBool `json:"crosszones"`
	// additional key/value details tied with template
	Details NullString `json:"details"`
	// the template display text
	DisplayText NullString `json:"displaytext"`
	// the name of the domain to which the template belongs
	Domain NullString `json:"domain"`
	// the ID of the domain to which the template belongs
	DomainId ID `json:"domainid"`
	// the format of the template.
	Format NullString `json:"format"`
	// the ID of the secondary storage host for the template
	HostId ID `json:"hostid"`
	// the name of the secondary storage host for the template
	HostName NullString `json:"hostname"`
	// the hypervisor on which the template runs
	Hypervisor NullString `json:"hypervisor"`
	// the template ID
	Id ID `json:"id"`
	// true if template contains XS/VMWare tools inorder to support dynamic scaling
	// of VM cpu/memory
	IsDynamicallyScalable NullBool `json:"isdynamicallyscalable"`
	// true if the template is extractable, false otherwise
	IsExtractable NullBool `json:"isextractable"`
	// true if this template is a featured template, false otherwise
	IsFeatured NullBool `json:"isfeatured"`
	// true if this template is a public template, false otherwise
	IsPublic NullBool `json:"ispublic"`
	// true if the template is ready to be deployed from, false otherwise.
	IsReady NullBool `json:"isready"`
	// the template name
	Name NullString `json:"name"`
	// the ID of the OS type for this template.
	OsTypeId ID `json:"ostypeid"`
	// the name of the OS type for this template.
	OsTypeName NullString `json:"ostypename"`
	// true if the reset password feature is enabled, false otherwise
	PasswordEnabled NullBool `json:"passwordenabled"`
	// the project name of the template
	Project NullString `json:"project"`
	// the project id of the template
	ProjectId ID `json:"projectid"`
	// the date this template was removed
	Removed NullString `json:"removed"`
	// the size of the template
	Size NullNumber `json:"size"`
	// the template ID of the parent template if present
	SourceTemplateId ID `json:"sourcetemplateid"`
	// true if template is sshkey enabled, false otherwise
	SshKeyEnabled NullBool `json:"sshkeyenabled"`
	// the status of the template
	Status NullString `json:"status"`
	// the list of resource tags associated with tempate
	Tags []Tag `json:"tags"`
	// the tag of this template
	Templatetag NullString `json:"templatetag"`
	// the type of the template
	TemplateType NullString `json:"templatetype"`
	// the ID of the zone for this template
	ZoneId ID `json:"zoneid"`
	// the name of the zone for this template
	ZoneName NullString `json:"zonename"`
}
