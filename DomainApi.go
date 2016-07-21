package cloudstack


// DeleteDomain represents the paramter of DeleteDomain
type DeleteDomainParameter struct {
	// true if all domain resources (child domains, accounts) have to be cleaned up,
	// false otherwise
	Cleanup NullBool
	// ID of domain to delete
	Id ID
}

func NewDeleteDomainParameter(id string) (p *DeleteDomainParameter) {
	p = new(DeleteDomainParameter)
	p.Id.Set(id)
	return p
}

// Deletes a specified domain
func (c *Client) DeleteDomain(p *DeleteDomainParameter) (*Result, error) {
	obj, err := c.Request("deleteDomain", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// CreateDomain represents the paramter of CreateDomain
type CreateDomainParameter struct {
	// Domain UUID, required for adding domain from another Region
	DomainId ID
	// creates domain with this name
	Name NullString
	// Network domain for networks in the domain
	NetworkDomain NullString
	// assigns new domain a parent domain by domain ID of the parent.  If no parent
	// domain is specied, the ROOT domain is assumed.
	ParentDomainId ID
}

func NewCreateDomainParameter(name string) (p *CreateDomainParameter) {
	p = new(CreateDomainParameter)
	p.Name.Set(name)
	return p
}

// Creates a domain
func (c *Client) CreateDomain(p *CreateDomainParameter) (*Domain, error) {
	obj, err := c.Request("createDomain", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Domain), err
}


// ListDomains represents the paramter of ListDomains
type ListDomainsParameter struct {
	// List domain by domain ID.
	Id ID
	// List by keyword
	Keyword NullString
	// List domains by domain level.
	Level NullNumber
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll NullBool
	// List domain by domain name.
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
}

func NewListDomainsParameter() (p *ListDomainsParameter) {
	p = new(ListDomainsParameter)
	return p
}

// Lists domains and provides detailed information for listed domains
func (c *Client) ListDomains(p *ListDomainsParameter) ([]*Domain, error) {
	obj, err := c.Request("listDomains", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Domain), err
}


// UpdateDomain represents the paramter of UpdateDomain
type UpdateDomainParameter struct {
	// ID of domain to update
	Id ID
	// updates domain with this name
	Name NullString
	// Network domain for the domain's networks; empty string will update domainName
	// with NULL value
	NetworkDomain NullString
}

func NewUpdateDomainParameter(id string) (p *UpdateDomainParameter) {
	p = new(UpdateDomainParameter)
	p.Id.Set(id)
	return p
}

// Updates a domain with a new name
func (c *Client) UpdateDomain(p *UpdateDomainParameter) (*Domain, error) {
	obj, err := c.Request("updateDomain", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Domain), err
}


// ListDomainChildren represents the paramter of ListDomainChildren
type ListDomainChildrenParameter struct {
	// list children domain by parent domain ID.
	Id ID
	// to return the entire tree, use the value "true". To return the first level
	// children, use the value "false".
	IsRecursive NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll NullBool
	// list children domains by name
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
}

func NewListDomainChildrenParameter() (p *ListDomainChildrenParameter) {
	p = new(ListDomainChildrenParameter)
	return p
}

// Lists all children domains belonging to a specified domain
func (c *Client) ListDomainChildren(p *ListDomainChildrenParameter) ([]*Domain, error) {
	obj, err := c.Request("listDomainChildren", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Domain), err
}

