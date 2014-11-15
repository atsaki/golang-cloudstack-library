package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type MigrateVirtualMachineWithVolumeParameter struct {
	Hostid           ID
	Migrateto        map[string]string
	Virtualmachineid ID
}

func (p *MigrateVirtualMachineWithVolumeParameter) SetHostid(s string) {
	p.Hostid.String = s
	p.Hostid.Valid = true
}
func (p *MigrateVirtualMachineWithVolumeParameter) SetMigrateto(m map[string]string) {
	p.Migrateto = m
}
func (p *MigrateVirtualMachineWithVolumeParameter) SetVirtualmachineid(s string) {
	p.Virtualmachineid.String = s
	p.Virtualmachineid.Valid = true
}
func (p *MigrateVirtualMachineWithVolumeParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Hostid.Valid {
		m["hostid"] = fmt.Sprint(p.Hostid.String)
	}
	if len(p.Migrateto) > 0 {
		i := 0
		for key, value := range m {
			m[fmt.Sprintf("migrateto[%d].key", i)] = key
			m[fmt.Sprintf("migrateto[%d].value", i)] = value
			i += 1
		}
	}
	if p.Virtualmachineid.Valid {
		m["virtualmachineid"] = fmt.Sprint(p.Virtualmachineid.String)
	}
	return m
}
func (c *Client) MigrateVirtualMachineWithVolume(p MigrateVirtualMachineWithVolumeParameter) (Virtualmachine, error) {
	var v map[string]json.RawMessage
	var ret Virtualmachine
	b, err := c.Request("migrateVirtualMachineWithVolume", p.ToMap())
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
