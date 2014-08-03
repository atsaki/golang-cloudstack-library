package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type ListZonesParameter struct {
	Available      NullBool
	Domainid       ID
	Id             ID
	Keyword        NullString
	Name           NullString
	Networktype    NullString
	Page           NullInt64
	Pagesize       NullInt64
	Showcapacities NullBool
	Tags           map[string]string
}

func (p *ListZonesParameter) SetAvailable(b bool) {
	p.Available.Bool = b
	p.Available.Valid = true
}
func (p *ListZonesParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *ListZonesParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *ListZonesParameter) SetKeyword(s string) {
	p.Keyword.String = s
	p.Keyword.Valid = true
}
func (p *ListZonesParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *ListZonesParameter) SetNetworktype(s string) {
	p.Networktype.String = s
	p.Networktype.Valid = true
}
func (p *ListZonesParameter) SetPage(n int64) {
	p.Page.Int64 = n
	p.Page.Valid = true
}
func (p *ListZonesParameter) SetPagesize(n int64) {
	p.Pagesize.Int64 = n
	p.Pagesize.Valid = true
}
func (p *ListZonesParameter) SetShowcapacities(b bool) {
	p.Showcapacities.Bool = b
	p.Showcapacities.Valid = true
}
func (p *ListZonesParameter) SetTags(m map[string]string) {
	p.Tags = m
}
func (p *ListZonesParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Available.Valid {
		m["available"] = fmt.Sprint(p.Available.Bool)
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Keyword.Valid {
		m["keyword"] = fmt.Sprint(p.Keyword.String)
	}
	if p.Name.Valid {
		m["name"] = fmt.Sprint(p.Name.String)
	}
	if p.Networktype.Valid {
		m["networktype"] = fmt.Sprint(p.Networktype.String)
	}
	if p.Page.Valid {
		m["page"] = fmt.Sprint(p.Page.Int64)
	}
	if p.Pagesize.Valid {
		m["pagesize"] = fmt.Sprint(p.Pagesize.Int64)
	}
	if p.Showcapacities.Valid {
		m["showcapacities"] = fmt.Sprint(p.Showcapacities.Bool)
	}
	if len(p.Tags) > 0 {
		i := 0
		for key, value := range m {
			m[fmt.Sprintf("tags[%d].key", i)] = key
			m[fmt.Sprintf("tags[%d].value", i)] = value
			i += 1
		}
	}
	return m
}

type ListZonesResponse struct {
	Count float64 `json:"count"`
	Zone  []struct {
		Allocationstate NullString `json:"allocationstate"`
		Capacity        []struct {
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
		} `json:"capacity"`
		Description           NullString `json:"description"`
		Dhcpprovider          NullString `json:"dhcpprovider"`
		Displaytext           NullString `json:"displaytext"`
		Dns1                  NullString `json:"dns1"`
		Dns2                  NullString `json:"dns2"`
		Domain                NullString `json:"domain"`
		Domainid              ID         `json:"domainid"`
		Domainname            NullString `json:"domainname"`
		Guestcidraddress      NullString `json:"guestcidraddress"`
		Id                    ID         `json:"id"`
		Internaldns1          NullString `json:"internaldns1"`
		Internaldns2          NullString `json:"internaldns2"`
		Ip6dns1               NullString `json:"ip6dns1"`
		Ip6dns2               NullString `json:"ip6dns2"`
		Localstorageenabled   NullBool   `json:"localstorageenabled"`
		Name                  NullString `json:"name"`
		Networktype           NullString `json:"networktype"`
		Resourcedetails       NullString `json:"resourcedetails"`
		Securitygroupsenabled NullBool   `json:"securitygroupsenabled"`
		Tags                  []struct {
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
		} `json:"tags"`
		Vlan      NullString `json:"vlan"`
		Zonetoken NullString `json:"zonetoken"`
	} `json:"zone"`
}

func (c *Client) ListZones(p ListZonesParameter) (*ListZonesResponse, error) {
	resp := new(ListZonesResponse)
	b, err := c.Request("listZones", p.ToMap())
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
