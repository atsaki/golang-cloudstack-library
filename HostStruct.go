package cloudstack

type GpuGroup struct {
	ResourceBase
	// Remaining capacity in terms of no. of more VMs that can be deployped with
	// this vGPU type
	RemainingCapacity NullNumber `json:"remainingcapacity`
	// Maximum displays per user
	MaxHeads NullNumber `json:"maxheads"`
	// Maximum X resolution per display
	MaxResolutionX NullNumber `json:"maxresolutionx"`
	// Model Name of vGPU
	VgpuType NullString `json:"vgputype"`
	// Maximum no. of vgpu per gpu card (pgpu)
	MaxVgpuPerPgpu NullNumber `json:"maxvgpuperpgpu"`
	// Maximum vgpu can be created with this vgpu type on the given gpu group
	MaxCapacity NullNumber `json:"maxcapacity"`
	// Maximum Y resolution per display
	MaxResolutionY NullNumber `json:"maxresolutiony"`
	// Video RAM for this vGPU type
	VideoRam NullNumber `json:"videoram"`
}

type Host struct {
	ResourceBase
	// the cpu average load on the host
	AverageLoad NullNumber `json:"averageload"`
	// capabilities of the host
	Capabilities NullString `json:"capabilities"`
	// the cluster ID of the host
	ClusterId ID `json:"clusterid"`
	// the cluster name of the host
	ClusterName NullString `json:"clustername"`
	// the cluster type of the cluster that host belongs to
	ClusterType NullString `json:"clustertype"`
	// the amount of the host's CPU currently allocated
	CpuAllocated NullString `json:"cpuallocated"`
	// the CPU number of the host
	CpuNumber NullNumber `json:"cpunumber"`
	// the number of CPU sockets on the host
	CpuSockets NullNumber `json:"cpusockets"`
	// the CPU speed of the host
	CpuSpeed NullNumber `json:"cpuspeed"`
	// the amount of the host's CPU currently used
	CpuUsed NullString `json:"cpuused"`
	// the amount of the host's CPU after applying the cpu.overprovisioning.factor
	CpuWithOverProvisioning NullString `json:"cpuwithoverprovisioning"`
	// the date and time the host was created
	Created NullString `json:"created"`
	// Host details in key/value pairs.
	Details map[string]string `json:"details"`
	// true if the host is disconnected. False otherwise.
	Disconnected NullString `json:"disconnected"`
	// the host's currently allocated disk size
	DiskSizeAllocated NullNumber `json:"disksizeallocated"`
	// the total disk size of the host
	DiskSizeTotal NullNumber `json:"disksizetotal"`
	// events available for the host
	Events NullString `json:"events"`
	// GPU cards present in the host
	GpuGroup []GpuGroup `json:"gpugroup"`
	// true if the host is Ha host (dedicated to vms started by HA process; false
	// otherwise
	HaHost NullBool `json:"hahost"`
	// true if this host has enough CPU and RAM capacity to migrate a VM to it,
	// false otherwise
	HasEnoughCapacity NullBool `json:"hasenoughcapacity"`
	// comma-separated list of tags for the host
	HostTags NullString `json:"hosttags"`
	// the host hypervisor
	Hypervisor NullString `json:"hypervisor"`
	// the hypervisor version
	HypervisorVersion NullString `json:"hypervisorversion"`
	// the ID of the host
	Id ID `json:"id"`
	// the IP address of the host
	IpAddress NullString `json:"ipaddress"`
	// true if local storage is active, false otherwise
	IsLocalStorageActive NullBool `json:"islocalstorageactive"`
	// the date and time the host was last pinged
	LastPinged NullString `json:"lastpinged"`
	// the management server ID of the host
	ManagementServerId ID `json:"managementserverid"`
	// the amount of the host's memory currently allocated
	MemoryAllocated NullNumber `json:"memoryallocated"`
	// the memory total of the host
	MemoryTotal NullNumber `json:"memorytotal"`
	// the amount of the host's memory currently used
	MemoryUsed NullNumber `json:"memoryused"`
	// the name of the host
	Name NullString `json:"name"`
	// the incoming network traffic on the host
	NetworkKbsRead NullNumber `json:"networkkbsread"`
	// the outgoing network traffic on the host
	NetworkKbsWrite NullNumber `json:"networkkbswrite"`
	// the OS category ID of the host
	OsCategoryId ID `json:"oscategoryid"`
	// the OS category name of the host
	OsCategoryName NullString `json:"oscategoryname"`
	// the Pod ID of the host
	PodId ID `json:"podid"`
	// the Pod name of the host
	PodName NullString `json:"podname"`
	// the date and time the host was removed
	Removed NullString `json:"removed"`
	// true if migrating a vm to this host requires storage motion, false otherwise
	RequiresStorageMotion NullBool `json:"requiresstoragemotion"`
	// the resource state of the host
	ResourceState NullString `json:"resourcestate"`
	// the state of the host
	State NullString `json:"state"`
	// true if this host is suitable(has enough capacity and satisfies all
	// conditions like hosttags, max guests vm limit etc) to migrate a VM to it ,
	// false otherwise
	SuitableForMigration NullBool `json:"suitableformigration"`
	// the host type
	Type NullString `json:"type"`
	// the host version
	Version NullString `json:"version"`
	// the Zone ID of the host
	ZoneId ID `json:"zoneid"`
	// the Zone name of the host
	ZoneName NullString `json:"zonename"`
}

type HostTag struct {
	ResourceBase
	// the host ID of the host tag
	HostId ID `json:"hostid"`
	// the ID of the host tag
	Id ID `json:"id"`
	// the name of the host tag
	Name NullString `json:"name"`
}

type DedicatedHost struct {
	ResourceBase
	// the Account ID of the host
	AccountId ID `json:"accountid"`
	// the Dedication Affinity Group ID of the host
	AffinityGroupId ID `json:"affinitygroupid"`
	// the domain ID of the host
	DomainId ID `json:"domainid"`
	// the ID of the host
	HostId ID `json:"hostid"`
	// the name of the host
	HostName NullString `json:"hostname"`
	// the ID of the dedicated resource
	Id ID `json:"id"`
}
