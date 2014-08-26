package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type ListTemplatesParameter struct {
	Account        NullString
	Domainid       ID
	Hypervisor     NullString
	Id             ID
	Isrecursive    NullBool
	Keyword        NullString
	Listall        NullBool
	Name           NullString
	Page           NullInt64
	Pagesize       NullInt64
	Projectid      ID
	Showremoved    NullBool
	Tags           map[string]string
	Templatefilter NullString
	Zoneid         ID
}

func (p *ListTemplatesParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *ListTemplatesParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *ListTemplatesParameter) SetHypervisor(s string) {
	p.Hypervisor.String = s
	p.Hypervisor.Valid = true
}
func (p *ListTemplatesParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *ListTemplatesParameter) SetIsrecursive(b bool) {
	p.Isrecursive.Bool = b
	p.Isrecursive.Valid = true
}
func (p *ListTemplatesParameter) SetKeyword(s string) {
	p.Keyword.String = s
	p.Keyword.Valid = true
}
func (p *ListTemplatesParameter) SetListall(b bool) {
	p.Listall.Bool = b
	p.Listall.Valid = true
}
func (p *ListTemplatesParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *ListTemplatesParameter) SetPage(n int64) {
	p.Page.Int64 = n
	p.Page.Valid = true
}
func (p *ListTemplatesParameter) SetPagesize(n int64) {
	p.Pagesize.Int64 = n
	p.Pagesize.Valid = true
}
func (p *ListTemplatesParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *ListTemplatesParameter) SetShowremoved(b bool) {
	p.Showremoved.Bool = b
	p.Showremoved.Valid = true
}
func (p *ListTemplatesParameter) SetTags(m map[string]string) {
	p.Tags = m
}
func (p *ListTemplatesParameter) SetTemplatefilter(s string) {
	p.Templatefilter.String = s
	p.Templatefilter.Valid = true
}
func (p *ListTemplatesParameter) SetZoneid(s string) {
	p.Zoneid.String = s
	p.Zoneid.Valid = true
}
func (p *ListTemplatesParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
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
	if p.Templatefilter.Valid {
		m["templatefilter"] = fmt.Sprint(p.Templatefilter.String)
	}
	if p.Zoneid.Valid {
		m["zoneid"] = fmt.Sprint(p.Zoneid.String)
	}
	return m
}

type ListTemplatesResponse struct {
	Count    float64 `json:"count"`
	Template []struct {
		Account               NullString `json:"account"`
		Accountid             ID         `json:"accountid"`
		Bootable              NullBool   `json:"bootable"`
		Checksum              NullString `json:"checksum"`
		Created               NullString `json:"created"`
		CrossZones            NullBool   `json:"crosszones"`
		Details               NullString `json:"details"`
		Displaytext           NullString `json:"displaytext"`
		Domain                NullString `json:"domain"`
		Domainid              ID         `json:"domainid"`
		Format                NullString `json:"format"`
		Hostid                ID         `json:"hostid"`
		Hostname              NullString `json:"hostname"`
		Hypervisor            NullString `json:"hypervisor"`
		Id                    ID         `json:"id"`
		Isdynamicallyscalable NullBool   `json:"isdynamicallyscalable"`
		Isextractable         NullBool   `json:"isextractable"`
		Isfeatured            NullBool   `json:"isfeatured"`
		Ispublic              NullBool   `json:"ispublic"`
		Isready               NullBool   `json:"isready"`
		Name                  NullString `json:"name"`
		Ostypeid              ID         `json:"ostypeid"`
		Ostypename            NullString `json:"ostypename"`
		Passwordenabled       NullBool   `json:"passwordenabled"`
		Project               NullString `json:"project"`
		Projectid             ID         `json:"projectid"`
		Removed               NullString `json:"removed"`
		Size                  NullInt64  `json:"size"`
		Sourcetemplateid      ID         `json:"sourcetemplateid"`
		Sshkeyenabled         NullBool   `json:"sshkeyenabled"`
		Status                NullString `json:"status"`
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
		Templatetag  NullString `json:"templatetag"`
		Templatetype NullString `json:"templatetype"`
		Zoneid       ID         `json:"zoneid"`
		Zonename     NullString `json:"zonename"`
	} `json:"template"`
}

func (c *Client) ListTemplates(p ListTemplatesParameter) (*ListTemplatesResponse, error) {
	resp := new(ListTemplatesResponse)
	b, err := c.Request("listTemplates", p.ToMap())
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