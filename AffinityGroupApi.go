package cloudstack


// CreateAffinityGroup represents the paramter of CreateAffinityGroup
type CreateAffinityGroupParameter struct {
	// an account for the affinity group. Must be used with domainId.
	Account NullString
	// optional description of the affinity group
	Description NullString
	// domainId of the account owning the affinity group
	DomainId ID
	// name of the affinity group
	Name NullString
	// Type of the affinity group from the available affinity/anti-affinity group
	// types
	Type NullString
}

func NewCreateAffinityGroupParameter(name string, typ string) (p *CreateAffinityGroupParameter) {
	p = new(CreateAffinityGroupParameter)
	p.Name.Set(name)
	p.Type.Set(typ)
	return p
}

// Creates an affinity/anti-affinity group
func (c *Client) CreateAffinityGroup(p *CreateAffinityGroupParameter) (*AffinityGroup, error) {
	obj, err := c.Request("createAffinityGroup", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*AffinityGroup), err
}


// ListAffinityGroupTypes represents the paramter of ListAffinityGroupTypes
type ListAffinityGroupTypesParameter struct {
	// List by keyword
	Keyword  NullString
	Page     NullNumber
	PageSize NullNumber
}

func NewListAffinityGroupTypesParameter() (p *ListAffinityGroupTypesParameter) {
	p = new(ListAffinityGroupTypesParameter)
	return p
}

// Lists affinity group types available
func (c *Client) ListAffinityGroupTypes(p *ListAffinityGroupTypesParameter) ([]*AffinityGroupType, error) {
	obj, err := c.Request("listAffinityGroupTypes", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*AffinityGroupType), err
}


// UpdateVMAffinityGroup represents the paramter of UpdateVMAffinityGroup
type UpdateVMAffinityGroupParameter struct {
	// comma separated list of affinity groups id that are going to be applied to
	// the virtual machine. Should be passed only when vm is created from a zone
	// with Basic Network support. Mutually exclusive with securitygroupnames
	// parameter
	AffinityGroupIds []string
	// comma separated list of affinity groups names that are going to be applied to
	// the virtual machine. Should be passed only when vm is created from a zone
	// with Basic Network support. Mutually exclusive with securitygroupids
	// parameter
	AffinityGroupNames []string
	// The ID of the virtual machine
	Id ID
}

func NewUpdateVMAffinityGroupParameter(id string) (p *UpdateVMAffinityGroupParameter) {
	p = new(UpdateVMAffinityGroupParameter)
	p.Id.Set(id)
	return p
}

// Updates the affinity/anti-affinity group associations of a virtual machine.
// The VM has to be stopped and restarted for the new properties to take effect.
func (c *Client) UpdateVMAffinityGroup(p *UpdateVMAffinityGroupParameter) (*VirtualMachine, error) {
	obj, err := c.Request("updateVMAffinityGroup", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// ListAffinityGroups represents the paramter of ListAffinityGroups
type ListAffinityGroupsParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// list only resources belonging to the domain specified
	DomainId ID
	// list the affinity group by the id provided
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
	// lists affinity groups by name
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
	// lists affinity groups by type
	Type NullString
	// lists affinity groups by virtual machine id
	VirtualMachineId ID
}

func NewListAffinityGroupsParameter() (p *ListAffinityGroupsParameter) {
	p = new(ListAffinityGroupsParameter)
	return p
}

// Lists affinity groups
func (c *Client) ListAffinityGroups(p *ListAffinityGroupsParameter) ([]*AffinityGroup, error) {
	obj, err := c.Request("listAffinityGroups", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*AffinityGroup), err
}


// DeleteAffinityGroup represents the paramter of DeleteAffinityGroup
type DeleteAffinityGroupParameter struct {
	// the account of the affinity group. Must be specified with domain ID
	Account NullString
	// the domain ID of account owning the affinity group
	DomainId ID
	// The ID of the affinity group. Mutually exclusive with name parameter
	Id ID
	// The name of the affinity group. Mutually exclusive with id parameter
	Name NullString
}

func NewDeleteAffinityGroupParameter() (p *DeleteAffinityGroupParameter) {
	p = new(DeleteAffinityGroupParameter)
	return p
}

// Deletes affinity group
func (c *Client) DeleteAffinityGroup(p *DeleteAffinityGroupParameter) (*Result, error) {
	obj, err := c.Request("deleteAffinityGroup", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}

