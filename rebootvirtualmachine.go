package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type RebootVirtualMachineParameter struct {
	Id ID
}

func (p *RebootVirtualMachineParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *RebootVirtualMachineParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	return m
}
func (c *Client) RebootVirtualMachine(p RebootVirtualMachineParameter) (Virtualmachine, error) {
	var v map[string]json.RawMessage
	var ret Virtualmachine
	b, err := c.Request("rebootVirtualMachine", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Println("failed:", err)
	}
	content, ok := v["virtualmachine"]
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
