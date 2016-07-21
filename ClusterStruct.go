package cloudstack

type Cluster struct {
	ResourceBase
	// the allocation state of the cluster
	AllocationState NullString `json:"allocationstate"`
	// the capacity of the Cluster
	Capacity []Capacity `json:"capacity"`
	// the type of the cluster
	ClusterType NullString `json:"clustertype"`
	// The cpu overcommit ratio of the cluster
	CpuOverCommitRatio NullString `json:"cpuovercommitratio"`
	// the hypervisor type of the cluster
	HypervisorType NullString `json:"hypervisortype"`
	// the cluster ID
	Id ID `json:"id"`
	// whether this cluster is managed by cloudstack
	ManagedState NullString `json:"managedstate"`
	// The memory overcommit ratio of the cluster
	MemoryOverCommitRatio NullString `json:"memoryovercommitratio"`
	// the cluster name
	Name NullString `json:"name"`
	// the Pod ID of the cluster
	PodId ID `json:"podid"`
	// the Pod name of the cluster
	PodName NullString `json:"podname"`
	// the Zone ID of the cluster
	ZoneId ID `json:"zoneid"`
	// the Zone name of the cluster
	ZoneName NullString `json:"zonename"`
}

type DedicatedCluster struct {
	ResourceBase
	// the Account ID of the cluster
	AccountId ID `json:"accountid"`
	// the Dedication Affinity Group ID of the cluster
	AffinityGroupId ID `json:"affinitygroupid"`
	// the ID of the cluster
	ClusterId ID `json:"clusterid"`
	// the name of the cluster
	ClusterName NullString `json:"clustername"`
	// the domain ID of the cluster
	DomainId ID `json:"domainid"`
	// the ID of the dedicated resource
	Id ID `json:"id"`
}
