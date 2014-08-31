package cloudstack

type Capacity struct {
	Capacitytotal NullInt64  `json:"capacitytotal"`
	Capacityused  NullInt64  `json:"capacityused"`
	Clusterid     ID         `json:"clusterid"`
	Clustername   NullString `json:"clustername"`
	Percentused   NullString `json:"percentused"`
	Podid         ID         `json:"podid"`
	Podname       NullString `json:"podname"`
	Type          NullInt64  `json:"type"`
	Zoneid        ID         `json:"zoneid"`
	Zonename      NullString `json:"zonename"`
}
