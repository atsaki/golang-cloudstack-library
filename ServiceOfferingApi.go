package cloudstack


// UpdateServiceOffering represents the paramter of UpdateServiceOffering
type UpdateServiceOfferingParameter struct {
	// the display text of the service offering to be updated
	DisplayText NullString
	// the ID of the service offering to be updated
	Id ID
	// the name of the service offering to be updated
	Name NullString
	// sort key of the service offering, integer
	SortKey NullNumber
}

func NewUpdateServiceOfferingParameter(id string) (p *UpdateServiceOfferingParameter) {
	p = new(UpdateServiceOfferingParameter)
	p.Id.Set(id)
	return p
}

// Updates a service offering.
func (c *Client) UpdateServiceOffering(p *UpdateServiceOfferingParameter) (*ServiceOffering, error) {
	obj, err := c.Request("updateServiceOffering", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*ServiceOffering), err
}


// CreateServiceOffering represents the paramter of CreateServiceOffering
type CreateServiceOfferingParameter struct {
	// bytes read rate of the disk offering
	BytesReadRate NullNumber
	// bytes write rate of the disk offering
	BytesWriteRate NullNumber
	// the CPU number of the service offering
	CpuNumber NullNumber
	// the CPU speed of the service offering in MHz.
	CpuSpeed NullNumber
	// whether compute offering iops is custom or not
	CustomizedIops NullBool
	// The deployment planner heuristics used to deploy a VM of this offering. If
	// null, value of global config vm.deployment.planner is used
	DeploymentPlanner NullString
	// the display text of the service offering
	DisplayText NullString
	// the ID of the containing domain, null for public offerings
	DomainId ID
	// the host tag for this service offering.
	HostTags NullString
	// Hypervisor snapshot reserve space as a percent of a volume (for managed
	// storage using Xen or VMware)
	HypervisorSnapshotReserve NullNumber
	// io requests read rate of the disk offering
	IopsReadRate NullNumber
	// io requests write rate of the disk offering
	IopsWriteRate NullNumber
	// is this a system vm offering
	IsSystem NullBool
	// true if the virtual machine needs to be volatile so that on every reboot of
	// VM, original root disk is dettached then destroyed and a fresh root disk is
	// created and attached to VM
	IsVolatile NullBool
	// restrict the CPU usage to committed service offering
	LimitCpuUse NullBool
	// max iops of the compute offering
	MaxIops NullNumber
	// the total memory of the service offering in MB
	Memory NullNumber
	// min iops of the compute offering
	MinIops NullNumber
	// the name of the service offering
	Name NullString
	// data transfer rate in megabits per second allowed. Supported only for
	// non-System offering and system offerings having "domainrouter" systemvmtype
	NetworkRate NullNumber
	// the HA for the service offering
	OfferHa NullBool
	// details for planner, used to store specific parameters
	ServiceOfferingDetails map[string]string
	// the storage type of the service offering. Values are local and shared.
	StorageType NullString
	// the system VM type. Possible types are "domainrouter", "consoleproxy" and
	// "secondarystoragevm".
	SystemVmType NullString
	// the tags for this service offering.
	Tags NullString
}

func NewCreateServiceOfferingParameter(displaytext string, name string) (p *CreateServiceOfferingParameter) {
	p = new(CreateServiceOfferingParameter)
	p.DisplayText.Set(displaytext)
	p.Name.Set(name)
	return p
}

// Creates a service offering.
func (c *Client) CreateServiceOffering(p *CreateServiceOfferingParameter) (*ServiceOffering, error) {
	obj, err := c.Request("createServiceOffering", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*ServiceOffering), err
}


// DeleteServiceOffering represents the paramter of DeleteServiceOffering
type DeleteServiceOfferingParameter struct {
	// the ID of the service offering
	Id ID
}

func NewDeleteServiceOfferingParameter(id string) (p *DeleteServiceOfferingParameter) {
	p = new(DeleteServiceOfferingParameter)
	p.Id.Set(id)
	return p
}

// Deletes a service offering.
func (c *Client) DeleteServiceOffering(p *DeleteServiceOfferingParameter) (*Result, error) {
	obj, err := c.Request("deleteServiceOffering", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// ListServiceOfferings represents the paramter of ListServiceOfferings
type ListServiceOfferingsParameter struct {
	// the ID of the domain associated with the service offering
	DomainId ID
	// ID of the service offering
	Id ID
	// is this a system vm offering
	IsSystem NullBool
	// List by keyword
	Keyword NullString
	// name of the service offering
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
	// the system VM type. Possible types are "consoleproxy", "secondarystoragevm"
	// or "domainrouter".
	SystemVmType NullString
	// the ID of the virtual machine. Pass this in if you want to see the available
	// service offering that a virtual machine can be changed to.
	VirtualMachineId ID
}

func NewListServiceOfferingsParameter() (p *ListServiceOfferingsParameter) {
	p = new(ListServiceOfferingsParameter)
	return p
}

// Lists all available service offerings.
func (c *Client) ListServiceOfferings(p *ListServiceOfferingsParameter) ([]*ServiceOffering, error) {
	obj, err := c.Request("listServiceOfferings", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*ServiceOffering), err
}

