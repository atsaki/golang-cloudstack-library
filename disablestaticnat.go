package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type DisableStaticNatParameter struct {
	Ipaddressid ID
}

func (p *DisableStaticNatParameter) SetIpaddressid(s string) {
	p.Ipaddressid.String = s
	p.Ipaddressid.Valid = true
}
func (p *DisableStaticNatParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Ipaddressid.Valid {
		m["ipaddressid"] = fmt.Sprint(p.Ipaddressid.String)
	}
	return m
}
func (c *Client) DisableStaticNat(p DisableStaticNatParameter) (Result, error) {
	var ret Result
	b, err := c.Request("disableStaticNat", p.ToMap())
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
