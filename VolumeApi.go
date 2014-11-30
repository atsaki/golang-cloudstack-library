package cloudstack


// ExtractVolume represents the paramter of ExtractVolume
type ExtractVolumeParameter struct {
	// the ID of the volume
	Id ID
	// the mode of extraction - HTTP_DOWNLOAD or FTP_UPLOAD
	Mode NullString
	// the url to which the volume would be extracted
	Url NullString
	// the ID of the zone where the volume is located
	ZoneId ID
}

func NewExtractVolumeParameter(id string, mode string, zoneid string) (p *ExtractVolumeParameter) {
	p = new(ExtractVolumeParameter)
	p.Id.Set(id)
	p.Mode.Set(mode)
	p.ZoneId.Set(zoneid)
	return p
}

// Extracts volume
func (c *Client) ExtractVolume(p *ExtractVolumeParameter) (*Volume, error) {
	obj, err := c.Request("extractVolume", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Volume), err
}


// ListVolumes represents the paramter of ListVolumes
type ListVolumesParameter struct {
	// list resources by account. Must be used with the domainId parameter.
	Account NullString
	// list volumes by disk offering
	DiskOfferingId ID
	// list resources by display flag; only ROOT admin is eligible to pass this
	// parameter
	DisplayVolume NullBool
	// list only resources belonging to the domain specified
	DomainId ID
	// list volumes on specified host
	HostId ID
	// the ID of the disk volume
	Id ID
	// defaults to false, but if true, lists all resources from the parent specified
	// by the domainId till leaves.
	IsRecursive NullBool
	// List by keyword
	Keyword NullString
	// If set to false, list only resources belonging to the command's caller; if
	// set to true - list resources that the caller is authorized to see. Default
	// value is false
	ListAll NullBool
	// the name of the disk volume
	Name     NullString
	Page     NullNumber
	PageSize NullNumber
	// the pod id the disk volume belongs to
	PodId ID
	// list objects by project
	ProjectId ID
	// the ID of the storage pool, available to ROOT admin only
	StorageId ID
	// List resources by tags (key/value pairs)
	Tags map[string]string
	// the type of disk volume
	Type NullString
	// the ID of the virtual machine
	VirtualMachineId ID
	// the ID of the availability zone
	ZoneId ID
}

func NewListVolumesParameter() (p *ListVolumesParameter) {
	p = new(ListVolumesParameter)
	return p
}

