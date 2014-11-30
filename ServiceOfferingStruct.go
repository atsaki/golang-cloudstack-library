package cloudstack

type ServiceOffering struct {
	ResourceBase
	// the number of CPU
	CpuNumber NullNumber `json:"cpunumber"`
	// the clock rate CPU speed in Mhz
	CpuSpeed NullNumber `json:"cpuspeed"`
	// the date this service offering was created
	Created NullString `json:"created"`
	// is this a  default system vm offering
	DefaultUse NullBool `json:"defaultuse"`
	// deployment strategy used to deploy VM.
	DeploymentPlanner NullString `json:"deploymentplanner"`
	// bytes read rate of the service offering
	DiskBytesReadRate NullNumber `json:"diskbytesreadrate"`
	// bytes write rate of the service offering
	DiskBytesWriteRate NullNumber `json:"diskbyteswriterate"`
	// io requests read rate of the service offering
	DiskIopsReadRate NullNumber `json:"diskiopsreadrate"`
	// io requests write rate of the service offering
	DiskIopsWriteRate NullNumber `json:"diskiopswriterate"`
	// an alternate display text of the service offering.
	DisplayText NullString `json:"displaytext"`
	// Domain name for the offering
	Domain NullString `json:"domain"`
	// the domain id of the service offering
	DomainId ID `json:"domainid"`
	// the host tag for the service offering
	HostTags NullString `json:"hosttags"`
	// Hypervisor snapshot reserve space as a percent of a volume (for managed
	// storage using Xen or VMware)
	HypervisorSnapshotReserve NullNumber `json:"hypervisorsnapshotreserve"`
	// the id of the service offering
	Id ID `json:"id"`
	// is true if the offering is customized
	IsCustomized NullBool `json:"iscustomized"`
	// true if disk offering uses custom iops, false otherwise
	IsCustomizedIops NullBool `json:"iscustomizediops"`
	// is this a system vm offering
	IsSystem NullBool `json:"issystem"`
	// true if the vm needs to be volatile, i.e., on every reboot of vm from API
	// root disk is discarded and creates a new root disk
	IsVolatile NullBool `json:"isvolatile"`
	// restrict the CPU usage to committed service offering
	LimitCpuUse NullBool `json:"limitcpuuse"`
	// the max iops of the disk offering
	MaxIops NullNumber `json:"maxiops"`
	// the memory in MB
	Memory NullNumber `json:"memory"`
	// the min iops of the disk offering
	MinIops NullNumber `json:"miniops"`
	// the name of the service offering
	Name NullString `json:"name"`
	// data transfer rate in megabits per second allowed.
	NetworkRate NullNumber `json:"networkrate"`
	// the ha support in the service offering
	OfferHa NullBool `json:"offerha"`
	// additional key/value details tied with this service offering
	ServiceOfferingDetails NullString `json:"serviceofferingdetails"`
	// the storage type for this service offering
	StorageType NullString `json:"storagetype"`
	// is this a the systemvm type for system vm offering
	SystemVmType NullString `json:"systemvmtype"`
	// the tags for the service offering
	Tags NullString `json:"tags"`
}
