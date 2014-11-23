package cloudstack

type Result struct {
	DisplayText NullString `json:"displaytext"`
	Success     NullBool   `json:"success"`
}
