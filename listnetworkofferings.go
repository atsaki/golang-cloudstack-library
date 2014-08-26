package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type ListNetworkOfferingsParameter struct {
	Availability       NullString
	Displaytext        NullString
	Forvpc             NullBool
	Guestiptype        NullString
	Id                 ID
	Isdefault          NullBool
	Istagged           NullBool
	Keyword            NullString
	Name               NullString
	Networkid          ID
	Page               NullInt64
	Pagesize           NullInt64
	Sourcenatsupported NullBool
	Specifyipranges    NullBool
	Specifyvlan        NullBool
	State              NullString
	Supportedservices  []string
	Tags               NullString
	Traffictype        NullString
	Zoneid             ID
}

func (p *ListNetworkOfferingsParameter) SetAvailability(s string) {
	p.Availability.String = s
	p.Availability.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetDisplaytext(s string) {
	p.Displaytext.String = s
	p.Displaytext.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetForvpc(b bool) {
	p.Forvpc.Bool = b
	p.Forvpc.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetGuestiptype(s string) {
	p.Guestiptype.String = s
	p.Guestiptype.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetIsdefault(b bool) {
	p.Isdefault.Bool = b
	p.Isdefault.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetIstagged(b bool) {
	p.Istagged.Bool = b
	p.Istagged.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetKeyword(s string) {
	p.Keyword.String = s
	p.Keyword.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetNetworkid(s string) {
	p.Networkid.String = s
	p.Networkid.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetPage(n int64) {
	p.Page.Int64 = n
	p.Page.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetPagesize(n int64) {
	p.Pagesize.Int64 = n
	p.Pagesize.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetSourcenatsupported(b bool) {
	p.Sourcenatsupported.Bool = b
	p.Sourcenatsupported.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetSpecifyipranges(b bool) {
	p.Specifyipranges.Bool = b
	p.Specifyipranges.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetSpecifyvlan(b bool) {
	p.Specifyvlan.Bool = b
	p.Specifyvlan.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetState(s string) {
	p.State.String = s
	p.State.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetSupportedservices(l []string) {
	p.Supportedservices = l
}
func (p *ListNetworkOfferingsParameter) SetTags(s string) {
	p.Tags.String = s
	p.Tags.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetTraffictype(s string) {
	p.Traffictype.String = s
	p.Traffictype.Valid = true
}
func (p *ListNetworkOfferingsParameter) SetZoneid(s string) {
	p.Zoneid.String = s
	p.Zoneid.Valid = true
}
func (p *ListNetworkOfferingsParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Availability.Valid {
		m["availability"] = fmt.Sprint(p.Availability.String)
	}
	if p.Displaytext.Valid {
		m["displaytext"] = fmt.Sprint(p.Displaytext.String)
	}
	if p.Forvpc.Valid {
		m["forvpc"] = fmt.Sprint(p.Forvpc.Bool)
	}
	if p.Guestiptype.Valid {
		m["guestiptype"] = fmt.Sprint(p.Guestiptype.String)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Isdefault.Valid {
		m["isdefault"] = fmt.Sprint(p.Isdefault.Bool)
	}
	if p.Istagged.Valid {
		m["istagged"] = fmt.Sprint(p.Istagged.Bool)
	}
	if p.Keyword.Valid {
		m["keyword"] = fmt.Sprint(p.Keyword.String)
	}
	if p.Name.Valid {
		m["name"] = fmt.Sprint(p.Name.String)
	}
	if p.Networkid.Valid {
		m["networkid"] = fmt.Sprint(p.Networkid.String)
	}
	if p.Page.Valid {
		m["page"] = fmt.Sprint(p.Page.Int64)
	}
	if p.Pagesize.Valid {
		m["pagesize"] = fmt.Sprint(p.Pagesize.Int64)
	}
	if p.Sourcenatsupported.Valid {
		m["sourcenatsupported"] = fmt.Sprint(p.Sourcenatsupported.Bool)
	}
	if p.Specifyipranges.Valid {
		m["specifyipranges"] = fmt.Sprint(p.Specifyipranges.Bool)
	}
	if p.Specifyvlan.Valid {
		m["specifyvlan"] = fmt.Sprint(p.Specifyvlan.Bool)
	}
	if p.State.Valid {
		m["state"] = fmt.Sprint(p.State.String)
	}
	if len(p.Supportedservices) > 0 {
		m["supportedservices"] = strings.Join(p.Supportedservices, ",")
	}
	if p.Tags.Valid {
		m["tags"] = fmt.Sprint(p.Tags.String)
	}
	if p.Traffictype.Valid {
		m["traffictype"] = fmt.Sprint(p.Traffictype.String)
	}
	if p.Zoneid.Valid {
		m["zoneid"] = fmt.Sprint(p.Zoneid.String)
	}
	return m
}

type ListNetworkOfferingsResponse struct {
	Count           float64 `json:"count"`
	Networkoffering []struct {
		Availability        NullString `json:"availability"`
		Conservemode        NullBool   `json:"conservemode"`
		Created             NullString `json:"created"`
		Details             NullString `json:"details"`
		Displaytext         NullString `json:"displaytext"`
		Egressdefaultpolicy NullBool   `json:"egressdefaultpolicy"`
		Forvpc              NullBool   `json:"forvpc"`
		Guestiptype         NullString `json:"guestiptype"`
		Id                  ID         `json:"id"`
		Isdefault           NullBool   `json:"isdefault"`
		Ispersistent        NullBool   `json:"ispersistent"`
		Maxconnections      NullInt64  `json:"maxconnections"`
		Name                NullString `json:"name"`
		Networkrate         NullInt64  `json:"networkrate"`
		Service             []struct {
			Capability []struct {
				Canchooseservicecapability NullBool   `json:"canchooseservicecapability"`
				Name                       NullString `json:"name"`
				Value                      NullString `json:"value"`
			} `json:"capability"`
			Name     NullString `json:"name"`
			Provider []struct {
				Canenableindividualservice   NullBool     `json:"canenableindividualservice"`
				Destinationphysicalnetworkid ID           `json:"destinationphysicalnetworkid"`
				Id                           ID           `json:"id"`
				Name                         NullString   `json:"name"`
				Physicalnetworkid            ID           `json:"physicalnetworkid"`
				Servicelist                  []NullString `json:"servicelist"`
				State                        NullString   `json:"state"`
			} `json:"provider"`
		} `json:"service"`
		Serviceofferingid        ID         `json:"serviceofferingid"`
		Specifyipranges          NullBool   `json:"specifyipranges"`
		Specifyvlan              NullBool   `json:"specifyvlan"`
		State                    NullString `json:"state"`
		Supportsstrechedl2subnet NullBool   `json:"supportsstrechedl2subnet"`
		Tags                     NullString `json:"tags"`
		Traffictype              NullString `json:"traffictype"`
	} `json:"networkoffering"`
}

func (c *Client) ListNetworkOfferings(p ListNetworkOfferingsParameter) (*ListNetworkOfferingsResponse, error) {
	resp := new(ListNetworkOfferingsResponse)
	b, err := c.Request("listNetworkOfferings", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return nil, err
	}
	err = json.Unmarshal(b, resp)
	if err != nil {
		log.Println("json.Unmarshal failed:", err)
	}
	return resp, err
}
