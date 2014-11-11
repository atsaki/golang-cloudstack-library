package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type ListPublicIpAddressesParameter struct {
	Account             NullString
	Allocatedonly       NullBool
	Associatednetworkid ID
	Domainid            ID
	Fordisplay          NullBool
	Forloadbalancing    NullBool
	Forvirtualnetwork   NullBool
	Id                  ID
	Ipaddress           NullString
	Isrecursive         NullBool
	Issourcenat         NullBool
	Isstaticnat         NullBool
	Keyword             NullString
	Listall             NullBool
	Page                NullInt64
	Pagesize            NullInt64
	Physicalnetworkid   ID
	Projectid           ID
	Tags                map[string]string
	Vlanid              ID
	Vpcid               ID
	Zoneid              ID
}

func (p *ListPublicIpAddressesParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetAllocatedonly(b bool) {
	p.Allocatedonly.Bool = b
	p.Allocatedonly.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetAssociatednetworkid(s string) {
	p.Associatednetworkid.String = s
	p.Associatednetworkid.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetFordisplay(b bool) {
	p.Fordisplay.Bool = b
	p.Fordisplay.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetForloadbalancing(b bool) {
	p.Forloadbalancing.Bool = b
	p.Forloadbalancing.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetForvirtualnetwork(b bool) {
	p.Forvirtualnetwork.Bool = b
	p.Forvirtualnetwork.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetIpaddress(s string) {
	p.Ipaddress.String = s
	p.Ipaddress.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetIsrecursive(b bool) {
	p.Isrecursive.Bool = b
	p.Isrecursive.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetIssourcenat(b bool) {
	p.Issourcenat.Bool = b
	p.Issourcenat.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetIsstaticnat(b bool) {
	p.Isstaticnat.Bool = b
	p.Isstaticnat.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetKeyword(s string) {
	p.Keyword.String = s
	p.Keyword.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetListall(b bool) {
	p.Listall.Bool = b
	p.Listall.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetPage(n int64) {
	p.Page.Int64 = n
	p.Page.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetPagesize(n int64) {
	p.Pagesize.Int64 = n
	p.Pagesize.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetPhysicalnetworkid(s string) {
	p.Physicalnetworkid.String = s
	p.Physicalnetworkid.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetTags(m map[string]string) {
	p.Tags = m
}
func (p *ListPublicIpAddressesParameter) SetVlanid(s string) {
	p.Vlanid.String = s
	p.Vlanid.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetVpcid(s string) {
	p.Vpcid.String = s
	p.Vpcid.Valid = true
}
func (p *ListPublicIpAddressesParameter) SetZoneid(s string) {
	p.Zoneid.String = s
	p.Zoneid.Valid = true
}
func (p *ListPublicIpAddressesParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if p.Allocatedonly.Valid {
		m["allocatedonly"] = fmt.Sprint(p.Allocatedonly.Bool)
	}
	if p.Associatednetworkid.Valid {
		m["associatednetworkid"] = fmt.Sprint(p.Associatednetworkid.String)
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Fordisplay.Valid {
		m["fordisplay"] = fmt.Sprint(p.Fordisplay.Bool)
	}
	if p.Forloadbalancing.Valid {
		m["forloadbalancing"] = fmt.Sprint(p.Forloadbalancing.Bool)
	}
	if p.Forvirtualnetwork.Valid {
		m["forvirtualnetwork"] = fmt.Sprint(p.Forvirtualnetwork.Bool)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Ipaddress.Valid {
		m["ipaddress"] = fmt.Sprint(p.Ipaddress.String)
	}
	if p.Isrecursive.Valid {
		m["isrecursive"] = fmt.Sprint(p.Isrecursive.Bool)
	}
	if p.Issourcenat.Valid {
		m["issourcenat"] = fmt.Sprint(p.Issourcenat.Bool)
	}
	if p.Isstaticnat.Valid {
		m["isstaticnat"] = fmt.Sprint(p.Isstaticnat.Bool)
	}
	if p.Keyword.Valid {
		m["keyword"] = fmt.Sprint(p.Keyword.String)
	}
	if p.Listall.Valid {
		m["listall"] = fmt.Sprint(p.Listall.Bool)
	}
	if p.Page.Valid {
		m["page"] = fmt.Sprint(p.Page.Int64)
	}
	if p.Pagesize.Valid {
		m["pagesize"] = fmt.Sprint(p.Pagesize.Int64)
	}
	if p.Physicalnetworkid.Valid {
		m["physicalnetworkid"] = fmt.Sprint(p.Physicalnetworkid.String)
	}
	if p.Projectid.Valid {
		m["projectid"] = fmt.Sprint(p.Projectid.String)
	}
	if len(p.Tags) > 0 {
		i := 0
		for key, value := range m {
			m[fmt.Sprintf("tags[%d].key", i)] = key
			m[fmt.Sprintf("tags[%d].value", i)] = value
			i += 1
		}
	}
	if p.Vlanid.Valid {
		m["vlanid"] = fmt.Sprint(p.Vlanid.String)
	}
	if p.Vpcid.Valid {
		m["vpcid"] = fmt.Sprint(p.Vpcid.String)
	}
	if p.Zoneid.Valid {
		m["zoneid"] = fmt.Sprint(p.Zoneid.String)
	}
	return m
}
func (c *Client) ListPublicIpAddresses(p ListPublicIpAddressesParameter) ([]Publicipaddress, error) {
	var v map[string]json.RawMessage
	var ret []Publicipaddress
	b, err := c.Request("listPublicIpAddresses", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["publicipaddress"]
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
