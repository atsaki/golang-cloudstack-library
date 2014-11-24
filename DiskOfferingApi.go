package cloudstack


// UpdateDiskOffering represents the paramter of UpdateDiskOffering
type UpdateDiskOfferingParameter struct {
	// an optional field, whether to display the offering to the end user or not.
	DisplayOffering NullBool
	// updates alternate display text of the disk offering with this value
	DisplayText NullString
	// ID of the disk offering
	Id ID
	// updates name of the disk offering with this value
	Name NullString
	// sort key of the disk offering, integer
	SortKey NullNumber
}

func NewUpdateDiskOfferingParameter(id string) (p *UpdateDiskOfferingParameter) {
	p = new(UpdateDiskOfferingParameter)
	p.Id.Set(id)
	return p
}

// Updates a disk offering.
func (c *Client) UpdateDiskOffering(p *UpdateDiskOfferingParameter) (obj *DiskOffering, err error) {
	err = c.request(p, &obj)
	return obj, err
}


// ListDiskOfferings represents the paramter of ListDiskOfferings
type ListDiskOfferingsParameter struct {
	// the ID of the domain of the disk offering.
	DomainId ID
	// ID of the disk offering
	Id ID
	// List by keyword
	Keyword NullString
	// name of the disk offering
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
}

func NewListDiskOfferingsParameter() (p *ListDiskOfferingsParameter) {
	p = new(ListDiskOfferingsParameter)
	return p
}

// Lists all available disk offerings.
func (c *Client) ListDiskOfferings(p *ListDiskOfferingsParameter) (objs []DiskOffering, err error) {
	err = c.request(p, &objs)
	return objs, err
}


// DeleteDiskOffering represents the paramter of DeleteDiskOffering
type DeleteDiskOfferingParameter struct {
	// ID of the disk offering
	Id ID
}

func NewDeleteDiskOfferingParameter(id string) (p *DeleteDiskOfferingParameter) {
	p = new(DeleteDiskOfferingParameter)
	p.Id.Set(id)
	return p
}

// Updates a disk offering.
func (c *Client) DeleteDiskOffering(p *DeleteDiskOfferingParameter) (obj *Result, err error) {
	err = c.request(p, &obj)
	return obj, err
}


// CreateDiskOffering represents the paramter of CreateDiskOffering
type CreateDiskOfferingParameter struct {
	// bytes read rate of the disk offering
	BytesReadRate NullNumber
	// bytes write rate of the disk offering
	BytesWriteRate NullNumber
	// whether disk offering size is custom or not
	Customized NullBool
	// whether disk offering iops is custom or not
	CustomizedIops NullBool
	// size of the disk offering in GB
	DiskSize NullNumber
	// an optional field, whether to display the offering to the end user or not.
	DisplayOffering NullBool
	// alternate display text of the disk offering
	DisplayText NullString
	// the ID of the containing domain, null for public offerings
	DomainId ID
	// Hypervisor snapshot reserve space as a percent of a volume (for managed
	// storage using Xen or VMware)
	HypervisorSnapshotReserve NullNumber
	// io requests read rate of the disk offering
	IopsReadRate NullNumber
	// io requests write rate of the disk offering
	IopsWriteRate NullNumber
	// max iops of the disk offering
	MaxIops NullNumber
	// min iops of the disk offering
	MinIops NullNumber
	// name of the disk offering
	Name NullString
	// the storage type of the disk offering. Values are local and shared.
	StorageType NullString
	// tags for the disk offering
	Tags NullString
}

func NewCreateDiskOfferingParameter(displaytext string, name string) (p *CreateDiskOfferingParameter) {
	p = new(CreateDiskOfferingParameter)
	p.DisplayText.Set(displaytext)
	p.Name.Set(name)
	return p
}

// Creates a disk offering.
func (c *Client) CreateDiskOffering(p *CreateDiskOfferingParameter) (obj *DiskOffering, err error) {
	err = c.request(p, &obj)
	return obj, err
}

