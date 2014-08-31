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
func (c *Client) ListZones(p ListZonesParameter) ([]Zone, error) {
	var v map[string]json.RawMessage
	var ret []Zone
	b, err := c.Request("listZones", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Println("json.Unmarshal failed:", err)
	}
	err = json.Unmarshal(v["zone"], &ret)
	if err != nil {
		log.Println("json.Unmarshal failed:", err)
	}
	return ret, err
}
