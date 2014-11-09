package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type AddNicToVirtualMachineParameter struct {
	Ipaddress        NullString
	Networkid        ID
	Virtualmachineid ID
}

func (p *AddNicToVirtualMachineParameter) SetIpaddress(s string) {
	p.Ipaddress.String = s
	p.Ipaddress.Valid = true
}
func (p *AddNicToVirtualMachineParameter) SetNetworkid(s string) {
	p.Networkid.String = s
	p.Networkid.Valid = true
}
func (p *AddNicToVirtualMachineParameter) SetVirtualmachineid(s string) {
	p.Virtualmachineid.String = s
	p.Virtualmachineid.Valid = true
}
func (p *AddNicToVirtualMachineParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Ipaddress.Valid {
		m["ipaddress"] = fmt.Sprint(p.Ipaddress.String)
	}
	if p.Networkid.Valid {
		m["networkid"] = fmt.Sprint(p.Networkid.String)
	}
	if p.Virtualmachineid.Valid {
		m["virtualmachineid"] = fmt.Sprint(p.Virtualmachineid.String)
	}
	return m
}
func (c *Client) AddNicToVirtualMachine(p AddNicToVirtualMachineParameter) (Virtualmachine, error) {
	var v map[string]json.RawMessage
	var ret Virtualmachine
	b, err := c.Request("addNicToVirtualMachine", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["virtualmachine"]
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
