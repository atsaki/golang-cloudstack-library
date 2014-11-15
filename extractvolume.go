package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type ExtractVolumeParameter struct {
	Id     ID
	Mode   NullString
	Url    NullString
	Zoneid ID
}

func (p *ExtractVolumeParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *ExtractVolumeParameter) SetMode(s string) {
	p.Mode.String = s
	p.Mode.Valid = true
}
func (p *ExtractVolumeParameter) SetUrl(s string) {
	p.Url.String = s
	p.Url.Valid = true
}
func (p *ExtractVolumeParameter) SetZoneid(s string) {
	p.Zoneid.String = s
	p.Zoneid.Valid = true
}
func (p *ExtractVolumeParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Mode.Valid {
		m["mode"] = fmt.Sprint(p.Mode.String)
	}
	if p.Url.Valid {
		m["url"] = fmt.Sprint(p.Url.String)
	}
	if p.Zoneid.Valid {
		m["zoneid"] = fmt.Sprint(p.Zoneid.String)
	}
	return m
}
func (c *Client) ExtractVolume(p ExtractVolumeParameter) (Volume, error) {
	var v map[string]json.RawMessage
	var ret Volume
	b, err := c.Request("extractVolume", p.ToMap())
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
