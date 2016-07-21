package cloudstack


// AddCluster represents the paramter of AddCluster
type AddClusterParameter struct {
	// Allocation state of this cluster for allocation of new resources
	AllocationState NullString
	// the cluster name
	ClusterName NullString
	// type of the cluster: CloudManaged, ExternalManaged
	ClusterType NullString
	// Name of virtual switch used for guest traffic in the cluster. This would
	// override zone wide traffic label setting.
	GuestVswitchName NullString
	// Type of virtual switch used for guest traffic in the cluster. Allowed values
	// are, vmwaresvs (for VMware standard vSwitch) and vmwaredvs (for VMware
	// distributed vSwitch)
	GuestVswitchType NullString
	// hypervisor type of the cluster:
	// XenServer,KVM,VMware,Hyperv,BareMetal,Simulator
	Hypervisor NullString
	// the password for the host
	Password NullString
	// the Pod ID for the host
	PodId ID
	// Name of virtual switch used for public traffic in the cluster.  This would
	// override zone wide traffic label setting.
	PublicVswitchName NullString
	// Type of virtual switch used for public traffic in the cluster. Allowed values
	// are, vmwaresvs (for VMware standard vSwitch) and vmwaredvs (for VMware
	// distributed vSwitch)
	PublicVswitchType NullString
	// the URL
	Url NullString
	// the username for the cluster
	UserName NullString
	// the ipaddress of the VSM associated with this cluster
	VsmIpAddress NullString
	// the password for the VSM associated with this cluster
	VsmPassword NullString
	// the username for the VSM associated with this cluster
	VsmUserName NullString
	// the Zone ID for the cluster
	ZoneId ID
}

func NewAddClusterParameter(clustername string, clustertype string, hypervisor string, podid string, zoneid string) (p *AddClusterParameter) {
	p = new(AddClusterParameter)
	p.ClusterName.Set(clustername)
	p.ClusterType.Set(clustertype)
	p.Hypervisor.Set(hypervisor)
	p.PodId.Set(podid)
	p.ZoneId.Set(zoneid)
	return p
}

// Adds a new cluster
func (c *Client) AddCluster(p *AddClusterParameter) (*Cluster, error) {
	obj, err := c.Request("addCluster", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Cluster), err
}


// ListClusters represents the paramter of ListClusters
type ListClustersParameter struct {
	// lists clusters by allocation state
	AllocationState NullString
	// lists clusters by cluster type
	ClusterType NullString
	// lists clusters by hypervisor type
	Hypervisor NullString
	// lists clusters by the cluster ID
	Id ID
	// List by keyword
	Keyword NullString
	// whether this cluster is managed by cloudstack
	ManagedState NullString
	// lists clusters by the cluster name
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
	// lists clusters by Pod ID
	PodId ID
	// flag to display the capacity of the clusters
	ShowCapacities NullBool
	// lists clusters by Zone ID
	ZoneId ID
}

func NewListClustersParameter() (p *ListClustersParameter) {
	p = new(ListClustersParameter)
	return p
}

// Lists clusters.
func (c *Client) ListClusters(p *ListClustersParameter) ([]*Cluster, error) {
	obj, err := c.Request("listClusters", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Cluster), err
}


// ReleaseDedicatedCluster represents the paramter of ReleaseDedicatedCluster
type ReleaseDedicatedClusterParameter struct {
	// the ID of the Cluster
	ClusterId ID
}

func NewReleaseDedicatedClusterParameter(clusterid string) (p *ReleaseDedicatedClusterParameter) {
	p = new(ReleaseDedicatedClusterParameter)
	p.ClusterId.Set(clusterid)
	return p
}

// Release the dedication for cluster
func (c *Client) ReleaseDedicatedCluster(p *ReleaseDedicatedClusterParameter) (*Result, error) {
	obj, err := c.Request("releaseDedicatedCluster", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// ListDedicatedClusters represents the paramter of ListDedicatedClusters
type ListDedicatedClustersParameter struct {
	// the name of the account associated with the cluster. Must be used with
	// domainId.
	Account NullString
	// list dedicated clusters by affinity group
	AffinityGroupId ID
	// the ID of the cluster
	ClusterId ID
	// the ID of the domain associated with the cluster
	DomainId ID
	// List by keyword
	Keyword  NullString
	Page     NullNumber
	PageSize NullNumber
}

func NewListDedicatedClustersParameter() (p *ListDedicatedClustersParameter) {
	p = new(ListDedicatedClustersParameter)
	return p
}

// Lists dedicated clusters.
func (c *Client) ListDedicatedClusters(p *ListDedicatedClustersParameter) ([]*DedicatedCluster, error) {
	obj, err := c.Request("listDedicatedClusters", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*DedicatedCluster), err
}


// DedicateCluster represents the paramter of DedicateCluster
type DedicateClusterParameter struct {
	// the name of the account which needs dedication. Must be used with domainId.
	Account NullString
	// the ID of the Cluster
	ClusterId ID
	// the ID of the containing domain
	DomainId ID
}

func NewDedicateClusterParameter(clusterid string, domainid string) (p *DedicateClusterParameter) {
	p = new(DedicateClusterParameter)
	p.ClusterId.Set(clusterid)
	p.DomainId.Set(domainid)
	return p
}

// Dedicate an existing cluster
func (c *Client) DedicateCluster(p *DedicateClusterParameter) (*DedicatedCluster, error) {
	obj, err := c.Request("dedicateCluster", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*DedicatedCluster), err
}


// UpdateCluster represents the paramter of UpdateCluster
type UpdateClusterParameter struct {
	// Allocation state of this cluster for allocation of new resources
	AllocationState NullString
	// the cluster name
	ClusterName NullString
	// hypervisor type of the cluster
	ClusterType NullString
	// hypervisor type of the cluster
	Hypervisor NullString
	// the ID of the Cluster
	Id ID
	// whether this cluster is managed by cloudstack
	ManagedState NullString
}

func NewUpdateClusterParameter(id string) (p *UpdateClusterParameter) {
	p = new(UpdateClusterParameter)
	p.Id.Set(id)
	return p
}

// Updates an existing cluster
func (c *Client) UpdateCluster(p *UpdateClusterParameter) (*Cluster, error) {
	obj, err := c.Request("updateCluster", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Cluster), err
}


// DeleteCluster represents the paramter of DeleteCluster
type DeleteClusterParameter struct {
	// the cluster ID
	Id ID
}

func NewDeleteClusterParameter(id string) (p *DeleteClusterParameter) {
	p = new(DeleteClusterParameter)
	p.Id.Set(id)
	return p
}

// Deletes a cluster.
func (c *Client) DeleteCluster(p *DeleteClusterParameter) (*Result, error) {
	obj, err := c.Request("deleteCluster", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}

