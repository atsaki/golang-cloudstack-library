package cloudstack

type StoragePool struct {
	ResourceBase
	// IOPS CloudStack can provision from this storage pool
	CapacityIops NullNumber `json:"capacityiops"`
	// the ID of the cluster for the storage pool
	ClusterId ID `json:"clusterid"`
	// the name of the cluster for the storage pool
	ClusterName NullString `json:"clustername"`
	// the date and time the storage pool was created
	Created NullString `json:"created"`
	// the host's currently allocated disk size
	DiskSizeAllocated NullNumber `json:"disksizeallocated"`
	// the total disk size of the storage pool
	DiskSizeTotal NullNumber `json:"disksizetotal"`
	// the host's currently used disk size
	DiskSizeUsed NullNumber `json:"disksizeused"`
	// the hypervisor type of the storage pool
	Hypervisor NullString `json:"hypervisor"`
	// the ID of the storage pool
	Id ID `json:"id"`
	// the IP address of the storage pool
	IpAddress NullString `json:"ipaddress"`
	// the name of the storage pool
	Name NullString `json:"name"`
	// the overprovisionfactor for the storage pool
	OverProvisionFactor NullString `json:"overprovisionfactor"`
	// the storage pool path
	Path NullString `json:"path"`
	// the Pod ID of the storage pool
	PodId ID `json:"podid"`
	// the Pod name of the storage pool
	PodName NullString `json:"podname"`
	// the scope of the storage pool
	Scope NullString `json:"scope"`
	// the state of the storage pool
	State NullString `json:"state"`
	// the storage pool capabilities
	StorageCapabilities map[string]string `json:"storagecapabilities"`
	// true if this pool is suitable to migrate a volume, false otherwise
	SuitableForMigration NullBool `json:"suitableformigration"`
	// the tags for the storage pool
	Tags NullString `json:"tags"`
	// the storage pool type
	Type NullString `json:"type"`
	// the Zone ID of the storage pool
	ZoneId ID `json:"zoneid"`
	// the Zone name of the storage pool
	ZoneName NullString `json:"zonename"`
}
