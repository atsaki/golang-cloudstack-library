package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type DestroyVirtualMachineParameter struct {
	Expunge NullBool
	Id      ID
}

func (p *DestroyVirtualMachineParameter) SetExpunge(b bool) {
	p.Expunge.Bool = b
	p.Expunge.Valid = true
}
func (p *DestroyVirtualMachineParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *DestroyVirtualMachineParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Expunge.Valid {
		m["expunge"] = fmt.Sprint(p.Expunge.Bool)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	return m
}
func (c *Client) DestroyVirtualMachine(p DestroyVirtualMachineParameter) (Virtualmachine, error) {
	var v map[string]json.RawMessage
	var ret Virtualmachine
	b, err := c.Request("destroyVirtualMachine", p.ToMap())
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
		errortext, _ := v["errortext"]
		return ret, fmt.Errorf(string(errortext))
	}
	err = json.Unmarshal(content, &ret)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(content))
	}
	return ret, nil
}
