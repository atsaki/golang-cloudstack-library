package cloudstack

type DiskOffering struct {
	ResourceBase
	// the cache mode to use for this disk offering. none, writeback or writethrough
	CacheMode NullString `json:"cachemode"`
	// the date this disk offering was created
	Created NullString `json:"created"`
	// bytes read rate of the disk offering
	DiskBytesReadRate NullNumber `json:"diskbytesreadrate"`
	// bytes write rate of the disk offering
	DiskBytesWriteRate NullNumber `json:"diskbyteswriterate"`
	// io requests read rate of the disk offering
	DiskIopsReadRate NullNumber `json:"diskiopsreadrate"`
	// io requests write rate of the disk offering
	DiskIopsWriteRate NullNumber `json:"diskiopswriterate"`
	// the size of the disk offering in GB
	DiskSize NullNumber `json:"disksize"`
	// whether to display the offering to the end user or not.
	DisplayOffering NullBool `json:"displayoffering"`
	// an alternate display text of the disk offering.
	DisplayText NullString `json:"displaytext"`
	// the domain name this disk offering belongs to. Ignore this information as it
	// is not currently applicable.
	Domain NullString `json:"domain"`
	// the domain ID this disk offering belongs to. Ignore this information as it is
	// not currently applicable.
	DomainId ID `json:"domainid"`
	// Hypervisor snapshot reserve space as a percent of a volume (for managed
	// storage using Xen or VMware)
	HypervisorSnapshotReserve NullNumber `json:"hypervisorsnapshotreserve"`
	// unique ID of the disk offering
	Id ID `json:"id"`
	// true if disk offering uses custom size, false otherwise
	IsCustomized NullBool `json:"iscustomized"`
	// true if disk offering uses custom iops, false otherwise
	IsCustomizedIops NullBool `json:"iscustomizediops"`
	// the max iops of the disk offering
	MaxIops NullNumber `json:"maxiops"`
	// the min iops of the disk offering
	MinIops NullNumber `json:"miniops"`
	// the name of the disk offering
	Name NullString `json:"name"`
	// the storage type for this disk offering
	StorageType NullString `json:"storagetype"`
	// the tags for the disk offering
	Tags NullString `json:"tags"`
}
