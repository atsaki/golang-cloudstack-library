package cloudstack

type VirtualMachine struct {
	ResourceBase
	// the account associated with the virtual machine
	Account NullString `json:"account"`
	// list of affinity groups associated with the virtual machine
	AffinityGroup []AffinityGroup `json:"affinitygroup"`
	// the number of cpu this virtual machine is running with
	CpuNumber NullNumber `json:"cpunumber"`
	// the speed of each cpu
	CpuSpeed NullNumber `json:"cpuspeed"`
	// the amount of the vm's CPU currently used
	CpuUsed NullString `json:"cpuused"`
	// the date when this virtual machine was created
	Created NullString `json:"created"`
	// Vm details in key/value pairs.
	Details NullString `json:"details"`
	// the read (io) of disk on the vm
	DiskIoRead NullNumber `json:"diskioread"`
	// the write (io) of disk on the vm
	DiskIoWrite NullNumber `json:"diskiowrite"`
	// the read (bytes) of disk on the vm
	DiskKbsRead NullNumber `json:"diskkbsread"`
	// the write (bytes) of disk on the vm
	DiskKbsWrite NullNumber `json:"diskkbswrite"`
	// the ID of the disk offering of the virtual machine
	DiskOfferingId ID `json:"diskofferingid"`
	// the name of the disk offering of the virtual machine
	DiskOfferingName NullString `json:"diskofferingname"`
	// user generated name. The name of the virtual machine is returned if no
	// displayname exists.
	DisplayName NullString `json:"displayname"`
	// an optional field whether to the display the vm to the end user or not.
	DisplayVm NullBool `json:"displayvm"`
	// the name of the domain in which the virtual machine exists
	Domain NullString `json:"domain"`
	// the ID of the domain in which the virtual machine exists
	DomainId ID `json:"domainid"`
	// the virtual network for the service offering
	ForVirtualNetwork NullBool `json:"forvirtualnetwork"`
	// the group name of the virtual machine
	Group NullString `json:"group"`
	// the group ID of the virtual machine
	GroupId ID `json:"groupid"`
	// Os type ID of the virtual machine
	GuestOsId ID `json:"guestosid"`
	// true if high-availability is enabled, false otherwise
	HaEnable NullBool `json:"haenable"`
	// the ID of the host for the virtual machine
	HostId ID `json:"hostid"`
	// the name of the host for the virtual machine
	HostName NullString `json:"hostname"`
	// the hypervisor on which the template runs
	Hypervisor NullString `json:"hypervisor"`
	// the ID of the virtual machine
	Id ID `json:"id"`
	// instance name of the user vm; this parameter is returned to the ROOT admin
	// only
	InstanceName NullString `json:"instancename"`
	// true if vm contains XS/VMWare tools inorder to support dynamic scaling of VM
	// cpu/memory.
	IsDynamicallyScalable NullBool `json:"isdynamicallyscalable"`
	// an alternate display text of the ISO attached to the virtual machine
	IsoDisplayText NullString `json:"isodisplaytext"`
	// the ID of the ISO attached to the virtual machine
	IsoId ID `json:"isoid"`
	// the name of the ISO attached to the virtual machine
	IsoName NullString `json:"isoname"`
	// ssh key-pair
	KeyPair NullString `json:"keypair"`
	// the memory allocated for the virtual machine
	Memory NullNumber `json:"memory"`
	// the name of the virtual machine
	Name NullString `json:"name"`
	// the incoming network traffic on the vm
	NetworkKbsRead NullNumber `json:"networkkbsread"`
	// the outgoing network traffic on the host
	NetworkKbsWrite NullNumber `json:"networkkbswrite"`
	// the list of nics associated with vm
	Nic []Nic `json:"nic"`
	// OS type id of the vm
	OsTypeId ID `json:"ostypeid"`
	// the password (if exists) of the virtual machine
	Password NullString `json:"password"`
	// true if the password rest feature is enabled, false otherwise
	PasswordEnabled NullBool `json:"passwordenabled"`
	// the project name of the vm
	Project NullString `json:"project"`
	// the project id of the vm
	ProjectId ID `json:"projectid"`
	// public IP address id associated with vm via Static nat rule
	PublicIp NullString `json:"publicip"`
	// public IP address id associated with vm via Static nat rule
	PublicIpId ID `json:"publicipid"`
	// device ID of the root volume
	RootDeviceId ID `json:"rootdeviceid"`
	// device type of the root volume
	RootDeviceType NullString `json:"rootdevicetype"`
	// list of security groups associated with the virtual machine
	SecurityGroup []SecurityGroup `json:"securitygroup"`
	// the ID of the service offering of the virtual machine
	ServiceOfferingId ID `json:"serviceofferingid"`
	// the name of the service offering of the virtual machine
	ServiceOfferingName NullString `json:"serviceofferingname"`
	// State of the Service from LB rule
	ServiceState NullString `json:"servicestate"`
	// the state of the virtual machine
	State NullString `json:"state"`
	// the list of resource tags associated with vm
	Tags []Tag `json:"tags"`
	//  an alternate display text of the template for the virtual machine
	TemplateDisplayText NullString `json:"templatedisplaytext"`
	// the ID of the template for the virtual machine. A -1 is returned if the
	// virtual machine was created from an ISO file.
	TemplateId ID `json:"templateid"`
	// the name of the template for the virtual machine
	TemplateName NullString `json:"templatename"`
	// the vgpu type used by the virtual machine
	Vgpu NullString `json:"vgpu"`
	// the ID of the availablility zone for the virtual machine
	ZoneId ID `json:"zoneid"`
	// the name of the availability zone for the virtual machine
	ZoneName NullString `json:"zonename"`
}
