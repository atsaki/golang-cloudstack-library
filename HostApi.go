package cloudstack

// ReconnectHost represents the paramter of ReconnectHost
type ReconnectHostParameter struct {
	// the host ID
	Id ID
}

func NewReconnectHostParameter(id string) (p *ReconnectHostParameter) {
	p = new(ReconnectHostParameter)
	p.Id.Set(id)
	return p
}

// Reconnects a host.
func (c *Client) ReconnectHost(p *ReconnectHostParameter) (*Host, error) {
	obj, err := c.Request("reconnectHost", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Host), err
}

// CancelHostMaintenance represents the paramter of CancelHostMaintenance
type CancelHostMaintenanceParameter struct {
	// the host ID
	Id ID
}

func NewCancelHostMaintenanceParameter(id string) (p *CancelHostMaintenanceParameter) {
	p = new(CancelHostMaintenanceParameter)
	p.Id.Set(id)
	return p
}

// Cancels host maintenance.
func (c *Client) CancelHostMaintenance(p *CancelHostMaintenanceParameter) (*Host, error) {
	obj, err := c.Request("cancelHostMaintenance", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Host), err
}

// ReleaseDedicatedHost represents the paramter of ReleaseDedicatedHost
type ReleaseDedicatedHostParameter struct {
	// the ID of the host
	HostId ID
}

func NewReleaseDedicatedHostParameter(hostid string) (p *ReleaseDedicatedHostParameter) {
	p = new(ReleaseDedicatedHostParameter)
	p.HostId.Set(hostid)
	return p
}

