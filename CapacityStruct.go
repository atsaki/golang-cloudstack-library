package cloudstack

type Capacity struct {
	ResourceBase
	// the total capacity available
	CapacityTotal NullNumber `json:"capacitytotal"`
	// the capacity currently in use
	CapacityUsed NullNumber `json:"capacityused"`
	// the Cluster ID
	ClusterId ID `json:"clusterid"`
	// the Cluster name
	ClusterName NullString `json:"clustername"`
	// the percentage of capacity currently in use
	PercentUsed NullString `json:"percentused"`
	// the Pod ID
	PodId ID `json:"podid"`
	// the Pod name
	PodName NullString `json:"podname"`
	// the capacity type
	Type NullNumber `json:"type"`
	// the Zone ID
	ZoneId ID `json:"zoneid"`
	// the Zone name
	ZoneName NullString `json:"zonename"`
}
