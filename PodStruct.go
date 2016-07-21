package cloudstack

type Pod struct {
	ResourceBase
	// the allocation state of the Pod
	AllocationState NullString `json:"allocationstate"`
	// the capacity of the Pod
	Capacity []Capacity `json:"capacity"`
	// the ending IP for the Pod
	EndIp NullString `json:"endip"`
	// the gateway of the Pod
	Gateway NullString `json:"gateway"`
	// the ID of the Pod
	Id ID `json:"id"`
	// the name of the Pod
	Name NullString `json:"name"`
	// the netmask of the Pod
	Netmask NullString `json:"netmask"`
	// the starting IP for the Pod
	StartIp NullString `json:"startip"`
	// the Zone ID of the Pod
	ZoneId ID `json:"zoneid"`
	// the Zone name of the Pod
	ZoneName NullString `json:"zonename"`
}

type DedicatedPod struct {
	ResourceBase
	// the Account Id to which the Pod is dedicated
	AccountId ID `json:"accountid"`
	// the Dedication Affinity Group ID of the pod
	AffinityGroupId ID `json:"affinitygroupid"`
	// the domain ID to which the Pod is dedicated
	DomainId ID `json:"domainid"`
	// the ID of the dedicated resource
	Id ID `json:"id"`
	// the ID of the Pod
	PodId ID `json:"podid"`
	// the Name of the Pod
	PodName NullString `json:"podname"`
}
