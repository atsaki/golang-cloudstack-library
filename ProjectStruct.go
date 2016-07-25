package cloudstack

type Project struct {
	ResourceBase
	// the account name of the project's owner
	Account NullString `json:"account"`
	// the total number of cpu cores available to be created for this project
	CpuAvailable NullString `json:"cpuavailable"`
	// the total number of cpu cores the project can own
	CpuLimit NullString `json:"cpulimit"`
	// the total number of cpu cores owned by project
	CpuTotal NullNumber `json:"cputotal"`
	// the displaytext of the project
	DisplayText NullString `json:"displaytext"`
	// the domain name where the project belongs to
	Domain NullString `json:"domain"`
	// the domain id the project belongs to
	DomainId ID `json:"domainid"`
	// the id of the project
	Id ID `json:"id"`
	// the total number of public ip addresses available for this project to acquire
	IpAvailable NullString `json:"ipavailable"`
	// the total number of public ip addresses this project can acquire
	IpLimit NullString `json:"iplimit"`
	// the total number of public ip addresses allocated for this project
	IpTotal NullNumber `json:"iptotal"`
	// the total memory (in MB) available to be created for this project
	MemoryAvailable NullString `json:"memoryavailable"`
	// the total memory (in MB) the project can own
	MemoryLimit NullString `json:"memorylimit"`
	// the total memory (in MB) owned by project
	MemoryTotal NullNumber `json:"memorytotal"`
	// the name of the project
	Name NullString `json:"name"`
	// the total number of networks available to be created for this project
	NetworkAvailable NullString `json:"networkavailable"`
	// the total number of networks the project can own
	NetworkLimit NullString `json:"networklimit"`
	// the total number of networks owned by project
	NetworkTotal NullNumber `json:"networktotal"`
	// the total primary storage space (in GiB) available to be used for this
	// project
	PrimaryStorageAvailable NullString `json:"primarystorageavailable"`
	// the total primary storage space (in GiB) the project can own
	PrimaryStorageLimit NullString `json:"primarystoragelimit"`
	// the total primary storage space (in GiB) owned by project
	PrimaryStorageTotal NullNumber `json:"primarystoragetotal"`
	// the total secondary storage space (in GiB) available to be used for this
	// project
	SecondaryStorageAvailable NullString `json:"secondarystorageavailable"`
	// the total secondary storage space (in GiB) the project can own
	SecondaryStorageLimit NullString `json:"secondarystoragelimit"`
	// the total secondary storage space (in GiB) owned by project
	SecondaryStorageTotal NullNumber `json:"secondarystoragetotal"`
	// the total number of snapshots available for this project
	SnapshotAvailable NullString `json:"snapshotavailable"`
	// the total number of snapshots which can be stored by this project
	SnapshotLimit NullString `json:"snapshotlimit"`
	// the total number of snapshots stored by this project
	SnapshotTotal NullNumber `json:"snapshottotal"`
	// the state of the project
	State NullString `json:"state"`
	// the list of resource tags associated with vm
	Tags []Tag `json:"tags"`
	// the total number of templates available to be created by this project
	TemplateAvailable NullString `json:"templateavailable"`
	// the total number of templates which can be created by this project
	TemplateLimit NullString `json:"templatelimit"`
	// the total number of templates which have been created by this project
	TemplateTotal NullNumber `json:"templatetotal"`
	// the total number of virtual machines available for this project to acquire
	VmAvailable NullString `json:"vmavailable"`
	// the total number of virtual machines that can be deployed by this project
	VmLimit NullString `json:"vmlimit"`
	// the total number of virtual machines running for this project
	VmRunning NullNumber `json:"vmrunning"`
	// the total number of virtual machines stopped for this project
	VmStopped NullNumber `json:"vmstopped"`
	// the total number of virtual machines deployed by this project
	VmTotal NullNumber `json:"vmtotal"`
	// the total volume available for this project
	VolumeAvailable NullString `json:"volumeavailable"`
	// the total volume which can be used by this project
	VolumeLimit NullString `json:"volumelimit"`
	// the total volume being used by this project
	VolumeTotal NullNumber `json:"volumetotal"`
	// the total number of vpcs available to be created for this project
	VpcAvailable NullString `json:"vpcavailable"`
	// the total number of vpcs the project can own
	VpcLimit NullString `json:"vpclimit"`
	// the total number of vpcs owned by project
	VpcTotal NullNumber `json:"vpctotal"`
}

type ProjectAccount struct {
	ResourceBase
	// project id
	ProjectId ID `json:"projectid"`
	// project name
	Project NullString `json:"project"`
	// the id of the account
	AccountId ID `json:"accountid"`
	// the name of the account
	Account NullString `json:"account"`
	// account type (admin, domain-admin, user)
	AccountType NullNumber `json:"accounttype"`
	// account role in the project (regular,owner)
	Role NullString `json:"role"`
	// id of the Domain the account belongs too
	DomainId ID `json:"domainid"`
	// name of the Domain the account belongs too
	Domain NullString `json:"domain"`
	// the list of users associated with account
	User []User `json:"user"`
}

type ProjectInvitation struct {
	ResourceBase
	// the account name of the project's owner
	Account NullString `json:"account"`
	// the domain name where the project belongs to
	Domain NullString `json:"domain"`
	// the domain id the project belongs to
	DomainId ID `json:"domainid"`
	// the email the invitation was sent to
	Email NullString `json:"email"`
	// the id of the invitation
	Id ID `json:"id"`
	// the name of the project
	Project NullString `json:"project"`
	// the id of the project
	ProjectId ID `json:"projectid"`
	// the invitation state
	State NullString `json:"state"`
}
