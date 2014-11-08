package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type UpdateVirtualMachineParameter struct {
	Customid              NullString
	Displayname           NullString
	Displayvm             NullBool
	Group                 NullString
	Haenable              NullBool
	Id                    ID
	Isdynamicallyscalable NullBool
	Name                  NullString
	Ostypeid              ID
	Userdata              NullString
}

func (p *UpdateVirtualMachineParameter) SetCustomid(s string) {
	p.Customid.String = s
	p.Customid.Valid = true
}
func (p *UpdateVirtualMachineParameter) SetDisplayname(s string) {
	p.Displayname.String = s
	p.Displayname.Valid = true
}
func (p *UpdateVirtualMachineParameter) SetDisplayvm(b bool) {
	p.Displayvm.Bool = b
	p.Displayvm.Valid = true
}
func (p *UpdateVirtualMachineParameter) SetGroup(s string) {
	p.Group.String = s
	p.Group.Valid = true
}
func (p *UpdateVirtualMachineParameter) SetHaenable(b bool) {
	p.Haenable.Bool = b
	p.Haenable.Valid = true
}
func (p *UpdateVirtualMachineParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *UpdateVirtualMachineParameter) SetIsdynamicallyscalable(b bool) {
	p.Isdynamicallyscalable.Bool = b
	p.Isdynamicallyscalable.Valid = true
}
func (p *UpdateVirtualMachineParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *UpdateVirtualMachineParameter) SetOstypeid(s string) {
	p.Ostypeid.String = s
	p.Ostypeid.Valid = true
}
func (p *UpdateVirtualMachineParameter) SetUserdata(s string) {
	p.Userdata.String = s
	p.Userdata.Valid = true
}
func (p *UpdateVirtualMachineParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Customid.Valid {
		m["customid"] = fmt.Sprint(p.Customid.String)
	}
	if p.Displayname.Valid {
		m["displayname"] = fmt.Sprint(p.Displayname.String)
	}
	if p.Displayvm.Valid {
		m["displayvm"] = fmt.Sprint(p.Displayvm.Bool)
	}
	if p.Group.Valid {
		m["group"] = fmt.Sprint(p.Group.String)
	}
	if p.Haenable.Valid {
		m["haenable"] = fmt.Sprint(p.Haenable.Bool)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Isdynamicallyscalable.Valid {
		m["isdynamicallyscalable"] = fmt.Sprint(p.Isdynamicallyscalable.Bool)
	}
	if p.Name.Valid {
		m["name"] = fmt.Sprint(p.Name.String)
	}
	if p.Ostypeid.Valid {
		m["ostypeid"] = fmt.Sprint(p.Ostypeid.String)
	}
	if p.Userdata.Valid {
		m["userdata"] = fmt.Sprint(p.Userdata.String)
	}
	return m
}
func (c *Client) UpdateVirtualMachine(p UpdateVirtualMachineParameter) (Virtualmachine, error) {
	var v map[string]json.RawMessage
	var ret Virtualmachine
	b, err := c.Request("updateVirtualMachine", p.ToMap())
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
