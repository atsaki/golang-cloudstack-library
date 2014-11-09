package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type ListNetworksParameter struct {
	Account           NullString
	Acltype           NullString
	Canusefordeploy   NullBool
	Displaynetwork    NullBool
	Domainid          ID
	Forvpc            NullBool
	Id                ID
	Isrecursive       NullBool
	Issystem          NullBool
	Keyword           NullString
	Listall           NullBool
	Page              NullInt64
	Pagesize          NullInt64
	Physicalnetworkid ID
	Projectid         ID
	Restartrequired   NullBool
	Specifyipranges   NullBool
	Supportedservices []string
	Tags              map[string]string
	Traffictype       NullString
	Type              NullString
	Vpcid             ID
	Zoneid            ID
}

func (p *ListNetworksParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *ListNetworksParameter) SetAcltype(s string) {
	p.Acltype.String = s
	p.Acltype.Valid = true
}
func (p *ListNetworksParameter) SetCanusefordeploy(b bool) {
	p.Canusefordeploy.Bool = b
	p.Canusefordeploy.Valid = true
}
func (p *ListNetworksParameter) SetDisplaynetwork(b bool) {
	p.Displaynetwork.Bool = b
	p.Displaynetwork.Valid = true
}
func (p *ListNetworksParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *ListNetworksParameter) SetForvpc(b bool) {
	p.Forvpc.Bool = b
	p.Forvpc.Valid = true
}
func (p *ListNetworksParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *ListNetworksParameter) SetIsrecursive(b bool) {
	p.Isrecursive.Bool = b
	p.Isrecursive.Valid = true
}
func (p *ListNetworksParameter) SetIssystem(b bool) {
	p.Issystem.Bool = b
	p.Issystem.Valid = true
}
func (p *ListNetworksParameter) SetKeyword(s string) {
	p.Keyword.String = s
	p.Keyword.Valid = true
}
func (p *ListNetworksParameter) SetListall(b bool) {
	p.Listall.Bool = b
	p.Listall.Valid = true
}
func (p *ListNetworksParameter) SetPage(n int64) {
	p.Page.Int64 = n
	p.Page.Valid = true
}
func (p *ListNetworksParameter) SetPagesize(n int64) {
	p.Pagesize.Int64 = n
	p.Pagesize.Valid = true
}
func (p *ListNetworksParameter) SetPhysicalnetworkid(s string) {
	p.Physicalnetworkid.String = s
	p.Physicalnetworkid.Valid = true
}
func (p *ListNetworksParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *ListNetworksParameter) SetRestartrequired(b bool) {
	p.Restartrequired.Bool = b
	p.Restartrequired.Valid = true
}
func (p *ListNetworksParameter) SetSpecifyipranges(b bool) {
	p.Specifyipranges.Bool = b
	p.Specifyipranges.Valid = true
}
func (p *ListNetworksParameter) SetSupportedservices(l []string) {
	p.Supportedservices = l
}
func (p *ListNetworksParameter) SetTags(m map[string]string) {
	p.Tags = m
}
func (p *ListNetworksParameter) SetTraffictype(s string) {
	p.Traffictype.String = s
	p.Traffictype.Valid = true
}
func (p *ListNetworksParameter) SetType(s string) {
	p.Type.String = s
	p.Type.Valid = true
}
func (p *ListNetworksParameter) SetVpcid(s string) {
	p.Vpcid.String = s
	p.Vpcid.Valid = true
}
func (p *ListNetworksParameter) SetZoneid(s string) {
	p.Zoneid.String = s
	p.Zoneid.Valid = true
}
func (p *ListNetworksParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if p.Acltype.Valid {
		m["acltype"] = fmt.Sprint(p.Acltype.String)
	}
	if p.Canusefordeploy.Valid {
		m["canusefordeploy"] = fmt.Sprint(p.Canusefordeploy.Bool)
	}
	if p.Displaynetwork.Valid {
		m["displaynetwork"] = fmt.Sprint(p.Displaynetwork.Bool)
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Forvpc.Valid {
		m["forvpc"] = fmt.Sprint(p.Forvpc.Bool)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Isrecursive.Valid {
		m["isrecursive"] = fmt.Sprint(p.Isrecursive.Bool)
	}
	if p.Issystem.Valid {
		m["issystem"] = fmt.Sprint(p.Issystem.Bool)
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
	if p.Restartrequired.Valid {
		m["restartrequired"] = fmt.Sprint(p.Restartrequired.Bool)
	}
	if p.Specifyipranges.Valid {
		m["specifyipranges"] = fmt.Sprint(p.Specifyipranges.Bool)
	}
	if len(p.Supportedservices) > 0 {
		m["supportedservices"] = strings.Join(p.Supportedservices, ",")
	}
	if len(p.Tags) > 0 {
		i := 0
		for key, value := range m {
			m[fmt.Sprintf("tags[%d].key", i)] = key
			m[fmt.Sprintf("tags[%d].value", i)] = value
			i += 1
		}
	}
	if p.Traffictype.Valid {
		m["traffictype"] = fmt.Sprint(p.Traffictype.String)
	}
	if p.Type.Valid {
		m["type"] = fmt.Sprint(p.Type.String)
	}
	if p.Vpcid.Valid {
		m["vpcid"] = fmt.Sprint(p.Vpcid.String)
	}
	if p.Zoneid.Valid {
		m["zoneid"] = fmt.Sprint(p.Zoneid.String)
	}
	return m
}
func (c *Client) ListNetworks(p ListNetworksParameter) ([]Network, error) {
	var v map[string]json.RawMessage
	var ret []Network
	b, err := c.Request("listNetworks", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["network"]
	if !ok {
		errortext, _ := v["errortext"]
		return ret, fmt.Errorf(string(errortext))
	}
	err = json.Unmarshal(content, &ret)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(content))
	}
	return ret, nil
}
