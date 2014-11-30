package cloudstack


// CopyTemplate represents the paramter of CopyTemplate
type CopyTemplateParameter struct {
	// ID of the zone the template is being copied to.
	DestZoneId ID
	// Template ID.
	Id ID
	// ID of the zone the template is currently hosted on. If not specified and
	// template is cross-zone, then we will sync this template to region wide image
	// store.
	SourceZoneId ID
}

func NewCopyTemplateParameter(destzoneid string, id string) (p *CopyTemplateParameter) {
	p = new(CopyTemplateParameter)
	p.DestZoneId.Set(destzoneid)
	p.Id.Set(id)
	return p
}

// Copies a template from one zone to another.
func (c *Client) CopyTemplate(p *CopyTemplateParameter) (*Template, error) {
	obj, err := c.Request("copyTemplate", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Template), err
}


// CreateTemplate represents the paramter of CreateTemplate
type CreateTemplateParameter struct {
	// 32 or 64 bit
	Bits NullNumber
	// Template details in key/value pairs.
	Details map[string]string
	// the display text of the template. This is usually used for display purposes.
	DisplayText NullString
	// true if template contains XS/VMWare tools inorder to support dynamic scaling
	// of VM cpu/memory
	IsDynamicallyScalable NullBool
	// true if this template is a featured template, false otherwise
	IsFeatured NullBool
	// true if this template is a public template, false otherwise
	IsPublic NullBool
	// the name of the template
	Name NullString
	// the ID of the OS Type that best represents the OS of this template.
	OsTypeId ID
	// true if the template supports the password reset feature; default is false
	PasswordEnabled NullBool
	// true if the template requres HVM, false otherwise
	RequiresHvm NullBool
	// the ID of the snapshot the template is being created from. Either this
	// parameter, or volumeId has to be passed in
	SnapshotId ID
	// the tag for this template.
	Templatetag NullString
	// Optional, only for baremetal hypervisor. The directory name where template
	// stored on CIFS server
	Url NullString
	// Optional, VM ID. If this presents, it is going to create a baremetal template
	// for VM this ID refers to. This is only for VM whose hypervisor type is
	// BareMetal
	VirtualMachineId ID
	// the ID of the disk volume the template is being created from. Either this
	// parameter, or snapshotId has to be passed in
	VolumeId ID
}

func NewCreateTemplateParameter(displaytext string, name string, ostypeid string) (p *CreateTemplateParameter) {
	p = new(CreateTemplateParameter)
	p.DisplayText.Set(displaytext)
	p.Name.Set(name)
	p.OsTypeId.Set(ostypeid)
	return p
}

// Creates a template of a virtual machine. The virtual machine must be in a
// STOPPED state. A template created from this command is automatically
// designated as a private template visible to the account that created it.
func (c *Client) CreateTemplate(p *CreateTemplateParameter) (*Template, error) {
	obj, err := c.Request("createTemplate", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Template), err
}


// ListTemplatePermissions represents the paramter of ListTemplatePermissions
type ListTemplatePermissionsParameter struct {
	// the template ID
	Id ID
}

func NewListTemplatePermissionsParameter(id string) (p *ListTemplatePermissionsParameter) {
	p = new(ListTemplatePermissionsParameter)
	p.Id.Set(id)
	return p
}

// List template visibility and all accounts that have permissions to view this
// template.
func (c *Client) ListTemplatePermissions(p *ListTemplatePermissionsParameter) ([]*TemplatePermission, error) {
	obj, err := c.Request("listTemplatePermissions", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*TemplatePermission), err
}


// DeleteTemplate represents the paramter of DeleteTemplate
type DeleteTemplateParameter struct {
	// the ID of the template
	Id ID
	// the ID of zone of the template
	ZoneId ID
}

func NewDeleteTemplateParameter(id string) (p *DeleteTemplateParameter) {
	p = new(DeleteTemplateParameter)
	p.Id.Set(id)
	return p
}

