package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type CreateVolumeParameter struct {
	Account          NullString
	Customid         NullString
	Diskofferingid   ID
	Displayvolume    NullBool
	Domainid         ID
	Maxiops          NullInt64
	Miniops          NullInt64
	Name             NullString
	Projectid        ID
	Size             NullInt64
	Snapshotid       ID
	Virtualmachineid ID
	Zoneid           ID
}

func (p *CreateVolumeParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *CreateVolumeParameter) SetCustomid(s string) {
	p.Customid.String = s
	p.Customid.Valid = true
}
func (p *CreateVolumeParameter) SetDiskofferingid(s string) {
	p.Diskofferingid.String = s
	p.Diskofferingid.Valid = true
}
func (p *CreateVolumeParameter) SetDisplayvolume(b bool) {
	p.Displayvolume.Bool = b
	p.Displayvolume.Valid = true
}
func (p *CreateVolumeParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *CreateVolumeParameter) SetMaxiops(n int64) {
	p.Maxiops.Int64 = n
	p.Maxiops.Valid = true
}
func (p *CreateVolumeParameter) SetMiniops(n int64) {
	p.Miniops.Int64 = n
	p.Miniops.Valid = true
}
func (p *CreateVolumeParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *CreateVolumeParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *CreateVolumeParameter) SetSize(n int64) {
	p.Size.Int64 = n
	p.Size.Valid = true
}
func (p *CreateVolumeParameter) SetSnapshotid(s string) {
	p.Snapshotid.String = s
	p.Snapshotid.Valid = true
}
func (p *CreateVolumeParameter) SetVirtualmachineid(s string) {
	p.Virtualmachineid.String = s
	p.Virtualmachineid.Valid = true
}
func (p *CreateVolumeParameter) SetZoneid(s string) {
	p.Zoneid.String = s
	p.Zoneid.Valid = true
}
func (p *CreateVolumeParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if p.Customid.Valid {
		m["customid"] = fmt.Sprint(p.Customid.String)
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
	if p.Maxiops.Valid {
		m["maxiops"] = fmt.Sprint(p.Maxiops.Int64)
	}
	if p.Miniops.Valid {
		m["miniops"] = fmt.Sprint(p.Miniops.Int64)
	}
	if p.Name.Valid {
		m["name"] = fmt.Sprint(p.Name.String)
	}
	if p.Projectid.Valid {
		m["projectid"] = fmt.Sprint(p.Projectid.String)
	}
	if p.Size.Valid {
		m["size"] = fmt.Sprint(p.Size.Int64)
	}
	if p.Snapshotid.Valid {
		m["snapshotid"] = fmt.Sprint(p.Snapshotid.String)
	}
	if p.Virtualmachineid.Valid {
		m["virtualmachineid"] = fmt.Sprint(p.Virtualmachineid.String)
	}
	if p.Zoneid.Valid {
		m["zoneid"] = fmt.Sprint(p.Zoneid.String)
	}
	return m
}
func (c *Client) CreateVolume(p CreateVolumeParameter) (Volume, error) {
	var v map[string]json.RawMessage
	var ret Volume
	b, err := c.Request("createVolume", p.ToMap())
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
		errortext, ok := v["errortext"]
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
