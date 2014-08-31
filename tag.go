package cloudstack

type Tag struct {
	Account      NullString `json:"account"`
	Customer     NullString `json:"customer"`
	Domain       NullString `json:"domain"`
	Domainid     ID         `json:"domainid"`
	Key          NullString `json:"key"`
	Project      NullString `json:"project"`
	Projectid    ID         `json:"projectid"`
	Resourceid   ID         `json:"resourceid"`
	Resourcetype NullString `json:"resourcetype"`
	Value        NullString `json:"value"`
}
