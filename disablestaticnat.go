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
	var v map[string]json.RawMessage
	var ret Result
	b, err := c.Request("disableStaticNat", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Println("failed:", err)
	}
	content, ok := v["result"]
	if !ok {
		log.Println("Content is empty.")
		return ret, nil
	}
	err = json.Unmarshal(content, &ret)
	if err != nil {
		log.Println("json.Unmarshal failed:", err)
	}
	return ret, err
}
