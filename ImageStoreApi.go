package cloudstack

// UpdateCloudToUseObjectStore represents the paramter of UpdateCloudToUseObjectStore
type UpdateCloudToUseObjectStoreParameter struct {
	// the details for the image store. Example:
	// details[0].key=accesskey&details[0].value=s389ddssaa&details[1].key=secretkey&details[1].value=8dshfsss
	Details map[string]string
	// the name for the image store
	Name NullString
	// the image store provider name
	Provider NullString
	// the URL for the image store
	Url NullString
}

func NewUpdateCloudToUseObjectStoreParameter(provider string) (p *UpdateCloudToUseObjectStoreParameter) {
	p = new(UpdateCloudToUseObjectStoreParameter)
	p.Provider.Set(provider)
	return p
}

// Migrate current NFS secondary storages to use object store.
func (c *Client) UpdateCloudToUseObjectStore(p *UpdateCloudToUseObjectStoreParameter) (*ImageStore, error) {
	obj, err := c.Request("updateCloudToUseObjectStore", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*ImageStore), err
}

// AddSecondaryStorage represents the paramter of AddSecondaryStorage
type AddSecondaryStorageParameter struct {
	// the URL for the secondary storage
	Url NullString
	// the Zone ID for the secondary storage
	ZoneId ID
}

func NewAddSecondaryStorageParameter(url string) (p *AddSecondaryStorageParameter) {
	p = new(AddSecondaryStorageParameter)
	p.Url.Set(url)
	return p
}

// Adds secondary storage.
func (c *Client) AddSecondaryStorage(p *AddSecondaryStorageParameter) (*ImageStore, error) {
	obj, err := c.Request("addSecondaryStorage", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*ImageStore), err
}

// ListImageStores represents the paramter of ListImageStores
type ListImageStoresParameter struct {
	// the ID of the storage pool
	Id ID
	// List by keyword
	Keyword NullString
	// the name of the image store
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
	// the image store protocol
	Protocol NullString
	// the image store provider
	Provider NullString
	// the Zone ID for the image store
	ZoneId ID
}

func NewListImageStoresParameter() (p *ListImageStoresParameter) {
	p = new(ListImageStoresParameter)
	return p
}

// Lists image stores.
func (c *Client) ListImageStores(p *ListImageStoresParameter) ([]*ImageStore, error) {
	obj, err := c.Request("listImageStores", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*ImageStore), err
}

// DeleteSecondaryStagingStore represents the paramter of DeleteSecondaryStagingStore
type DeleteSecondaryStagingStoreParameter struct {
	// the staging store ID
	Id ID
}

func NewDeleteSecondaryStagingStoreParameter(id string) (p *DeleteSecondaryStagingStoreParameter) {
	p = new(DeleteSecondaryStagingStoreParameter)
	p.Id.Set(id)
	return p
}

// Deletes a secondary staging store .
func (c *Client) DeleteSecondaryStagingStore(p *DeleteSecondaryStagingStoreParameter) (*Result, error) {
	obj, err := c.Request("deleteSecondaryStagingStore", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}

// ListSecondaryStagingStores represents the paramter of ListSecondaryStagingStores
type ListSecondaryStagingStoresParameter struct {
	// the ID of the staging store
	Id ID
	// List by keyword
	Keyword NullString
	// the name of the staging store
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
	// the staging store protocol
	Protocol NullString
	// the staging store provider
	Provider NullString
	// the Zone ID for the staging store
	ZoneId ID
}

func NewListSecondaryStagingStoresParameter() (p *ListSecondaryStagingStoresParameter) {
	p = new(ListSecondaryStagingStoresParameter)
	return p
}

// Lists secondary staging stores.
func (c *Client) ListSecondaryStagingStores(p *ListSecondaryStagingStoresParameter) ([]*ImageStore, error) {
	obj, err := c.Request("listSecondaryStagingStores", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*ImageStore), err
}

// AddImageStore represents the paramter of AddImageStore
type AddImageStoreParameter struct {
	// the details for the image store. Example:
	// details[0].key=accesskey&details[0].value=s389ddssaa&details[1].key=secretkey&details[1].value=8dshfsss
	Details map[string]string
	// the name for the image store
	Name NullString
	// the image store provider name
	Provider NullString
	// the URL for the image store
	Url NullString
	// the Zone ID for the image store
	ZoneId ID
}

func NewAddImageStoreParameter(provider string) (p *AddImageStoreParameter) {
	p = new(AddImageStoreParameter)
	p.Provider.Set(provider)
	return p
}

// Adds backup image store.
func (c *Client) AddImageStore(p *AddImageStoreParameter) (*ImageStore, error) {
	obj, err := c.Request("addImageStore", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*ImageStore), err
}

// CreateSecondaryStagingStore represents the paramter of CreateSecondaryStagingStore
type CreateSecondaryStagingStoreParameter struct {
	// the details for the staging store
	Details map[string]string
	// the staging store provider name
	Provider NullString
	// the scope of the staging store: zone only for now
	Scope NullString
	// the URL for the staging store
	Url NullString
	// the Zone ID for the staging store
	ZoneId ID
}

func NewCreateSecondaryStagingStoreParameter(url string) (p *CreateSecondaryStagingStoreParameter) {
	p = new(CreateSecondaryStagingStoreParameter)
	p.Url.Set(url)
	return p
}

// create secondary staging store.
func (c *Client) CreateSecondaryStagingStore(p *CreateSecondaryStagingStoreParameter) (*ImageStore, error) {
	obj, err := c.Request("createSecondaryStagingStore", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*ImageStore), err
}

// DeleteImageStore represents the paramter of DeleteImageStore
type DeleteImageStoreParameter struct {
	// the image store ID
	Id ID
}

func NewDeleteImageStoreParameter(id string) (p *DeleteImageStoreParameter) {
	p = new(DeleteImageStoreParameter)
	p.Id.Set(id)
	return p
}

// Deletes an image store .
func (c *Client) DeleteImageStore(p *DeleteImageStoreParameter) (*Result, error) {
	obj, err := c.Request("deleteImageStore", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}
