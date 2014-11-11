package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type ListServiceOfferingsParameter struct {
	Domainid         ID
	Id               ID
	Issystem         NullBool
	Keyword          NullString
	Name             NullString
	Page             NullInt64
	Pagesize         NullInt64
	Systemvmtype     NullString
	Virtualmachineid ID
}

func (p *ListServiceOfferingsParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *ListServiceOfferingsParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *ListServiceOfferingsParameter) SetIssystem(b bool) {
	p.Issystem.Bool = b
	p.Issystem.Valid = true
}
func (p *ListServiceOfferingsParameter) SetKeyword(s string) {
	p.Keyword.String = s
	p.Keyword.Valid = true
}
func (p *ListServiceOfferingsParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *ListServiceOfferingsParameter) SetPage(n int64) {
	p.Page.Int64 = n
	p.Page.Valid = true
}
func (p *ListServiceOfferingsParameter) SetPagesize(n int64) {
	p.Pagesize.Int64 = n
	p.Pagesize.Valid = true
}
func (p *ListServiceOfferingsParameter) SetSystemvmtype(s string) {
	p.Systemvmtype.String = s
	p.Systemvmtype.Valid = true
}
func (p *ListServiceOfferingsParameter) SetVirtualmachineid(s string) {
	p.Virtualmachineid.String = s
	p.Virtualmachineid.Valid = true
}
func (p *ListServiceOfferingsParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Issystem.Valid {
		m["issystem"] = fmt.Sprint(p.Issystem.Bool)
	}
	if p.Keyword.Valid {
		m["keyword"] = fmt.Sprint(p.Keyword.String)
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
	if p.Systemvmtype.Valid {
		m["systemvmtype"] = fmt.Sprint(p.Systemvmtype.String)
	}
	if p.Virtualmachineid.Valid {
		m["virtualmachineid"] = fmt.Sprint(p.Virtualmachineid.String)
	}
	return m
}
func (c *Client) ListServiceOfferings(p ListServiceOfferingsParameter) ([]Serviceoffering, error) {
	var v map[string]json.RawMessage
	var ret []Serviceoffering
	b, err := c.Request("listServiceOfferings", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["serviceoffering"]
	if !ok {
		if len(v) == 0 {
			return ret, nil
		}
		errortext, _ := v["errortext"]
		if ok {
			return ret, fmt.Errorf(string(errortext))
		} else {
			return ret, fmt.Errorf("Unexpected format")
		}
	}
	err = json.Unmarshal(content, &ret)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(content))
	}
	return ret, nil
}
