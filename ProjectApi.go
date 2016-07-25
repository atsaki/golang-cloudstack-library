package cloudstack

// SuspendProject represents the paramter of SuspendProject
type SuspendProjectParameter struct {
	// id of the project to be suspended
	Id ID
}

func NewSuspendProjectParameter(id string) (p *SuspendProjectParameter) {
	p = new(SuspendProjectParameter)
	p.Id.Set(id)
	return p
}

// Suspends a project
func (c *Client) SuspendProject(p *SuspendProjectParameter) (*Project, error) {
	obj, err := c.Request("suspendProject", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Project), err
}

// ListProjects represents the paramter of ListProjects
type ListProjectsParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// list projects by display text
	DisplayText NullString
	// list only resources belonging to the domain specified
	DomainId ID
	// list projects by project ID
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
	// list projects by name
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
	// list projects by state
	State NullString
	// List projects by tags (key/value pairs)
	Tags map[string]string
}

func NewListProjectsParameter() (p *ListProjectsParameter) {
	p = new(ListProjectsParameter)
	return p
}

// Lists projects and provides detailed information for listed projects
func (c *Client) ListProjects(p *ListProjectsParameter) ([]*Project, error) {
	obj, err := c.Request("listProjects", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Project), err
}

// DeleteProject represents the paramter of DeleteProject
type DeleteProjectParameter struct {
	// id of the project to be deleted
	Id ID
}

func NewDeleteProjectParameter(id string) (p *DeleteProjectParameter) {
	p = new(DeleteProjectParameter)
	p.Id.Set(id)
	return p
}

// Deletes a project
func (c *Client) DeleteProject(p *DeleteProjectParameter) (*Result, error) {
	obj, err := c.Request("deleteProject", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}

// ActivateProject represents the paramter of ActivateProject
type ActivateProjectParameter struct {
	// id of the project to be modified
	Id ID
}

func NewActivateProjectParameter(id string) (p *ActivateProjectParameter) {
	p = new(ActivateProjectParameter)
	p.Id.Set(id)
	return p
}

// Activates a project
func (c *Client) ActivateProject(p *ActivateProjectParameter) (*Project, error) {
	obj, err := c.Request("activateProject", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Project), err
}

// ListProjectAccounts represents the paramter of ListProjectAccounts
type ListProjectAccountsParameter struct {
	// list accounts of the project by account name
	Account NullString
	// List by keyword
	Keyword  NullString
	Page     NullNumber
	PageSize NullNumber
	// id of the project
	ProjectId ID
	// list accounts of the project by role
	Role NullString
}

func NewListProjectAccountsParameter(projectid string) (p *ListProjectAccountsParameter) {
	p = new(ListProjectAccountsParameter)
	p.ProjectId.Set(projectid)
	return p
}

// Lists project's accounts
func (c *Client) ListProjectAccounts(p *ListProjectAccountsParameter) ([]*ProjectAccount, error) {
	obj, err := c.Request("listProjectAccounts", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*ProjectAccount), err
}

// ListProjectInvitations represents the paramter of ListProjectInvitations
type ListProjectInvitationsParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// if true, list only active invitations - having Pending state and ones that
	// are not timed out yet
	ActiveOnly NullBool
	// list only resources belonging to the domain specified
	DomainId ID
	// list invitations by id
	Id ID
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll  NullBool
	Page     NullNumber
	PageSize NullNumber
	// list by project id
	ProjectId ID
	// list invitations by state
	State NullString
}

func NewListProjectInvitationsParameter() (p *ListProjectInvitationsParameter) {
	p = new(ListProjectInvitationsParameter)
	return p
}

// Lists project invitations and provides detailed information for listed
// invitations
func (c *Client) ListProjectInvitations(p *ListProjectInvitationsParameter) ([]*ProjectInvitation, error) {
	obj, err := c.Request("listProjectInvitations", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*ProjectInvitation), err
}

// UpdateProjectInvitation represents the paramter of UpdateProjectInvitation
type UpdateProjectInvitationParameter struct {
	// if true, accept the invitation, decline if false. True by default
	Accept NullBool
	// account that is joining the project
	Account NullString
	// id of the project to join
	ProjectId ID
	// list invitations for specified account; this parameter has to be specified
	// with domainId
	Token NullString
}

func NewUpdateProjectInvitationParameter(projectid string) (p *UpdateProjectInvitationParameter) {
	p = new(UpdateProjectInvitationParameter)
	p.ProjectId.Set(projectid)
	return p
}

// Accepts or declines project invitation
func (c *Client) UpdateProjectInvitation(p *UpdateProjectInvitationParameter) (*Result, error) {
	obj, err := c.Request("updateProjectInvitation", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}

// UpdateProject represents the paramter of UpdateProject
type UpdateProjectParameter struct {
	// new Admin account for the project
	Account NullString
	// display text of the project
	DisplayText NullString
	// id of the project to be modified
	Id ID
}

func NewUpdateProjectParameter(id string) (p *UpdateProjectParameter) {
	p = new(UpdateProjectParameter)
	p.Id.Set(id)
	return p
}

// Updates a project
func (c *Client) UpdateProject(p *UpdateProjectParameter) (*Project, error) {
	obj, err := c.Request("updateProject", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Project), err
}

// DeleteProjectInvitation represents the paramter of DeleteProjectInvitation
type DeleteProjectInvitationParameter struct {
	// id of the invitation
	Id ID
}

func NewDeleteProjectInvitationParameter(id string) (p *DeleteProjectInvitationParameter) {
	p = new(DeleteProjectInvitationParameter)
	p.Id.Set(id)
	return p
}

// Deletes project invitation
func (c *Client) DeleteProjectInvitation(p *DeleteProjectInvitationParameter) (*Result, error) {
	obj, err := c.Request("deleteProjectInvitation", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}

// CreateProject represents the paramter of CreateProject
type CreateProjectParameter struct {
	// account who will be Admin for the project
	Account NullString
	// display text of the project
	DisplayText NullString
	// domain ID of the account owning a project
	DomainId ID
	// name of the project
	Name NullString
}

func NewCreateProjectParameter(displaytext string, name string) (p *CreateProjectParameter) {
	p = new(CreateProjectParameter)
	p.DisplayText.Set(displaytext)
	p.Name.Set(name)
	return p
}

// Creates a project
func (c *Client) CreateProject(p *CreateProjectParameter) (*Project, error) {
	obj, err := c.Request("createProject", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Project), err
}
