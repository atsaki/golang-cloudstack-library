package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type ResizeVolumeParameter struct {
	Diskofferingid ID
	Id             ID
	Shrinkok       NullBool
	Size           NullInt64
}

func (p *ResizeVolumeParameter) SetDiskofferingid(s string) {
	p.Diskofferingid.String = s
	p.Diskofferingid.Valid = true
}
func (p *ResizeVolumeParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *ResizeVolumeParameter) SetShrinkok(b bool) {
	p.Shrinkok.Bool = b
	p.Shrinkok.Valid = true
}
func (p *ResizeVolumeParameter) SetSize(n int64) {
	p.Size.Int64 = n
	p.Size.Valid = true
}
func (p *ResizeVolumeParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Diskofferingid.Valid {
		m["diskofferingid"] = fmt.Sprint(p.Diskofferingid.String)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Shrinkok.Valid {
		m["shrinkok"] = fmt.Sprint(p.Shrinkok.Bool)
	}
	if p.Size.Valid {
		m["size"] = fmt.Sprint(p.Size.Int64)
	}
	return m
}
func (c *Client) ResizeVolume(p ResizeVolumeParameter) (Volume, error) {
	var v map[string]json.RawMessage
	var ret Volume
	b, err := c.Request("resizeVolume", p.ToMap())
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
