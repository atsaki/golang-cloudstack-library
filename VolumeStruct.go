package cloudstack

type Volume struct {
	ResourceBase
	// the account associated with the disk volume
	Account NullString `json:"account"`
	// the date the volume was attached to a VM instance
	Attached NullString `json:"attached"`
	// the chain info of the volume
	ChainInfo NullString `json:"chaininfo"`
	// the date the disk volume was created
	Created NullString `json:"created"`
	// the boolean state of whether the volume is destroyed or not
	Destroyed NullBool `json:"destroyed"`
	// the ID of the device on user vm the volume is attahed to. This tag is not
	// returned when the volume is detached.
	DeviceId ID `json:"deviceid"`
	// bytes read rate of the disk volume
	DiskBytesReadRate NullNumber `json:"diskbytesreadrate"`
	// bytes write rate of the disk volume
	DiskBytesWriteRate NullNumber `json:"diskbyteswriterate"`
	// io requests read rate of the disk volume
	DiskIopsReadRate NullNumber `json:"diskiopsreadrate"`
	// io requests write rate of the disk volume
	DiskIopsWriteRate NullNumber `json:"diskiopswriterate"`
	// the display text of the disk offering
	DiskOfferingDisplayText NullString `json:"diskofferingdisplaytext"`
	// ID of the disk offering
	DiskOfferingId ID `json:"diskofferingid"`
	// name of the disk offering
	DiskOfferingName NullString `json:"diskofferingname"`
	// an optional field whether to the display the volume to the end user or not.
	DisplayVolume NullBool `json:"displayvolume"`
	// the domain associated with the disk volume
	Domain NullString `json:"domain"`
	// the ID of the domain associated with the disk volume
	DomainId ID `json:"domainid"`
	// Hypervisor the volume belongs to
	Hypervisor NullString `json:"hypervisor"`
	// ID of the disk volume
	Id ID `json:"id"`
	// true if the volume is extractable, false otherwise
	IsExtractable NullBool `json:"isextractable"`
	// an alternate display text of the ISO attached to the virtual machine
	IsoDisplayText NullString `json:"isodisplaytext"`
	// the ID of the ISO attached to the virtual machine
	IsoId ID `json:"isoid"`
	// the name of the ISO attached to the virtual machine
	IsoName NullString `json:"isoname"`
	// max iops of the disk volume
	MaxIops NullNumber `json:"maxiops"`
	// min iops of the disk volume
	MinIops NullNumber `json:"miniops"`
	// name of the disk volume
	Name NullString `json:"name"`
	// the path of the volume
	Path NullString `json:"path"`
	// the project name of the vpn
	Project NullString `json:"project"`
	// the project id of the vpn
	ProjectId ID `json:"projectid"`
	// need quiesce vm or not when taking snapshot
	QuiesceVm NullBool `json:"quiescevm"`
	// the display text of the service offering for root disk
	ServiceOfferingDisplayText NullString `json:"serviceofferingdisplaytext"`
	// ID of the service offering for root disk
	ServiceOfferingId ID `json:"serviceofferingid"`
	// name of the service offering for root disk
	ServiceOfferingName NullString `json:"serviceofferingname"`
	// size of the disk volume
	Size NullNumber `json:"size"`
	// ID of the snapshot from which this volume was created
	SnapshotId ID `json:"snapshotid"`
	// the state of the disk volume
	State NullString `json:"state"`
	// the status of the volume
	Status NullString `json:"status"`
	// name of the primary storage hosting the disk volume
	Storage NullString `json:"storage"`
	// id of the primary storage hosting the disk volume; returned to admin user
	// only
	StorageId ID `json:"storageid"`
	// shared or local storage
	StorageType NullString `json:"storagetype"`
	// the list of resource tags associated with volume
	Tags []Tag `json:"tags"`
	//  an alternate display text of the template for the virtual machine
	TemplateDisplayText NullString `json:"templatedisplaytext"`
	// the ID of the template for the virtual machine. A -1 is returned if the
	// virtual machine was created from an ISO file.
	TemplateId ID `json:"templateid"`
	// the name of the template for the virtual machine
	TemplateName NullString `json:"templatename"`
	// type of the disk volume (ROOT or DATADISK)
	Type NullString `json:"type"`
	// id of the virtual machine
	VirtualMachineId ID `json:"virtualmachineid"`
	// display name of the virtual machine
	VmDisplayName NullString `json:"vmdisplayname"`
	// name of the virtual machine
	VmName NullString `json:"vmname"`
	// state of the virtual machine
	VmState NullString `json:"vmstate"`
	// ID of the availability zone
	ZoneId ID `json:"zoneid"`
	// name of the availability zone
	ZoneName NullString `json:"zonename"`
}
