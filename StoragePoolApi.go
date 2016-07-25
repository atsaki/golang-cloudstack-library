package cloudstack


// EnableStorageMaintenance represents the paramter of EnableStorageMaintenance
type EnableStorageMaintenanceParameter struct {
	// Primary storage ID
	Id ID
}

func NewEnableStorageMaintenanceParameter(id string) (p *EnableStorageMaintenanceParameter) {
	p = new(EnableStorageMaintenanceParameter)
	p.Id.Set(id)
	return p
}

// Puts storage pool into maintenance state
func (c *Client) EnableStorageMaintenance(p *EnableStorageMaintenanceParameter) (*StoragePool, error) {
	obj, err := c.Request("enableStorageMaintenance", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*StoragePool), err
}


// UpdateStoragePool represents the paramter of UpdateStoragePool
type UpdateStoragePoolParameter struct {
	// bytes CloudStack can provision from this storage pool
	CapacityBytes NullNumber
	// IOPS CloudStack can provision from this storage pool
	CapacityIops NullNumber
	// the Id of the storage pool
	Id ID
	// comma-separated list of tags for the storage pool
	Tags []string
}

func NewUpdateStoragePoolParameter(id string) (p *UpdateStoragePoolParameter) {
	p = new(UpdateStoragePoolParameter)
	p.Id.Set(id)
	return p
}

// Updates a storage pool.
func (c *Client) UpdateStoragePool(p *UpdateStoragePoolParameter) (*StoragePool, error) {
	obj, err := c.Request("updateStoragePool", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*StoragePool), err
}


// CancelStorageMaintenance represents the paramter of CancelStorageMaintenance
type CancelStorageMaintenanceParameter struct {
	// the primary storage ID
	Id ID
}

func NewCancelStorageMaintenanceParameter(id string) (p *CancelStorageMaintenanceParameter) {
	p = new(CancelStorageMaintenanceParameter)
	p.Id.Set(id)
	return p
}

// Cancels maintenance for primary storage
func (c *Client) CancelStorageMaintenance(p *CancelStorageMaintenanceParameter) (*StoragePool, error) {
	obj, err := c.Request("cancelStorageMaintenance", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*StoragePool), err
}


// CreateStoragePool represents the paramter of CreateStoragePool
type CreateStoragePoolParameter struct {
	// bytes CloudStack can provision from this storage pool
	CapacityBytes NullNumber
	// IOPS CloudStack can provision from this storage pool
	CapacityIops NullNumber
	// the cluster ID for the storage pool
	ClusterId ID
	// the details for the storage pool
	Details map[string]string
	// hypervisor type of the hosts in zone that will be attached to this storage
	// pool. KVM, VMware supported as of now.
	Hypervisor NullString
	// whether the storage should be managed by CloudStack
	Managed NullBool
	// the name for the storage pool
	Name NullString
	// the Pod ID for the storage pool
	PodId ID
	// the storage provider name
	Provider NullString
	// the scope of the storage: cluster or zone
	Scope NullString
	// the tags for the storage pool
	Tags NullString
	// the URL of the storage pool
	Url NullString
	// the Zone ID for the storage pool
	ZoneId ID
}

func NewCreateStoragePoolParameter(name string, url string, zoneid string) (p *CreateStoragePoolParameter) {
	p = new(CreateStoragePoolParameter)
	p.Name.Set(name)
	p.Url.Set(url)
	p.ZoneId.Set(zoneid)
	return p
}

// Creates a storage pool.
func (c *Client) CreateStoragePool(p *CreateStoragePoolParameter) (*StoragePool, error) {
	obj, err := c.Request("createStoragePool", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*StoragePool), err
}


// ListStoragePools represents the paramter of ListStoragePools
type ListStoragePoolsParameter struct {
	// list storage pools belongig to the specific cluster
	ClusterId ID
	// the ID of the storage pool
	Id ID
	// the IP address for the storage pool
	IpAddress NullString
	// List by keyword
	Keyword NullString
	// the name of the storage pool
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
	// the storage pool path
	Path NullString
	// the Pod ID for the storage pool
	PodId ID
	// the ID of the storage pool
	Scope NullString
	// the Zone ID for the storage pool
	ZoneId ID
}

func NewListStoragePoolsParameter() (p *ListStoragePoolsParameter) {
	p = new(ListStoragePoolsParameter)
	return p
}

// Lists storage pools.
func (c *Client) ListStoragePools(p *ListStoragePoolsParameter) ([]*StoragePool, error) {
	obj, err := c.Request("listStoragePools", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*StoragePool), err
}


// DeleteStoragePool represents the paramter of DeleteStoragePool
type DeleteStoragePoolParameter struct {
	// Force destroy storage pool (force expunge volumes in Destroyed state as a
	// part of pool removal)
	Forced NullBool
	// Storage pool id
	Id ID
}

func NewDeleteStoragePoolParameter(id string) (p *DeleteStoragePoolParameter) {
	p = new(DeleteStoragePoolParameter)
	p.Id.Set(id)
	return p
}

// Deletes a storage pool.
func (c *Client) DeleteStoragePool(p *DeleteStoragePoolParameter) (*Result, error) {
	obj, err := c.Request("deleteStoragePool", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// FindStoragePoolsForMigration represents the paramter of FindStoragePoolsForMigration
type FindStoragePoolsForMigrationParameter struct {
	// the ID of the volume
	Id ID
	// List by keyword
	Keyword  NullString
	Page     NullNumber
	PageSize NullNumber
}

func NewFindStoragePoolsForMigrationParameter(id string) (p *FindStoragePoolsForMigrationParameter) {
	p = new(FindStoragePoolsForMigrationParameter)
	p.Id.Set(id)
	return p
}

// Lists storage pools available for migration of a volume.
func (c *Client) FindStoragePoolsForMigration(p *FindStoragePoolsForMigrationParameter) (*StoragePool, error) {
	obj, err := c.Request("findStoragePoolsForMigration", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*StoragePool), err
}