// Release the dedication for host
func (c *Client) ReleaseDedicatedHost(p *ReleaseDedicatedHostParameter) (*Result, error) {
	obj, err := c.Request("releaseDedicatedHost", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}

// AddBaremetalHost represents the paramter of AddBaremetalHost
type AddBaremetalHostParameter struct {
	// Allocation state of this Host for allocation of new resources
	AllocationState NullString
	// the cluster ID for the host
	ClusterId ID
	// the cluster name for the host
	ClusterName NullString
	// list of tags to be added to the host
	HostTags []string
	// hypervisor type of the host
	Hypervisor NullString
	// ip address intentionally allocated to this host after provisioning
	IpAddress NullString
	// the password for the host
	Password NullString
	// the Pod ID for the host
	PodId ID
	// the host URL
	Url NullString
	// the username for the host
	UserName NullString
	// the Zone ID for the host
	ZoneId ID
}

func NewAddBaremetalHostParameter(hypervisor string, password string, podid string, url string, username string, zoneid string) (p *AddBaremetalHostParameter) {
	p = new(AddBaremetalHostParameter)
	p.Hypervisor.Set(hypervisor)
	p.Password.Set(password)
	p.PodId.Set(podid)
	p.Url.Set(url)
	p.UserName.Set(username)
	p.ZoneId.Set(zoneid)
	return p
}

// add a baremetal host
func (c *Client) AddBaremetalHost(p *AddBaremetalHostParameter) (*Host, error) {
	obj, err := c.Request("addBaremetalHost", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Host), err
}

// UpdateHostPassword represents the paramter of UpdateHostPassword
type UpdateHostPasswordParameter struct {
	// the cluster ID
	ClusterId ID
	// the host ID
	HostId ID
	// the new password for the host/cluster
	Password NullString
	// the username for the host/cluster
	UserName NullString
}

func NewUpdateHostPasswordParameter(password string, username string) (p *UpdateHostPasswordParameter) {
	p = new(UpdateHostPasswordParameter)
	p.Password.Set(password)
	p.UserName.Set(username)
	return p
}

// Update password of a host/pool on management server.
func (c *Client) UpdateHostPassword(p *UpdateHostPasswordParameter) (*Result, error) {
	obj, err := c.Request("updateHostPassword", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}

// ListDedicatedHosts represents the paramter of ListDedicatedHosts
type ListDedicatedHostsParameter struct {
	// the name of the account associated with the host. Must be used with domainId.
	Account NullString
	// list dedicated hosts by affinity group
	AffinityGroupId ID
	// the ID of the domain associated with the host
	DomainId ID
	// the ID of the host
	HostId ID
	// List by keyword
	Keyword  NullString
	Page     NullNumber
	PageSize NullNumber
}

func NewListDedicatedHostsParameter() (p *ListDedicatedHostsParameter) {
	p = new(ListDedicatedHostsParameter)
	return p
}

// Lists dedicated hosts.
func (c *Client) ListDedicatedHosts(p *ListDedicatedHostsParameter) ([]*DedicatedHost, error) {
	obj, err := c.Request("listDedicatedHosts", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*DedicatedHost), err
}

// ReleaseHostReservation represents the paramter of ReleaseHostReservation
type ReleaseHostReservationParameter struct {
	// the host ID
	Id ID
}

func NewReleaseHostReservationParameter(id string) (p *ReleaseHostReservationParameter) {
	p = new(ReleaseHostReservationParameter)
	p.Id.Set(id)
	return p
}

// Releases host reservation.
func (c *Client) ReleaseHostReservation(p *ReleaseHostReservationParameter) (*Result, error) {
	obj, err := c.Request("releaseHostReservation", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}

// AddHost represents the paramter of AddHost
type AddHostParameter struct {
	// Allocation state of this Host for allocation of new resources
	AllocationState NullString
	// the cluster ID for the host
	ClusterId ID
	// the cluster name for the host
	ClusterName NullString
	// list of tags to be added to the host
	HostTags []string
	// hypervisor type of the host
	Hypervisor NullString
	// the password for the host
	Password NullString
	// the Pod ID for the host
	PodId ID
	// the host URL
	Url NullString
	// the username for the host
	UserName NullString
	// the Zone ID for the host
	ZoneId ID
}

func NewAddHostParameter(hypervisor string, password string, podid string, url string, username string, zoneid string) (p *AddHostParameter) {
	p = new(AddHostParameter)
	p.Hypervisor.Set(hypervisor)
	p.Password.Set(password)
	p.PodId.Set(podid)
	p.Url.Set(url)
	p.UserName.Set(username)
	p.ZoneId.Set(zoneid)
	return p
}

// Adds a new host.
func (c *Client) AddHost(p *AddHostParameter) (*Host, error) {
	obj, err := c.Request("addHost", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Host), err
}

// UpdateHost represents the paramter of UpdateHost
type UpdateHostParameter struct {
	// Change resource state of host, valid values are [Enable, Disable]. Operation
	// may failed if host in states not allowing Enable/Disable
	AllocationState NullString
	// list of tags to be added to the host
	HostTags []string
	// the ID of the host to update
	Id ID
	// the id of Os category to update the host with
	OsCategoryId ID
	// the new uri for the secondary storage: nfs://host/path
	Url NullString
}

func NewUpdateHostParameter(id string) (p *UpdateHostParameter) {
	p = new(UpdateHostParameter)
	p.Id.Set(id)
	return p
}

// Updates a host.
func (c *Client) UpdateHost(p *UpdateHostParameter) (*Host, error) {
	obj, err := c.Request("updateHost", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Host), err
}

// FindHostsForMigration represents the paramter of FindHostsForMigration
type FindHostsForMigrationParameter struct {
	// List by keyword
	Keyword  NullString
	Page     NullNumber
	PageSize NullNumber
	// find hosts to which this VM can be migrated and flag the hosts with enough
	// CPU/RAM to host the VM
	VirtualMachineId ID
}

func NewFindHostsForMigrationParameter(virtualmachineid string) (p *FindHostsForMigrationParameter) {
	p = new(FindHostsForMigrationParameter)
	p.VirtualMachineId.Set(virtualmachineid)
	return p
}

// Find hosts suitable for migrating a virtual machine.
func (c *Client) FindHostsForMigration(p *FindHostsForMigrationParameter) (*Host, error) {
	obj, err := c.Request("findHostsForMigration", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Host), err
}

// PrepareHostForMaintenance represents the paramter of PrepareHostForMaintenance
type PrepareHostForMaintenanceParameter struct {
	// the host ID
	Id ID
}

func NewPrepareHostForMaintenanceParameter(id string) (p *PrepareHostForMaintenanceParameter) {
	p = new(PrepareHostForMaintenanceParameter)
	p.Id.Set(id)
	return p
}

// Prepares a host for maintenance.
func (c *Client) PrepareHostForMaintenance(p *PrepareHostForMaintenanceParameter) (*Host, error) {
	obj, err := c.Request("prepareHostForMaintenance", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Host), err
}

// DedicateHost represents the paramter of DedicateHost
type DedicateHostParameter struct {
	// the name of the account which needs dedication. Must be used with domainId.
	Account NullString
	// the ID of the containing domain
	DomainId ID
	// the ID of the host to update
	HostId ID
}

func NewDedicateHostParameter(domainid string, hostid string) (p *DedicateHostParameter) {
	p = new(DedicateHostParameter)
	p.DomainId.Set(domainid)
	p.HostId.Set(hostid)
	return p
}

// Dedicates a host.
func (c *Client) DedicateHost(p *DedicateHostParameter) (*DedicatedHost, error) {
	obj, err := c.Request("dedicateHost", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*DedicatedHost), err
}

// ListHostTags represents the paramter of ListHostTags
type ListHostTagsParameter struct {
	// List by keyword
	Keyword  NullString
	Page     NullNumber
	PageSize NullNumber
}

func NewListHostTagsParameter() (p *ListHostTagsParameter) {
	p = new(ListHostTagsParameter)
	return p
}

// Lists host tags
func (c *Client) ListHostTags(p *ListHostTagsParameter) ([]*HostTag, error) {
	obj, err := c.Request("listHostTags", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*HostTag), err
}

// ListHosts represents the paramter of ListHosts
type ListHostsParameter struct {
	// lists hosts existing in particular cluster
	ClusterId ID
	// comma separated list of host details requested, value can be a list of [ min,
	// all, capacity, events, stats]
	Details []string
	// if true, list only hosts dedicated to HA
	HaHost NullBool
	// hypervisor type of host: XenServer,KVM,VMware,Hyperv,BareMetal,Simulator
	Hypervisor NullString
	// the id of the host
	Id ID
	// List by keyword
	Keyword NullString
	// the name of the host
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
	// the Pod ID for the host
	PodId ID
	// list hosts by resource state. Resource state represents current state
	// determined by admin of host, valule can be one of [Enabled, Disabled,
	// Unmanaged, PrepareForMaintenance, ErrorInMaintenance, Maintenance, Error]
	ResourceState NullString
	// the state of the host
	State NullString
	// the host type
	Type NullString
	// lists hosts in the same cluster as this VM and flag hosts with enough CPU/RAm
	// to host this VM
	VirtualMachineId ID
	// the Zone ID for the host
	ZoneId ID
}

func NewListHostsParameter() (p *ListHostsParameter) {
	p = new(ListHostsParameter)
	return p
}

// Lists hosts.
func (c *Client) ListHosts(p *ListHostsParameter) ([]*Host, error) {
	obj, err := c.Request("listHosts", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Host), err
}

// DeleteHost represents the paramter of DeleteHost
type DeleteHostParameter struct {
	// Force delete the host. All HA enabled vms running on the host will be put to
	// HA; HA disabled ones will be stopped
	Forced NullBool
	// Force destroy local storage on this host. All VMs created on this local
	// storage will be destroyed
	ForceDestroyLocalStorage NullBool
	// the host ID
	Id ID
}

func NewDeleteHostParameter(id string) (p *DeleteHostParameter) {
	p = new(DeleteHostParameter)
	p.Id.Set(id)
	return p
}

// Deletes a host.
func (c *Client) DeleteHost(p *DeleteHostParameter) (*Result, error) {
	obj, err := c.Request("deleteHost", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}
