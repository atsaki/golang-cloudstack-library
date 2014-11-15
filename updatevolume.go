package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type UpdateVolumeParameter struct {
	Chaininfo     NullString
	Customid      NullString
	Displayvolume NullBool
	Id            ID
	Path          NullString
	State         NullString
	Storageid     ID
}

func (p *UpdateVolumeParameter) SetChaininfo(s string) {
	p.Chaininfo.String = s
	p.Chaininfo.Valid = true
}
func (p *UpdateVolumeParameter) SetCustomid(s string) {
	p.Customid.String = s
	p.Customid.Valid = true
}
func (p *UpdateVolumeParameter) SetDisplayvolume(b bool) {
	p.Displayvolume.Bool = b
	p.Displayvolume.Valid = true
}
func (p *UpdateVolumeParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *UpdateVolumeParameter) SetPath(s string) {
	p.Path.String = s
	p.Path.Valid = true
}
func (p *UpdateVolumeParameter) SetState(s string) {
	p.State.String = s
	p.State.Valid = true
}
func (p *UpdateVolumeParameter) SetStorageid(s string) {
	p.Storageid.String = s
	p.Storageid.Valid = true
}
func (p *UpdateVolumeParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Chaininfo.Valid {
		m["chaininfo"] = fmt.Sprint(p.Chaininfo.String)
	}
	if p.Customid.Valid {
		m["customid"] = fmt.Sprint(p.Customid.String)
	}
	if p.Displayvolume.Valid {
		m["displayvolume"] = fmt.Sprint(p.Displayvolume.Bool)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Path.Valid {
		m["path"] = fmt.Sprint(p.Path.String)
	}
	if p.State.Valid {
		m["state"] = fmt.Sprint(p.State.String)
	}
	if p.Storageid.Valid {
		m["storageid"] = fmt.Sprint(p.Storageid.String)
	}
	return m
}
func (c *Client) UpdateVolume(p UpdateVolumeParameter) (Volume, error) {
	var v map[string]json.RawMessage
	var ret Volume
	b, err := c.Request("updateVolume", p.ToMap())
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