// Deletes a template from the system. All virtual machines using the deleted
// template will not be affected.
func (c *Client) DeleteTemplate(p *DeleteTemplateParameter) (*Result, error) {
	obj, err := c.Request("deleteTemplate", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// RegisterTemplate represents the paramter of RegisterTemplate
type RegisterTemplateParameter struct {
	// an optional accountName. Must be used with domainId.
	Account NullString
	// 32 or 64 bits support. 64 by default
	Bits NullNumber
	// the MD5 checksum value of this template
	Checksum NullString
	// Template details in key/value pairs.
	Details map[string]string
	// the display text of the template. This is usually used for display purposes.
	DisplayText NullString
	// an optional domainId. If the account parameter is used, domainId must also be
	// used.
	DomainId ID
	// the format for the template. Possible values include QCOW2, RAW, and VHD.
	Format NullString
	// the target hypervisor for the template
	Hypervisor NullString
	// true if template contains XS/VMWare tools inorder to support dynamic scaling
	// of VM cpu/memory
	IsDynamicallyScalable NullBool
	// true if the template or its derivatives are extractable; default is false
	IsExtractable NullBool
	// true if this template is a featured template, false otherwise
	IsFeatured NullBool
	// true if the template is available to all accounts; default is true
	IsPublic NullBool
	// true if the template type is routing i.e., if template is used to deploy
	// router
	IsRouting NullBool
	// the name of the template
	Name NullString
	// the ID of the OS Type that best represents the OS of this template.
	OsTypeId ID
	// true if the template supports the password reset feature; default is false
	PasswordEnabled NullBool
	// Register template for the project
	ProjectId ID
	// true if this template requires HVM
	RequiresHvm NullBool
	// true if the template supports the sshkey upload feature; default is false
	SshKeyEnabled NullBool
	// the tag for this template.
	Templatetag NullString
	// the URL of where the template is hosted. Possible URL include http:// and
	// https://
	Url NullString
	// the ID of the zone the template is to be hosted on
	ZoneId ID
}

func NewRegisterTemplateParameter(displaytext string, format string, hypervisor string, name string, ostypeid string, url string, zoneid string) (p *RegisterTemplateParameter) {
	p = new(RegisterTemplateParameter)
	p.DisplayText.Set(displaytext)
	p.Format.Set(format)
	p.Hypervisor.Set(hypervisor)
	p.Name.Set(name)
	p.OsTypeId.Set(ostypeid)
	p.Url.Set(url)
	p.ZoneId.Set(zoneid)
	return p
}

// Registers an existing template into the CloudStack cloud.
func (c *Client) RegisterTemplate(p *RegisterTemplateParameter) (*Template, error) {
	obj, err := c.Request("registerTemplate", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Template), err
}


// ListTemplates represents the paramter of ListTemplates
type ListTemplatesParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// list only resources belonging to the domain specified
	DomainId ID
	// the hypervisor for which to restrict the search
	Hypervisor NullString
	// the template ID
	Id ID
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll NullBool
	// the template name
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
	// list objects by project
	ProjectId ID
	// show removed templates as well
	ShowRemoved NullBool
	// List resources by tags (key/value pairs)
	Tags map[string]string
	// possible values are "featured", "self",
	// "selfexecutable","sharedexecutable","executable", and "community". * featured
	// : templates that have been marked as featured and public. * self : templates
	// that have been registered or created by the calling user. * selfexecutable :
	// same as self, but only returns templates that can be used to deploy a new VM.
	// * sharedexecutable : templates ready to be deployed that have been granted to
	// the calling user by another user. * executable : templates that are owned by
	// the calling user, or public templates, that can be used to deploy a VM. *
	// community : templates that have been marked as public but not featured. * all
	// : all templates (only usable by admins).
	Templatefilter NullString
	// list templates by zoneId
	ZoneId ID
}

func NewListTemplatesParameter(templatefilter string) (p *ListTemplatesParameter) {
	p = new(ListTemplatesParameter)
	p.Templatefilter.Set(templatefilter)
	return p
}

// List all public, private, and privileged templates.
func (c *Client) ListTemplates(p *ListTemplatesParameter) ([]*Template, error) {
	obj, err := c.Request("listTemplates", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Template), err
}


// PrepareTemplate represents the paramter of PrepareTemplate
type PrepareTemplateParameter struct {
	// template ID of the template to be prepared in primary storage(s).
	TemplateId ID
	// zone ID of the template to be prepared in primary storage(s).
	ZoneId ID
}

func NewPrepareTemplateParameter(templateid string, zoneid string) (p *PrepareTemplateParameter) {
	p = new(PrepareTemplateParameter)
	p.TemplateId.Set(templateid)
	p.ZoneId.Set(zoneid)
	return p
}

// load template into primary storage
func (c *Client) PrepareTemplate(p *PrepareTemplateParameter) (*Template, error) {
	obj, err := c.Request("prepareTemplate", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Template), err
}


// UpdateTemplate represents the paramter of UpdateTemplate
type UpdateTemplateParameter struct {
	// true if image is bootable, false otherwise
	Bootable NullBool
	// the display text of the image
	DisplayText NullString
	// the format for the image
	Format NullString
	// the ID of the image file
	Id ID
	// true if template/ISO contains XS/VMWare tools inorder to support dynamic
	// scaling of VM cpu/memory
	IsDynamicallyScalable NullBool
	// true if the template type is routing i.e., if template is used to deploy
	// router
	IsRouting NullBool
	// the name of the image file
	Name NullString
	// the ID of the OS type that best represents the OS of this image.
	OsTypeId ID
	// true if the image supports the password reset feature; default is false
	PasswordEnabled NullBool
	// sort key of the template, integer
	SortKey NullNumber
}

func NewUpdateTemplateParameter(id string) (p *UpdateTemplateParameter) {
	p = new(UpdateTemplateParameter)
	p.Id.Set(id)
	return p
}

// Updates attributes of a template.
func (c *Client) UpdateTemplate(p *UpdateTemplateParameter) (*Template, error) {
	obj, err := c.Request("updateTemplate", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Template), err
}


// ExtractTemplate represents the paramter of ExtractTemplate
type ExtractTemplateParameter struct {
	// the ID of the template
	Id ID
	// the mode of extraction - HTTP_DOWNLOAD or FTP_UPLOAD
	Mode NullString
	// the url to which the ISO would be extracted
	Url NullString
	// the ID of the zone where the ISO is originally located
	ZoneId ID
}

func NewExtractTemplateParameter(id string, mode string) (p *ExtractTemplateParameter) {
	p = new(ExtractTemplateParameter)
	p.Id.Set(id)
	p.Mode.Set(mode)
	return p
}

// Extracts a template
func (c *Client) ExtractTemplate(p *ExtractTemplateParameter) (*Template, error) {
	obj, err := c.Request("extractTemplate", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Template), err
}


// UpdateTemplatePermissions represents the paramter of UpdateTemplatePermissions
type UpdateTemplatePermissionsParameter struct {
	// a comma delimited list of accounts. If specified, "op" parameter has to be
	// passed in.
	Accounts []string
	// the template ID
	Id ID
	// true if the template/iso is extractable, false other wise. Can be set only by
	// root admin
	IsExtractable NullBool
	// true for featured template/iso, false otherwise
	IsFeatured NullBool
	// true for public template/iso, false for private templates/isos
	IsPublic NullBool
	// permission operator (add, remove, reset)
	Op NullString
	// a comma delimited list of projects. If specified, "op" parameter has to be
	// passed in.
	ProjectIds []string
}

func NewUpdateTemplatePermissionsParameter(id string) (p *UpdateTemplatePermissionsParameter) {
	p = new(UpdateTemplatePermissionsParameter)
	p.Id.Set(id)
	return p
}

// Updates a template visibility permissions. A public template is visible to
// all accounts within the same domain. A private template is visible only to
// the owner of the template. A priviledged template is a private template with
// account permissions added. Only accounts specified under the template
// permissions are visible to them.
func (c *Client) UpdateTemplatePermissions(p *UpdateTemplatePermissionsParameter) (*Result, error) {
	obj, err := c.Request("updateTemplatePermissions", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}