// Lists all volumes.
func (c *Client) ListVolumes(p *ListVolumesParameter) ([]*Volume, error) {
	obj, err := c.Request("listVolumes", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.([]*Volume), err
}


// MigrateVolume represents the paramter of MigrateVolume
type MigrateVolumeParameter struct {
	// if the volume should be live migrated when it is attached to a running vm
	LiveMigrate NullBool
	// destination storage pool ID to migrate the volume to
	StorageId ID
	// the ID of the volume
	VolumeId ID
}

func NewMigrateVolumeParameter(storageid string, volumeid string) (p *MigrateVolumeParameter) {
	p = new(MigrateVolumeParameter)
	p.StorageId.Set(storageid)
	p.VolumeId.Set(volumeid)
	return p
}

// Migrate volume
func (c *Client) MigrateVolume(p *MigrateVolumeParameter) (*Volume, error) {
	obj, err := c.Request("migrateVolume", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Volume), err
}


// AttachVolume represents the paramter of AttachVolume
type AttachVolumeParameter struct {
	// the ID of the device to map the volume to within the guest OS. If no deviceId
	// is passed in, the next available deviceId will be chosen. Possible values for
	// a Linux OS are:* 1 - /dev/xvdb* 2 - /dev/xvdc* 4 - /dev/xvde* 5 - /dev/xvdf*
	// 6 - /dev/xvdg* 7 - /dev/xvdh* 8 - /dev/xvdi* 9 - /dev/xvdj
	DeviceId ID
	// the ID of the disk volume
	Id ID
	//     the ID of the virtual machine
	VirtualMachineId ID
}

func NewAttachVolumeParameter(id string, virtualmachineid string) (p *AttachVolumeParameter) {
	p = new(AttachVolumeParameter)
	p.Id.Set(id)
	p.VirtualMachineId.Set(virtualmachineid)
	return p
}

// Attaches a disk volume to a virtual machine.
func (c *Client) AttachVolume(p *AttachVolumeParameter) (*Volume, error) {
	obj, err := c.Request("attachVolume", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Volume), err
}


// UpdateVolume represents the paramter of UpdateVolume
type UpdateVolumeParameter struct {
	// The chain info of the volume
	ChainInfo NullString
	// an optional field, in case you want to set a custom id to the resource.
	// Allowed to Root Admins only
	CustomId ID
	// an optional field, whether to the display the volume to the end user or not.
	DisplayVolume NullBool
	// the ID of the disk volume
	Id ID
	// The path of the volume
	Path NullString
	// The state of the volume
	State NullString
	// Destination storage pool UUID for the volume
	StorageId ID
}

func NewUpdateVolumeParameter() (p *UpdateVolumeParameter) {
	p = new(UpdateVolumeParameter)
	return p
}

// Updates the volume.
func (c *Client) UpdateVolume(p *UpdateVolumeParameter) (*Volume, error) {
	obj, err := c.Request("updateVolume", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Volume), err
}


// ResizeVolume represents the paramter of ResizeVolume
type ResizeVolumeParameter struct {
	// new disk offering id
	DiskOfferingId ID
	// the ID of the disk volume
	Id ID
	// Verify OK to Shrink
	ShrinkOk NullBool
	// New volume size in G
	Size NullNumber
}

func NewResizeVolumeParameter(id string) (p *ResizeVolumeParameter) {
	p = new(ResizeVolumeParameter)
	p.Id.Set(id)
	return p
}

// Resizes a volume
func (c *Client) ResizeVolume(p *ResizeVolumeParameter) (*Volume, error) {
	obj, err := c.Request("resizeVolume", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Volume), err
}


// DetachVolume represents the paramter of DetachVolume
type DetachVolumeParameter struct {
	// the device ID on the virtual machine where volume is detached from
	DeviceId ID
	// the ID of the disk volume
	Id ID
	// the ID of the virtual machine where the volume is detached from
	VirtualMachineId ID
}

func NewDetachVolumeParameter() (p *DetachVolumeParameter) {
	p = new(DetachVolumeParameter)
	return p
}

// Detaches a disk volume from a virtual machine.
func (c *Client) DetachVolume(p *DetachVolumeParameter) (*Volume, error) {
	obj, err := c.Request("detachVolume", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Volume), err
}


// DeleteVolume represents the paramter of DeleteVolume
type DeleteVolumeParameter struct {
	// The ID of the disk volume
	Id ID
}

func NewDeleteVolumeParameter(id string) (p *DeleteVolumeParameter) {
	p = new(DeleteVolumeParameter)
	p.Id.Set(id)
	return p
}

// Deletes a detached disk volume.
func (c *Client) DeleteVolume(p *DeleteVolumeParameter) (*Result, error) {
	obj, err := c.Request("deleteVolume", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Result), err
}


// CreateVolume represents the paramter of CreateVolume
type CreateVolumeParameter struct {
	// the account associated with the disk volume. Must be used with the domainId
	// parameter.
	Account NullString
	// an optional field, in case you want to set a custom id to the resource.
	// Allowed to Root Admins only
	CustomId ID
	// the ID of the disk offering. Either diskOfferingId or snapshotId must be
	// passed in.
	DiskOfferingId ID
	// an optional field, whether to display the volume to the end user or not.
	DisplayVolume NullBool
	// the domain ID associated with the disk offering. If used with the account
	// parameter returns the disk volume associated with the account for the
	// specified domain.
	DomainId ID
	// max iops
	MaxIops NullNumber
	// min iops
	MinIops NullNumber
	// the name of the disk volume
	Name NullString
	// the project associated with the volume. Mutually exclusive with account
	// parameter
	ProjectId ID
	// Arbitrary volume size
	Size NullNumber
	// the snapshot ID for the disk volume. Either diskOfferingId or snapshotId must
	// be passed in.
	SnapshotId ID
	// the ID of the virtual machine; to be used with snapshot Id, VM to which the
	// volume gets attached after creation
	VirtualMachineId ID
	// the ID of the availability zone
	ZoneId ID
}

func NewCreateVolumeParameter(name string) (p *CreateVolumeParameter) {
	p = new(CreateVolumeParameter)
	p.Name.Set(name)
	return p
}

// Creates a disk volume from a disk offering. This disk volume must still be
// attached to a virtual machine to make use of it.
func (c *Client) CreateVolume(p *CreateVolumeParameter) (*Volume, error) {
	obj, err := c.Request("createVolume", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Volume), err
}


// UploadVolume represents the paramter of UploadVolume
type UploadVolumeParameter struct {
	// an optional accountName. Must be used with domainId.
	Account NullString
	// the MD5 checksum value of this volume
	Checksum NullString
	// an optional domainId. If the account parameter is used, domainId must also be
	// used.
	DomainId ID
	// the format for the volume. Possible values include QCOW2, OVA, and VHD.
	Format NullString
	// Image store uuid
	ImageStoreUuid ID
	// the name of the volume
	Name NullString
	// Upload volume for the project
	ProjectId ID
	// the URL of where the volume is hosted. Possible URL include http:// and
	// https://
	Url NullString
	// the ID of the zone the volume is to be hosted on
	ZoneId ID
}

func NewUploadVolumeParameter(format string, name string, url string, zoneid string) (p *UploadVolumeParameter) {
	p = new(UploadVolumeParameter)
	p.Format.Set(format)
	p.Name.Set(name)
	p.Url.Set(url)
	p.ZoneId.Set(zoneid)
	return p
}

// Uploads a data disk.
func (c *Client) UploadVolume(p *UploadVolumeParameter) (*Volume, error) {
	obj, err := c.Request("uploadVolume", convertParamToMap(p))
	if err != nil {
		return nil, err
	}
	return obj.(*Volume), err
}

