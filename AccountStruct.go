package cloudstack

type Account struct {
	ResourceBase
	// details for the account
	AccountDetails NullString `json:"accountdetails"`
	// account type (admin, domain-admin, user)
	AccountType NullNumber `json:"accounttype"`
	// the total number of cpu cores available to be created for this account
	CpuAvailable NullString `json:"cpuavailable"`
	// the total number of cpu cores the account can own
	CpuLimit NullString `json:"cpulimit"`
	// the total number of cpu cores owned by account
	CpuTotal NullNumber `json:"cputotal"`
	// the default zone of the account
	DefaultZoneId ID `json:"defaultzoneid"`
	// name of the Domain the account belongs too
	Domain NullString `json:"domain"`
	// id of the Domain the account belongs too
	DomainId ID `json:"domainid"`
	// the list of acl groups that account belongs to
	Groups []NullString `json:"groups"`
	// the id of the account
	Id ID `json:"id"`
	// the total number of public ip addresses available for this account to acquire
	IpAvailable NullString `json:"ipavailable"`
	// the total number of public ip addresses this account can acquire
	IpLimit NullString `json:"iplimit"`
	// the total number of public ip addresses allocated for this account
	IpTotal NullNumber `json:"iptotal"`
	// true if the account requires cleanup
	IsCleanupRequired NullBool `json:"iscleanuprequired"`
	// true if account is default, false otherwise
	IsDefault NullBool `json:"isdefault"`
	// the total memory (in MB) available to be created for this account
	MemoryAvailable NullString `json:"memoryavailable"`
	// the total memory (in MB) the account can own
	MemoryLimit NullString `json:"memorylimit"`
	// the total memory (in MB) owned by account
	MemoryTotal NullNumber `json:"memorytotal"`
	// the name of the account
	Name NullString `json:"name"`
	// the total number of networks available to be created for this account
	NetworkAvailable NullString `json:"networkavailable"`
	// the network domain
	NetworkDomain NullString `json:"networkdomain"`
	// the total number of networks the account can own
	NetworkLimit NullString `json:"networklimit"`
	// the total number of networks owned by account
	NetworkTotal NullNumber `json:"networktotal"`
	// the total primary storage space (in GiB) available to be used for this
	// account
	PrimaryStorageAvailable NullString `json:"primarystorageavailable"`
	// the total primary storage space (in GiB) the account can own
	PrimaryStorageLimit NullString `json:"primarystoragelimit"`
	// the total primary storage space (in GiB) owned by account
	PrimaryStorageTotal NullNumber `json:"primarystoragetotal"`
	// the total number of projects available for administration by this account
	ProjectAvailable NullString `json:"projectavailable"`
	// the total number of projects the account can own
	ProjectLimit NullString `json:"projectlimit"`
	// the total number of projects being administrated by this account
	ProjectTotal NullNumber `json:"projecttotal"`
	// the total number of network traffic bytes received
	ReceivedBytes NullNumber `json:"receivedbytes"`
	// the total secondary storage space (in GiB) available to be used for this
	// account
	SecondaryStorageAvailable NullString `json:"secondarystorageavailable"`
	// the total secondary storage space (in GiB) the account can own
	SecondaryStorageLimit NullString `json:"secondarystoragelimit"`
	// the total secondary storage space (in GiB) owned by account
	SecondaryStorageTotal NullNumber `json:"secondarystoragetotal"`
	// the total number of network traffic bytes sent
	SentBytes NullNumber `json:"sentbytes"`
	// the total number of snapshots available for this account
	SnapshotAvailable NullString `json:"snapshotavailable"`
	// the total number of snapshots which can be stored by this account
	SnapshotLimit NullString `json:"snapshotlimit"`
	// the total number of snapshots stored by this account
	SnapshotTotal NullNumber `json:"snapshottotal"`
	// the state of the account
	State NullString `json:"state"`
	// the total number of templates available to be created by this account
	TemplateAvailable NullString `json:"templateavailable"`
	// the total number of templates which can be created by this account
	TemplateLimit NullString `json:"templatelimit"`
	// the total number of templates which have been created by this account
	TemplateTotal NullNumber `json:"templatetotal"`
	// the list of users associated with account
	User []User `json:"user"`
	// the total number of virtual machines available for this account to acquire
	VmAvailable NullString `json:"vmavailable"`
	// the total number of virtual machines that can be deployed by this account
	VmLimit NullString `json:"vmlimit"`
	// the total number of virtual machines running for this account
	VmRunning NullNumber `json:"vmrunning"`
	// the total number of virtual machines stopped for this account
	VmStopped NullNumber `json:"vmstopped"`
	// the total number of virtual machines deployed by this account
	VmTotal NullNumber `json:"vmtotal"`
	// the total volume available for this account
	VolumeAvailable NullString `json:"volumeavailable"`
	// the total volume which can be used by this account
	VolumeLimit NullString `json:"volumelimit"`
	// the total volume being used by this account
	VolumeTotal NullNumber `json:"volumetotal"`
	// the total number of vpcs available to be created for this account
	VpcAvailable NullString `json:"vpcavailable"`
	// the total number of vpcs the account can own
	VpcLimit NullString `json:"vpclimit"`
	// the total number of vpcs owned by account
	VpcTotal NullNumber `json:"vpctotal"`
}
