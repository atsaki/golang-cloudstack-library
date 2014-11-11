package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type RemoveNicFromVirtualMachineParameter struct {
	Nicid            ID
	Virtualmachineid ID
}

func (p *RemoveNicFromVirtualMachineParameter) SetNicid(s string) {
	p.Nicid.String = s
	p.Nicid.Valid = true
}
func (p *RemoveNicFromVirtualMachineParameter) SetVirtualmachineid(s string) {
	p.Virtualmachineid.String = s
	p.Virtualmachineid.Valid = true
}
func (p *RemoveNicFromVirtualMachineParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Nicid.Valid {
		m["nicid"] = fmt.Sprint(p.Nicid.String)
	}
	if p.Virtualmachineid.Valid {
		m["virtualmachineid"] = fmt.Sprint(p.Virtualmachineid.String)
	}
	return m
}
func (c *Client) RemoveNicFromVirtualMachine(p RemoveNicFromVirtualMachineParameter) (Virtualmachine, error) {
	var v map[string]json.RawMessage
	var ret Virtualmachine
	b, err := c.Request("removeNicFromVirtualMachine", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["virtualmachine"]
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
