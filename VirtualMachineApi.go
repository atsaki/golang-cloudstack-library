package cloudstack


// AddNicToVirtualMachine represents the paramter of AddNicToVirtualMachine
type AddNicToVirtualMachineParameter struct {
	// IP Address for the new network
	IpAddress NullString
	// Network ID
	NetworkId ID
	// Virtual Machine ID
	VirtualMachineId ID
}

func NewAddNicToVirtualMachineParameter(networkid string, virtualmachineid string) (p *AddNicToVirtualMachineParameter) {
	p = new(AddNicToVirtualMachineParameter)
	p.NetworkId.Set(networkid)
	p.VirtualMachineId.Set(virtualmachineid)
	return p
}

// Adds VM to specified network by creating a NIC
func (c *Client) AddNicToVirtualMachine(p *AddNicToVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("addNicToVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// GetVMPassword represents the paramter of GetVMPassword
type GetVMPasswordParameter struct {
	// The ID of the virtual machine
	Id ID
}

func NewGetVMPasswordParameter(id string) (p *GetVMPasswordParameter) {
	p = new(GetVMPasswordParameter)
	p.Id.Set(id)
	return p
}

// Returns an encrypted password for the VM
func (c *Client) GetVMPassword(p *GetVMPasswordParameter) (*VMPassword, error) {
	obj, err := c.Request("getVMPassword", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VMPassword), err
}


// CleanVMReservations represents the paramter of CleanVMReservations
type CleanVMReservationsParameter struct {
}

func NewCleanVMReservationsParameter() (p *CleanVMReservationsParameter) {
	p = new(CleanVMReservationsParameter)
	return p
}

// Cleanups VM reservations in the database.
func (c *Client) CleanVMReservations(p *CleanVMReservationsParameter) (*Result, error) {
	obj, err := c.Request("cleanVMReservations", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// RecoverVirtualMachine represents the paramter of RecoverVirtualMachine
type RecoverVirtualMachineParameter struct {
	// The ID of the virtual machine
	Id ID
}

func NewRecoverVirtualMachineParameter(id string) (p *RecoverVirtualMachineParameter) {
	p = new(RecoverVirtualMachineParameter)
	p.Id.Set(id)
	return p
}

// Recovers a virtual machine.
func (c *Client) RecoverVirtualMachine(p *RecoverVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("recoverVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// DeployVirtualMachine represents the paramter of DeployVirtualMachine
type DeployVirtualMachineParameter struct {
	// an optional account for the virtual machine. Must be used with domainId.
	Account NullString
	// comma separated list of affinity groups id that are going to be applied to
	// the virtual machine. Mutually exclusive with affinitygroupnames parameter
	AffinityGroupIds []string
	// comma separated list of affinity groups names that are going to be applied to
	// the virtual machine.Mutually exclusive with affinitygroupids parameter
	AffinityGroupNames []string
	// an optional field, in case you want to set a custom id to the resource.
	// Allowed to Root Admins only
	CustomId ID
	// Deployment planner to use for vm allocation. Available to ROOT admin only
	DeploymentPlanner NullString
	// used to specify the custom parameters.
	Details map[string]string
	// the ID of the disk offering for the virtual machine. If the template is of
	// ISO format, the diskOfferingId is for the root disk volume. Otherwise this
	// parameter is used to indicate the offering for the data disk volume. If the
	// templateId parameter passed is from a Template object, the diskOfferingId
	// refers to a DATA Disk Volume created. If the templateId parameter passed is
	// from an ISO object, the diskOfferingId refers to a ROOT Disk Volume created.
	DiskOfferingId ID
	// an optional user generated name for the virtual machine
	DisplayName NullString
	// an optional field, whether to the display the vm to the end user or not.
	DisplayVm NullBool
	// an optional domainId for the virtual machine. If the account parameter is
	// used, domainId must also be used.
	DomainId ID
	// an optional group for the virtual machine
	Group NullString
	// destination Host ID to deploy the VM to - parameter available for root admin
	// only
	HostId ID
	// the hypervisor on which to deploy the virtual machine
	Hypervisor NullString
	// the ipv6 address for default vm's network
	Ip6Address NullString
	// the ip address for default vm's network
	IpAddress NullString
	// ip to network mapping. Can't be specified with networkIds parameter. Example:
	// iptonetworklist[0].ip=10.10.10.11&iptonetworklist[0].ipv6=fc00:1234:5678::abcd&iptonetworklist[0].networkid=uuid
	// - requests to use ip 10.10.10.11 in network id=uuid
	IpToNetworkList map[string]string
	// an optional keyboard device type for the virtual machine. valid value can be
	// one of de,de-ch,es,fi,fr,fr-be,fr-ch,is,it,jp,nl-be,no,pt,uk,us
	Keyboard NullString
	// name of the ssh key pair used to login to the virtual machine
	KeyPair NullString
	// host name for the virtual machine
	Name NullString
	// list of network ids used by virtual machine. Can't be specified with
	// ipToNetworkList parameter
	NetworkIds []string
	// Deploy vm for the project
	ProjectId ID
	// Optional field to resize root disk on deploy. Only applies to template-based
	// deployments. Analogous to details[0].rootdisksize, which takes precedence
	// over this parameter if both are provided
	RootDiskSize NullNumber
	// comma separated list of security groups id that going to be applied to the
	// virtual machine. Should be passed only when vm is created from a zone with
	// Basic Network support. Mutually exclusive with securitygroupnames parameter
	SecurityGroupIds []string
	// comma separated list of security groups names that going to be applied to the
	// virtual machine. Should be passed only when vm is created from a zone with
	// Basic Network support. Mutually exclusive with securitygroupids parameter
	SecurityGroupNames []string
	// the ID of the service offering for the virtual machine
	ServiceOfferingId ID
	// the arbitrary size for the DATADISK volume. Mutually exclusive with
	// diskOfferingId
	Size NullNumber
	// true if network offering supports specifying ip ranges; defaulted to true if
	// not specified
	StartVm NullBool
	// the ID of the template for the virtual machine
	TemplateId ID
	// an optional binary data that can be sent to the virtual machine upon a
	// successful deployment. This binary data must be base64 encoded before adding
	// it to the request. Using HTTP GET (via querystring), you can send up to 2KB
	// of data after base64 encoding. Using HTTP POST(via POST body), you can send
	// up to 32K of data after base64 encoding.
	UserData NullString
	// availability zone for the virtual machine
	ZoneId ID
}

func NewDeployVirtualMachineParameter(serviceofferingid string, templateid string, zoneid string) (p *DeployVirtualMachineParameter) {
	p = new(DeployVirtualMachineParameter)
	p.ServiceOfferingId.Set(serviceofferingid)
	p.TemplateId.Set(templateid)
	p.ZoneId.Set(zoneid)
	return p
}

// Creates and automatically starts a virtual machine based on a service
// offering, disk offering, and template.
func (c *Client) DeployVirtualMachine(p *DeployVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("deployVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// ExpungeVirtualMachine represents the paramter of ExpungeVirtualMachine
type ExpungeVirtualMachineParameter struct {
	// The ID of the virtual machine
	Id ID
}

func NewExpungeVirtualMachineParameter(id string) (p *ExpungeVirtualMachineParameter) {
	p = new(ExpungeVirtualMachineParameter)
	p.Id.Set(id)
	return p
}

// Expunge a virtual machine. Once expunged, it cannot be recoverd.
func (c *Client) ExpungeVirtualMachine(p *ExpungeVirtualMachineParameter) (*Result, error) {
	obj, err := c.Request("expungeVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// ScaleVirtualMachine represents the paramter of ScaleVirtualMachine
type ScaleVirtualMachineParameter struct {
	// name value pairs of custom parameters for cpu,memory and cpunumber. example
	// details[i].name=value
	Details map[string]string
	// The ID of the virtual machine
	Id ID
	// the ID of the service offering for the virtual machine
	ServiceOfferingId ID
}

func NewScaleVirtualMachineParameter(id string, serviceofferingid string) (p *ScaleVirtualMachineParameter) {
	p = new(ScaleVirtualMachineParameter)
	p.Id.Set(id)
	p.ServiceOfferingId.Set(serviceofferingid)
	return p
}

// Scales the virtual machine to a new service offering.
func (c *Client) ScaleVirtualMachine(p *ScaleVirtualMachineParameter) (*Result, error) {
	obj, err := c.Request("scaleVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// RebootVirtualMachine represents the paramter of RebootVirtualMachine
type RebootVirtualMachineParameter struct {
	// The ID of the virtual machine
	Id ID
}

func NewRebootVirtualMachineParameter(id string) (p *RebootVirtualMachineParameter) {
	p = new(RebootVirtualMachineParameter)
	p.Id.Set(id)
	return p
}

// Reboots a virtual machine.
func (c *Client) RebootVirtualMachine(p *RebootVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("rebootVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// RemoveNicFromVirtualMachine represents the paramter of RemoveNicFromVirtualMachine
type RemoveNicFromVirtualMachineParameter struct {
	// NIC ID
	NicId ID
	// Virtual Machine ID
	VirtualMachineId ID
}

func NewRemoveNicFromVirtualMachineParameter(nicid string, virtualmachineid string) (p *RemoveNicFromVirtualMachineParameter) {
	p = new(RemoveNicFromVirtualMachineParameter)
	p.NicId.Set(nicid)
	p.VirtualMachineId.Set(virtualmachineid)
	return p
}

// Removes VM from specified network by deleting a NIC
func (c *Client) RemoveNicFromVirtualMachine(p *RemoveNicFromVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("removeNicFromVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// UpdateDefaultNicForVirtualMachine represents the paramter of UpdateDefaultNicForVirtualMachine
type UpdateDefaultNicForVirtualMachineParameter struct {
	// NIC ID
	NicId ID
	// Virtual Machine ID
	VirtualMachineId ID
}

func NewUpdateDefaultNicForVirtualMachineParameter(nicid string, virtualmachineid string) (p *UpdateDefaultNicForVirtualMachineParameter) {
	p = new(UpdateDefaultNicForVirtualMachineParameter)
	p.NicId.Set(nicid)
	p.VirtualMachineId.Set(virtualmachineid)
	return p
}

// Changes the default NIC on a VM
func (c *Client) UpdateDefaultNicForVirtualMachine(p *UpdateDefaultNicForVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("updateDefaultNicForVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// ResetPasswordForVirtualMachine represents the paramter of ResetPasswordForVirtualMachine
type ResetPasswordForVirtualMachineParameter struct {
	// The ID of the virtual machine
	Id ID
}

func NewResetPasswordForVirtualMachineParameter(id string) (p *ResetPasswordForVirtualMachineParameter) {
	p = new(ResetPasswordForVirtualMachineParameter)
	p.Id.Set(id)
	return p
}

// Resets the password for virtual machine. The virtual machine must be in a
// "Stopped" state and the template must already support this feature for this
// command to take effect. [async]
func (c *Client) ResetPasswordForVirtualMachine(p *ResetPasswordForVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("resetPasswordForVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// ChangeServiceForVirtualMachine represents the paramter of ChangeServiceForVirtualMachine
type ChangeServiceForVirtualMachineParameter struct {
	// name value pairs of custom parameters for cpu, memory and cpunumber. example
	// details[i].name=value
	Details map[string]string
	// The ID of the virtual machine
	Id ID
	// the service offering ID to apply to the virtual machine
	ServiceOfferingId ID
}

func NewChangeServiceForVirtualMachineParameter(id string, serviceofferingid string) (p *ChangeServiceForVirtualMachineParameter) {
	p = new(ChangeServiceForVirtualMachineParameter)
	p.Id.Set(id)
	p.ServiceOfferingId.Set(serviceofferingid)
	return p
}

// Changes the service offering for a virtual machine. The virtual machine must
// be in a "Stopped" state for this command to take effect.
func (c *Client) ChangeServiceForVirtualMachine(p *ChangeServiceForVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("changeServiceForVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// StartVirtualMachine represents the paramter of StartVirtualMachine
type StartVirtualMachineParameter struct {
	// Deployment planner to use for vm allocation. Available to ROOT admin only
	DeploymentPlanner NullString
	// destination Host ID to deploy the VM to - parameter available for root admin
	// only
	HostId ID
	// The ID of the virtual machine
	Id ID
}

func NewStartVirtualMachineParameter(id string) (p *StartVirtualMachineParameter) {
	p = new(StartVirtualMachineParameter)
	p.Id.Set(id)
	return p
}

// Starts a virtual machine.
func (c *Client) StartVirtualMachine(p *StartVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("startVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// MigrateVirtualMachine represents the paramter of MigrateVirtualMachine
type MigrateVirtualMachineParameter struct {
	// Destination Host ID to migrate VM to. Required for live migrating a VM from
	// host to host
	HostId ID
	// Destination storage pool ID to migrate VM volumes to. Required for migrating
	// the root disk volume
	StorageId ID
	// the ID of the virtual machine
	VirtualMachineId ID
}

func NewMigrateVirtualMachineParameter(virtualmachineid string) (p *MigrateVirtualMachineParameter) {
	p = new(MigrateVirtualMachineParameter)
	p.VirtualMachineId.Set(virtualmachineid)
	return p
}

// Attempts Migration of a VM to a different host or Root volume of the vm to a
// different storage pool
func (c *Client) MigrateVirtualMachine(p *MigrateVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("migrateVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// AssignVirtualMachine represents the paramter of AssignVirtualMachine
type AssignVirtualMachineParameter struct {
	// account name of the new VM owner.
	Account NullString
	// domain id of the new VM owner.
	DomainId ID
	// list of new network ids in which the moved VM will participate. In case no
	// network ids are provided the VM will be part of the default network for that
	// zone. In case there is no network yet created for the new account the default
	// network will be created.
	NetworkIds []string
	// list of security group ids to be applied on the virtual machine. In case no
	// security groups are provided the VM is part of the default security group.
	SecurityGroupIds []string
	// id of the VM to be moved
	VirtualMachineId ID
}

func NewAssignVirtualMachineParameter(account string, domainid string, virtualmachineid string) (p *AssignVirtualMachineParameter) {
	p = new(AssignVirtualMachineParameter)
	p.Account.Set(account)
	p.DomainId.Set(domainid)
	p.VirtualMachineId.Set(virtualmachineid)
	return p
}

// Change ownership of a VM from one account to another. This API is available
// for Basic zones with security groups and Advanced zones with guest networks.
// A root administrator can reassign a VM from any account to any other account
// in any domain. A domain administrator can reassign a VM to any account in the
// same domain.
func (c *Client) AssignVirtualMachine(p *AssignVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("assignVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// UpdateVirtualMachine represents the paramter of UpdateVirtualMachine
type UpdateVirtualMachineParameter struct {
	// an optional field, in case you want to set a custom id to the resource.
	// Allowed to Root Admins only
	CustomId ID
	// user generated name
	DisplayName NullString
	// an optional field, whether to the display the vm to the end user or not.
	DisplayVm NullBool
	// group of the virtual machine
	Group NullString
	// true if high-availability is enabled for the virtual machine, false otherwise
	HaEnable NullBool
	// The ID of the virtual machine
	Id ID
	// true if VM contains XS/VMWare tools inorder to support dynamic scaling of VM
	// cpu/memory
	IsDynamicallyScalable NullBool
	// new host name of the vm. The VM has to be stopped/started for this update to
	// take affect
	Name NullString
	// the ID of the OS type that best represents this VM.
	OsTypeId ID
	// an optional binary data that can be sent to the virtual machine upon a
	// successful deployment. This binary data must be base64 encoded before adding
	// it to the request. Using HTTP GET (via querystring), you can send up to 2KB
	// of data after base64 encoding. Using HTTP POST(via POST body), you can send
	// up to 32K of data after base64 encoding.
	UserData NullString
}

func NewUpdateVirtualMachineParameter(id string) (p *UpdateVirtualMachineParameter) {
	p = new(UpdateVirtualMachineParameter)
	p.Id.Set(id)
	return p
}

// Updates properties of a virtual machine. The VM has to be stopped and
// restarted for the new properties to take effect. UpdateVirtualMachine does
// not first check whether the VM is stopped. Therefore, stop the VM manually
// before issuing this call.
func (c *Client) UpdateVirtualMachine(p *UpdateVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("updateVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// RestoreVirtualMachine represents the paramter of RestoreVirtualMachine
type RestoreVirtualMachineParameter struct {
	// an optional template Id to restore vm from the new template. This can be an
	// ISO id in case of restore vm deployed using ISO
	TemplateId ID
	// Virtual Machine ID
	VirtualMachineId ID
}

func NewRestoreVirtualMachineParameter(virtualmachineid string) (p *RestoreVirtualMachineParameter) {
	p = new(RestoreVirtualMachineParameter)
	p.VirtualMachineId.Set(virtualmachineid)
	return p
}

// Restore a VM to original template/ISO or new template/ISO
func (c *Client) RestoreVirtualMachine(p *RestoreVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("restoreVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// MigrateVirtualMachineWithVolume represents the paramter of MigrateVirtualMachineWithVolume
type MigrateVirtualMachineWithVolumeParameter struct {
	// Destination Host ID to migrate VM to.
	HostId ID
	// Map of pool to which each volume should be migrated (volume/pool pair)
	MigrateTo map[string]string
	// the ID of the virtual machine
	VirtualMachineId ID
}

func NewMigrateVirtualMachineWithVolumeParameter(hostid string, virtualmachineid string) (p *MigrateVirtualMachineWithVolumeParameter) {
	p = new(MigrateVirtualMachineWithVolumeParameter)
	p.HostId.Set(hostid)
	p.VirtualMachineId.Set(virtualmachineid)
	return p
}

// Attempts Migration of a VM with its volumes to a different host
func (c *Client) MigrateVirtualMachineWithVolume(p *MigrateVirtualMachineWithVolumeParameter) (*VirtualMachine, error) {
	obj, err := c.Request("migrateVirtualMachineWithVolume", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// StopVirtualMachine represents the paramter of StopVirtualMachine
type StopVirtualMachineParameter struct {
	// Force stop the VM (vm is marked as Stopped even when command fails to be send
	// to the backend).  The caller knows the VM is stopped.
	Forced NullBool
	// The ID of the virtual machine
	Id ID
}

func NewStopVirtualMachineParameter(id string) (p *StopVirtualMachineParameter) {
	p = new(StopVirtualMachineParameter)
	p.Id.Set(id)
	return p
}

// Stops a virtual machine.
func (c *Client) StopVirtualMachine(p *StopVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("stopVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// DestroyVirtualMachine represents the paramter of DestroyVirtualMachine
type DestroyVirtualMachineParameter struct {
	// If true is passed, the vm is expunged immediately. False by default.
	// Parameter can be passed to the call by ROOT/Domain admin only
	Expunge NullBool
	// The ID of the virtual machine
	Id ID
}

func NewDestroyVirtualMachineParameter(id string) (p *DestroyVirtualMachineParameter) {
	p = new(DestroyVirtualMachineParameter)
	p.Id.Set(id)
	return p
}

// Destroys a virtual machine. Once destroyed, only the administrator can
// recover it.
func (c *Client) DestroyVirtualMachine(p *DestroyVirtualMachineParameter) (*VirtualMachine, error) {
	obj, err := c.Request("destroyVirtualMachine", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), err
}


// ListVirtualMachines represents the paramter of ListVirtualMachines
type ListVirtualMachinesParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// list vms by affinity group
	AffinityGroupId ID
	// comma separated list of host details requested, value can be a list of [all,
	// group, nics, stats, secgrp, tmpl, servoff, diskoff, iso, volume, min,
	// affgrp]. If no parameter is passed in, the details will be defaulted to all
	Details []string
	// list resources by display flag; only ROOT admin is eligible to pass this
	// parameter
	DisplayVm NullBool
	// list only resources belonging to the domain specified
	DomainId ID
	// list by network type; true if need to list vms using Virtual Network, false
	// otherwise
	ForVirtualNetwork NullBool
	// the group ID
	GroupId ID
	// the host ID
	HostId ID
	// the target hypervisor for the template
	Hypervisor NullString
	// the ID of the virtual machine
	Id ID
	// the IDs of the virtual machines, mutually exclusive with id
	Ids []string
	// list vms by iso
	IsoId ID
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll NullBool
	// name of the virtual machine
	Name NullString
	// list by network id
	NetworkId ID
	Page      NullNumber
	PageSize  NullNumber
	// the pod ID
	PodId ID
	// list objects by project
	ProjectId ID
	// list by the service offering
	ServiceOfferingId ID
	// state of the virtual machine
	State NullString
	// the storage ID where vm's volumes belong to
	StorageId ID
	// List resources by tags (key/value pairs)
	Tags map[string]string
	// list vms by template
	TemplateId ID
	// list vms by vpc
	VpcId ID
	// the availability zone ID
	ZoneId ID
}

func NewListVirtualMachinesParameter() (p *ListVirtualMachinesParameter) {
	p = new(ListVirtualMachinesParameter)
	return p
}

// List the virtual machines owned by the account.
func (c *Client) ListVirtualMachines(p *ListVirtualMachinesParameter) ([]*VirtualMachine, error) {
	obj, err := c.Request("listVirtualMachines", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*VirtualMachine), err
}

