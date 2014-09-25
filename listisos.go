package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type ListIsosParameter struct {
	Account     NullString
	Bootable    NullBool
	Domainid    ID
	Hypervisor  NullString
	Id          ID
	Isofilter   NullString
	Ispublic    NullBool
	Isready     NullBool
	Isrecursive NullBool
	Keyword     NullString
	Listall     NullBool
	Name        NullString
	Page        NullInt64
	Pagesize    NullInt64
	Projectid   ID
	Showremoved NullBool
	Tags        map[string]string
	Zoneid      ID
}

func (p *ListIsosParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *ListIsosParameter) SetBootable(b bool) {
	p.Bootable.Bool = b
	p.Bootable.Valid = true
}
func (p *ListIsosParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *ListIsosParameter) SetHypervisor(s string) {
	p.Hypervisor.String = s
	p.Hypervisor.Valid = true
}
func (p *ListIsosParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *ListIsosParameter) SetIsofilter(s string) {
	p.Isofilter.String = s
	p.Isofilter.Valid = true
}
func (p *ListIsosParameter) SetIspublic(b bool) {
	p.Ispublic.Bool = b
	p.Ispublic.Valid = true
}
func (p *ListIsosParameter) SetIsready(b bool) {
	p.Isready.Bool = b
	p.Isready.Valid = true
}
func (p *ListIsosParameter) SetIsrecursive(b bool) {
	p.Isrecursive.Bool = b
	p.Isrecursive.Valid = true
}
func (p *ListIsosParameter) SetKeyword(s string) {
	p.Keyword.String = s
	p.Keyword.Valid = true
}
func (p *ListIsosParameter) SetListall(b bool) {
	p.Listall.Bool = b
	p.Listall.Valid = true
}
func (p *ListIsosParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *ListIsosParameter) SetPage(n int64) {
	p.Page.Int64 = n
	p.Page.Valid = true
}
func (p *ListIsosParameter) SetPagesize(n int64) {
	p.Pagesize.Int64 = n
	p.Pagesize.Valid = true
}
func (p *ListIsosParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *ListIsosParameter) SetShowremoved(b bool) {
	p.Showremoved.Bool = b
	p.Showremoved.Valid = true
}
func (p *ListIsosParameter) SetTags(m map[string]string) {
	p.Tags = m
}
func (p *ListIsosParameter) SetZoneid(s string) {
	p.Zoneid.String = s
	p.Zoneid.Valid = true
}
func (p *ListIsosParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if p.Bootable.Valid {
		m["bootable"] = fmt.Sprint(p.Bootable.Bool)
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Hypervisor.Valid {
		m["hypervisor"] = fmt.Sprint(p.Hypervisor.String)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Isofilter.Valid {
		m["isofilter"] = fmt.Sprint(p.Isofilter.String)
	}
	if p.Ispublic.Valid {
		m["ispublic"] = fmt.Sprint(p.Ispublic.Bool)
	}
	if p.Isready.Valid {
		m["isready"] = fmt.Sprint(p.Isready.Bool)
	}
	if p.Isrecursive.Valid {
		m["isrecursive"] = fmt.Sprint(p.Isrecursive.Bool)
	}
	if p.Keyword.Valid {
		m["keyword"] = fmt.Sprint(p.Keyword.String)
	}
	if p.Listall.Valid {
		m["listall"] = fmt.Sprint(p.Listall.Bool)
	}
	if p.Name.Valid {
		m["name"] = fmt.Sprint(p.Name.String)
	}
	if p.Page.Valid {
		m["page"] = fmt.Sprint(p.Page.Int64)
	}
	if p.Pagesize.Valid {
		m["pagesize"] = fmt.Sprint(p.Pagesize.Int64)
	}
	if p.Projectid.Valid {
		m["projectid"] = fmt.Sprint(p.Projectid.String)
	}
	if p.Showremoved.Valid {
		m["showremoved"] = fmt.Sprint(p.Showremoved.Bool)
	}
	if len(p.Tags) > 0 {
		i := 0
		for key, value := range m {
			m[fmt.Sprintf("tags[%d].key", i)] = key
			m[fmt.Sprintf("tags[%d].value", i)] = value
			i += 1
		}
	}
	if p.Zoneid.Valid {
		m["zoneid"] = fmt.Sprint(p.Zoneid.String)
	}
	return m
}
func (c *Client) ListIsos(p ListIsosParameter) ([]Iso, error) {
	var v map[string]json.RawMessage
	var ret []Iso
	b, err := c.Request("listIsos", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Println("json.Unmarshal failed:", err)
	}
	content, ok := v["iso"]
	if !ok {
		log.Println("Content is empty.")
		return ret, nil
	}
	err = json.Unmarshal(content, &ret)
	if err != nil {
		log.Println("json.Unmarshal failed:", err)
	}
	return ret, err
}
