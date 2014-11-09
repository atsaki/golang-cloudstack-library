package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type DisassociateIpAddressParameter struct {
	Id ID
}

func (p *DisassociateIpAddressParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *DisassociateIpAddressParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	return m
}
func (c *Client) DisassociateIpAddress(p DisassociateIpAddressParameter) (Result, error) {
	var v map[string]json.RawMessage
	var ret Result
	b, err := c.Request("disassociateIpAddress", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["result"]
	if !ok {
		errortext, _ := v["errortext"]
		return ret, fmt.Errorf(string(errortext))
	}
	err = json.Unmarshal(content, &ret)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal: %s", string(content))
	}
	return ret, nil
}
