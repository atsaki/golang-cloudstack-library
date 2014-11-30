package cloudstack


// AuthorizeSecurityGroupIngress represents the paramter of AuthorizeSecurityGroupIngress
type AuthorizeSecurityGroupIngressParameter struct {
	// an optional account for the security group. Must be used with domainId.
	Account NullString
	// the cidr list associated
	CidrList []string
	// an optional domainId for the security group. If the account parameter is
	// used, domainId must also be used.
	DomainId ID
	// end port for this ingress rule
	EndPort NullNumber
	// error code for this icmp message
	IcmpCode NullNumber
	// type of the icmp message being sent
	IcmpType NullNumber
	// an optional project of the security group
	ProjectId ID
	// TCP is default. UDP is the other supported protocol
	Protocol NullString
	// The ID of the security group. Mutually exclusive with securityGroupName
	// parameter
	SecurityGroupId ID
	// The name of the security group. Mutually exclusive with securityGroupName
	// parameter
	SecurityGroupName NullString
	// start port for this ingress rule
	StartPort NullNumber
	// user to security group mapping
	UserSecurityGroupList map[string]string
}

func NewAuthorizeSecurityGroupIngressParameter() (p *AuthorizeSecurityGroupIngressParameter) {
	p = new(AuthorizeSecurityGroupIngressParameter)
	return p
}

// Authorizes a particular ingress rule for this security group
func (c *Client) AuthorizeSecurityGroupIngress(p *AuthorizeSecurityGroupIngressParameter) (*SecurityGroupIngress, error) {
	obj, err := c.Request("authorizeSecurityGroupIngress", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*SecurityGroupIngress), err
}


// CreateSecurityGroup represents the paramter of CreateSecurityGroup
type CreateSecurityGroupParameter struct {
	// an optional account for the security group. Must be used with domainId.
	Account NullString
	// the description of the security group
	Description NullString
	// an optional domainId for the security group. If the account parameter is
	// used, domainId must also be used.
	DomainId ID
	// name of the security group
	Name NullString
	// Create security group for project
	ProjectId ID
}

func NewCreateSecurityGroupParameter(name string) (p *CreateSecurityGroupParameter) {
	p = new(CreateSecurityGroupParameter)
	p.Name.Set(name)
	return p
}

// Creates a security group
func (c *Client) CreateSecurityGroup(p *CreateSecurityGroupParameter) (*SecurityGroup, error) {
	obj, err := c.Request("createSecurityGroup", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*SecurityGroup), err
}


// RevokeSecurityGroupEgress represents the paramter of RevokeSecurityGroupEgress
type RevokeSecurityGroupEgressParameter struct {
	// The ID of the egress rule
	Id ID
}

func NewRevokeSecurityGroupEgressParameter(id string) (p *RevokeSecurityGroupEgressParameter) {
	p = new(RevokeSecurityGroupEgressParameter)
	p.Id.Set(id)
	return p
}

// Deletes a particular egress rule from this security group
func (c *Client) RevokeSecurityGroupEgress(p *RevokeSecurityGroupEgressParameter) (*Result, error) {
	obj, err := c.Request("revokeSecurityGroupEgress", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// DeleteSecurityGroup represents the paramter of DeleteSecurityGroup
type DeleteSecurityGroupParameter struct {
	// the account of the security group. Must be specified with domain ID
	Account NullString
	// the domain ID of account owning the security group
	DomainId ID
	// The ID of the security group. Mutually exclusive with name parameter
	Id ID
	// The ID of the security group. Mutually exclusive with id parameter
	Name NullString
	// the project of the security group
	ProjectId ID
}

func NewDeleteSecurityGroupParameter() (p *DeleteSecurityGroupParameter) {
	p = new(DeleteSecurityGroupParameter)
	return p
}

// Deletes security group
func (c *Client) DeleteSecurityGroup(p *DeleteSecurityGroupParameter) (*Result, error) {
	obj, err := c.Request("deleteSecurityGroup", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// AuthorizeSecurityGroupEgress represents the paramter of AuthorizeSecurityGroupEgress
type AuthorizeSecurityGroupEgressParameter struct {
	// an optional account for the security group. Must be used with domainId.
	Account NullString
	// the cidr list associated
	CidrList []string
	// an optional domainId for the security group. If the account parameter is
	// used, domainId must also be used.
	DomainId ID
	// end port for this egress rule
	EndPort NullNumber
	// error code for this icmp message
	IcmpCode NullNumber
	// type of the icmp message being sent
	IcmpType NullNumber
	// an optional project of the security group
	ProjectId ID
	// TCP is default. UDP is the other supported protocol
	Protocol NullString
	// The ID of the security group. Mutually exclusive with securityGroupName
	// parameter
	SecurityGroupId ID
	// The name of the security group. Mutually exclusive with securityGroupName
	// parameter
	SecurityGroupName NullString
	// start port for this egress rule
	StartPort NullNumber
	// user to security group mapping
	UserSecurityGroupList map[string]string
}

func NewAuthorizeSecurityGroupEgressParameter() (p *AuthorizeSecurityGroupEgressParameter) {
	p = new(AuthorizeSecurityGroupEgressParameter)
	return p
}

// Authorizes a particular egress rule for this security group
func (c *Client) AuthorizeSecurityGroupEgress(p *AuthorizeSecurityGroupEgressParameter) (*SecurityGroupEgress, error) {
	obj, err := c.Request("authorizeSecurityGroupEgress", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*SecurityGroupEgress), err
}


// ListSecurityGroups represents the paramter of ListSecurityGroups
type ListSecurityGroupsParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// list only resources belonging to the domain specified
	DomainId ID
	// list the security group by the id provided
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
	// list objects by project
	ProjectId ID
	// lists security groups by name
	SecurityGroupName NullString
	// List resources by tags (key/value pairs)
	Tags map[string]string
	// lists security groups by virtual machine id
	VirtualMachineId ID
}

func NewListSecurityGroupsParameter() (p *ListSecurityGroupsParameter) {
	p = new(ListSecurityGroupsParameter)
	return p
}

// Lists security groups
func (c *Client) ListSecurityGroups(p *ListSecurityGroupsParameter) ([]*SecurityGroup, error) {
	obj, err := c.Request("listSecurityGroups", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*SecurityGroup), err
}


// RevokeSecurityGroupIngress represents the paramter of RevokeSecurityGroupIngress
type RevokeSecurityGroupIngressParameter struct {
	// The ID of the ingress rule
	Id ID
}

func NewRevokeSecurityGroupIngressParameter(id string) (p *RevokeSecurityGroupIngressParameter) {
	p = new(RevokeSecurityGroupIngressParameter)
	p.Id.Set(id)
	return p
}

// Deletes a particular ingress rule from this security group
func (c *Client) RevokeSecurityGroupIngress(p *RevokeSecurityGroupIngressParameter) (*Result, error) {
	obj, err := c.Request("revokeSecurityGroupIngress", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}

