package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type ListVolumesParameter struct {
	Account          NullString
	Diskofferingid   ID
	Displayvolume    NullBool
	Domainid         ID
	Hostid           ID
	Id               ID
	Isrecursive      NullBool
	Keyword          NullString
	Listall          NullBool
	Name             NullString
	Page             NullInt64
	Pagesize         NullInt64
	Podid            ID
	Projectid        ID
	Storageid        ID
	Tags             map[string]string
	Type             NullString
	Virtualmachineid ID
	Zoneid           ID
}

func (p *ListVolumesParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *ListVolumesParameter) SetDiskofferingid(s string) {
	p.Diskofferingid.String = s
	p.Diskofferingid.Valid = true
}
func (p *ListVolumesParameter) SetDisplayvolume(b bool) {
	p.Displayvolume.Bool = b
	p.Displayvolume.Valid = true
}
func (p *ListVolumesParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *ListVolumesParameter) SetHostid(s string) {
	p.Hostid.String = s
	p.Hostid.Valid = true
}
func (p *ListVolumesParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *ListVolumesParameter) SetIsrecursive(b bool) {
	p.Isrecursive.Bool = b
	p.Isrecursive.Valid = true
}
func (p *ListVolumesParameter) SetKeyword(s string) {
	p.Keyword.String = s
	p.Keyword.Valid = true
}
func (p *ListVolumesParameter) SetListall(b bool) {
	p.Listall.Bool = b
	p.Listall.Valid = true
}
func (p *ListVolumesParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *ListVolumesParameter) SetPage(n int64) {
	p.Page.Int64 = n
	p.Page.Valid = true
}
func (p *ListVolumesParameter) SetPagesize(n int64) {
	p.Pagesize.Int64 = n
	p.Pagesize.Valid = true
}
func (p *ListVolumesParameter) SetPodid(s string) {
	p.Podid.String = s
	p.Podid.Valid = true
}
func (p *ListVolumesParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *ListVolumesParameter) SetStorageid(s string) {
	p.Storageid.String = s
	p.Storageid.Valid = true
}
func (p *ListVolumesParameter) SetTags(m map[string]string) {
	p.Tags = m
}
func (p *ListVolumesParameter) SetType(s string) {
	p.Type.String = s
	p.Type.Valid = true
}
func (p *ListVolumesParameter) SetVirtualmachineid(s string) {
	p.Virtualmachineid.String = s
	p.Virtualmachineid.Valid = true
}
func (p *ListVolumesParameter) SetZoneid(s string) {
	p.Zoneid.String = s
	p.Zoneid.Valid = true
}
func (p *ListVolumesParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if p.Diskofferingid.Valid {
		m["diskofferingid"] = fmt.Sprint(p.Diskofferingid.String)
	}
	if p.Displayvolume.Valid {
		m["displayvolume"] = fmt.Sprint(p.Displayvolume.Bool)
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Hostid.Valid {
		m["hostid"] = fmt.Sprint(p.Hostid.String)
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
	if p.Podid.Valid {
		m["podid"] = fmt.Sprint(p.Podid.String)
	}
	if p.Projectid.Valid {
		m["projectid"] = fmt.Sprint(p.Projectid.String)
	}
	if p.Storageid.Valid {
		m["storageid"] = fmt.Sprint(p.Storageid.String)
	}
	if len(p.Tags) > 0 {
		i := 0
		for key, value := range m {
			m[fmt.Sprintf("tags[%d].key", i)] = key
			m[fmt.Sprintf("tags[%d].value", i)] = value
			i += 1
		}
	}
	if p.Type.Valid {
		m["type"] = fmt.Sprint(p.Type.String)
	}
	if p.Virtualmachineid.Valid {
		m["virtualmachineid"] = fmt.Sprint(p.Virtualmachineid.String)
	}
	if p.Zoneid.Valid {
		m["zoneid"] = fmt.Sprint(p.Zoneid.String)
	}
	return m
}
func (c *Client) ListVolumes(p ListVolumesParameter) ([]Volume, error) {
	var v map[string]json.RawMessage
	var ret []Volume
	b, err := c.Request("listVolumes", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["volume"]
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
