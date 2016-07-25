package cloudstack


// ListStorageProviders represents the paramter of ListStorageProviders
type ListStorageProvidersParameter struct {
	// List by keyword
	Keyword  NullString
	Page     NullNumber
	PageSize NullNumber
	// the type of storage provider: either primary or image
	Type NullString
}

func NewListStorageProvidersParameter(typ string) (p *ListStorageProvidersParameter) {
	p = new(ListStorageProvidersParameter)
	p.Type.Set(typ)
	return p
}

// Lists storage providers.
func (c *Client) ListStorageProviders(p *ListStorageProvidersParameter) ([]*StorageProvider, error) {
	obj, err := c.Request("listStorageProviders", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*StorageProvider), err
}

