package cloudstack


// ListDedicatedPods represents the paramter of ListDedicatedPods
type ListDedicatedPodsParameter struct {
	// the name of the account associated with the pod. Must be used with domainId.
	Account NullString
	// list dedicated pods by affinity group
	AffinityGroupId ID
	// the ID of the domain associated with the pod
	DomainId ID
	// List by keyword
	Keyword  NullString
	Page     NullNumber
	PageSize NullNumber
	// the ID of the pod
	PodId ID
}

func NewListDedicatedPodsParameter() (p *ListDedicatedPodsParameter) {
	p = new(ListDedicatedPodsParameter)
	return p
}

// Lists dedicated pods.
func (c *Client) ListDedicatedPods(p *ListDedicatedPodsParameter) ([]*DedicatedPod, error) {
	obj, err := c.Request("listDedicatedPods", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*DedicatedPod), err
}


// DeletePod represents the paramter of DeletePod
type DeletePodParameter struct {
	// the ID of the Pod
	Id ID
}

func NewDeletePodParameter(id string) (p *DeletePodParameter) {
	p = new(DeletePodParameter)
	p.Id.Set(id)
	return p
}

// Deletes a Pod.
func (c *Client) DeletePod(p *DeletePodParameter) (*Result, error) {
	obj, err := c.Request("deletePod", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// ListPods represents the paramter of ListPods
type ListPodsParameter struct {
	// list pods by allocation state
	AllocationState NullString
	// list Pods by ID
	Id ID
	// List by keyword
	Keyword NullString
	// list Pods by name
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
	// flag to display the capacity of the pods
	ShowCapacities NullBool
	// list Pods by Zone ID
	ZoneId ID
}

func NewListPodsParameter() (p *ListPodsParameter) {
	p = new(ListPodsParameter)
	return p
}

// Lists all Pods.
func (c *Client) ListPods(p *ListPodsParameter) ([]*Pod, error) {
	obj, err := c.Request("listPods", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Pod), err
}


// ReleaseDedicatedPod represents the paramter of ReleaseDedicatedPod
type ReleaseDedicatedPodParameter struct {
	// the ID of the Pod
	PodId ID
}

func NewReleaseDedicatedPodParameter(podid string) (p *ReleaseDedicatedPodParameter) {
	p = new(ReleaseDedicatedPodParameter)
	p.PodId.Set(podid)
	return p
}

// Release the dedication for the pod
func (c *Client) ReleaseDedicatedPod(p *ReleaseDedicatedPodParameter) (*Result, error) {
	obj, err := c.Request("releaseDedicatedPod", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// DedicatePod represents the paramter of DedicatePod
type DedicatePodParameter struct {
	// the name of the account which needs dedication. Must be used with domainId.
	Account NullString
	// the ID of the containing domain
	DomainId ID
	// the ID of the Pod
	PodId ID
}

func NewDedicatePodParameter(domainid string, podid string) (p *DedicatePodParameter) {
	p = new(DedicatePodParameter)
	p.DomainId.Set(domainid)
	p.PodId.Set(podid)
	return p
}

// Dedicates a Pod.
func (c *Client) DedicatePod(p *DedicatePodParameter) (*DedicatedPod, error) {
	obj, err := c.Request("dedicatePod", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*DedicatedPod), err
}


// CreatePod represents the paramter of CreatePod
type CreatePodParameter struct {
	// Allocation state of this Pod for allocation of new resources
	AllocationState NullString
	// the ending IP address for the Pod
	EndIp NullString
	// the gateway for the Pod
	Gateway NullString
	// the name of the Pod
	Name NullString
	// the netmask for the Pod
	Netmask NullString
	// the starting IP address for the Pod
	StartIp NullString
	// the Zone ID in which the Pod will be created
	ZoneId ID
}

func NewCreatePodParameter(gateway string, name string, netmask string, startip string, zoneid string) (p *CreatePodParameter) {
	p = new(CreatePodParameter)
	p.Gateway.Set(gateway)
	p.Name.Set(name)
	p.Netmask.Set(netmask)
	p.StartIp.Set(startip)
	p.ZoneId.Set(zoneid)
	return p
}

// Creates a new Pod.
func (c *Client) CreatePod(p *CreatePodParameter) (*Pod, error) {
	obj, err := c.Request("createPod", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Pod), err
}


// UpdatePod represents the paramter of UpdatePod
type UpdatePodParameter struct {
	// Allocation state of this cluster for allocation of new resources
	AllocationState NullString
	// the ending IP address for the Pod
	EndIp NullString
	// the gateway for the Pod
	Gateway NullString
	// the ID of the Pod
	Id ID
	// the name of the Pod
	Name NullString
	// the netmask of the Pod
	Netmask NullString
	// the starting IP address for the Pod
	StartIp NullString
}

func NewUpdatePodParameter(id string) (p *UpdatePodParameter) {
	p = new(UpdatePodParameter)
	p.Id.Set(id)
	return p
}

// Updates a Pod.
func (c *Client) UpdatePod(p *UpdatePodParameter) (*Pod, error) {
	obj, err := c.Request("updatePod", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Pod), err
}

