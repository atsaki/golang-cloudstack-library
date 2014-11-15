package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type DeleteVolumeParameter struct {
	Id ID
}

func (p *DeleteVolumeParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *DeleteVolumeParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	return m
}
func (c *Client) DeleteVolume(p DeleteVolumeParameter) (Result, error) {
	var ret Result
	b, err := c.Request("deleteVolume", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	return ret, nil
}
