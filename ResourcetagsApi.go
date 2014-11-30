package cloudstack


// DeleteTags represents the paramter of DeleteTags
type DeleteTagsParameter struct {
	// Delete tags for resource id(s)
	ResourceIds []string
	// Delete tag by resource type
	ResourceType NullString
	// Delete tags matching key/value pairs
	Tags map[string]string
}

func NewDeleteTagsParameter(resourceids []string, resourcetype string) (p *DeleteTagsParameter) {
	p = new(DeleteTagsParameter)
	p.ResourceIds = resourceids
	p.ResourceType.Set(resourcetype)
	return p
}

// Deleting resource tag(s)
func (c *Client) DeleteTags(p *DeleteTagsParameter) (*Result, error) {
	obj, err := c.Request("deleteTags", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// ListTags represents the paramter of ListTags
type ListTagsParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// list by customer name
	Customer NullString
	// list only resources belonging to the domain specified
	DomainId ID
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// list by key
	Key NullString
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
	// list by resource id
	ResourceId ID
	// list by resource type
	ResourceType NullString
	// list by value
	Value NullString
}

func NewListTagsParameter() (p *ListTagsParameter) {
	p = new(ListTagsParameter)
	return p
}

// List resource tag(s)
func (c *Client) ListTags(p *ListTagsParameter) ([]*Tag, error) {
	obj, err := c.Request("listTags", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Tag), err
}


// CreateTags represents the paramter of CreateTags
type CreateTagsParameter struct {
	// identifies client specific tag. When the value is not null, the tag can't be
	// used by cloudStack code internally
	Customer NullString
	// list of resources to create the tags for
	ResourceIds []string
	// type of the resource
	ResourceType NullString
	// Map of tags (key/value pairs)
	Tags map[string]string
}

func NewCreateTagsParameter(resourceids []string, resourcetype string, tags map[string]string) (p *CreateTagsParameter) {
	p = new(CreateTagsParameter)
	p.ResourceIds = resourceids
	p.ResourceType.Set(resourcetype)
	p.Tags = tags
	return p
}

// Creates resource tag(s)
func (c *Client) CreateTags(p *CreateTagsParameter) (*Result, error) {
	obj, err := c.Request("createTags", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}

