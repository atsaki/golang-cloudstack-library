package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type StartVirtualMachineParameter struct {
	Deploymentplanner NullString
	Hostid            ID
	Id                ID
}

func (p *StartVirtualMachineParameter) SetDeploymentplanner(s string) {
	p.Deploymentplanner.String = s
	p.Deploymentplanner.Valid = true
}
func (p *StartVirtualMachineParameter) SetHostid(s string) {
	p.Hostid.String = s
	p.Hostid.Valid = true
}
func (p *StartVirtualMachineParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *StartVirtualMachineParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Deploymentplanner.Valid {
		m["deploymentplanner"] = fmt.Sprint(p.Deploymentplanner.String)
	}
	if p.Hostid.Valid {
		m["hostid"] = fmt.Sprint(p.Hostid.String)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	return m
}
func (c *Client) StartVirtualMachine(p StartVirtualMachineParameter) (Virtualmachine, error) {
	var v map[string]json.RawMessage
	var ret Virtualmachine
	b, err := c.Request("startVirtualMachine", p.ToMap())
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
