package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type DetachVolumeParameter struct {
	Deviceid         NullInt64
	Id               ID
	Virtualmachineid ID
}

func (p *DetachVolumeParameter) SetDeviceid(n int64) {
	p.Deviceid.Int64 = n
	p.Deviceid.Valid = true
}
func (p *DetachVolumeParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *DetachVolumeParameter) SetVirtualmachineid(s string) {
	p.Virtualmachineid.String = s
	p.Virtualmachineid.Valid = true
}
func (p *DetachVolumeParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Deviceid.Valid {
		m["deviceid"] = fmt.Sprint(p.Deviceid.Int64)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Virtualmachineid.Valid {
		m["virtualmachineid"] = fmt.Sprint(p.Virtualmachineid.String)
	}
	return m
}
func (c *Client) DetachVolume(p DetachVolumeParameter) (Volume, error) {
	var v map[string]json.RawMessage
	var ret Volume
	b, err := c.Request("detachVolume", p.ToMap())
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
