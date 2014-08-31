package cloudstack

type Networkoffering struct {
	Availability             NullString       `json:"availability"`
	Conservemode             NullBool         `json:"conservemode"`
	Created                  NullString       `json:"created"`
	Details                  NullString       `json:"details"`
	Displaytext              NullString       `json:"displaytext"`
	Egressdefaultpolicy      NullBool         `json:"egressdefaultpolicy"`
	Forvpc                   NullBool         `json:"forvpc"`
	Guestiptype              NullString       `json:"guestiptype"`
	Id                       ID               `json:"id"`
	Isdefault                NullBool         `json:"isdefault"`
	Ispersistent             NullBool         `json:"ispersistent"`
	Maxconnections           NullInt64        `json:"maxconnections"`
	Name                     NullString       `json:"name"`
	Networkrate              NullInt64        `json:"networkrate"`
	Service                  []Networkservice `json:"service"`
	Serviceofferingid        ID               `json:"serviceofferingid"`
	Specifyipranges          NullBool         `json:"specifyipranges"`
	Specifyvlan              NullBool         `json:"specifyvlan"`
	State                    NullString       `json:"state"`
	Supportsstrechedl2subnet NullBool         `json:"supportsstrechedl2subnet"`
	Tags                     NullString       `json:"tags"`
	Traffictype              NullString       `json:"traffictype"`
}
