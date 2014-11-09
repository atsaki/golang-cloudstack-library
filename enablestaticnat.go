package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type EnableStaticNatParameter struct {
	Ipaddressid      ID
	Networkid        ID
	Virtualmachineid ID
	Vmguestip        NullString
}

func (p *EnableStaticNatParameter) SetIpaddressid(s string) {
	p.Ipaddressid.String = s
	p.Ipaddressid.Valid = true
}
func (p *EnableStaticNatParameter) SetNetworkid(s string) {
	p.Networkid.String = s
	p.Networkid.Valid = true
}
func (p *EnableStaticNatParameter) SetVirtualmachineid(s string) {
	p.Virtualmachineid.String = s
	p.Virtualmachineid.Valid = true
}
func (p *EnableStaticNatParameter) SetVmguestip(s string) {
	p.Vmguestip.String = s
	p.Vmguestip.Valid = true
}
func (p *EnableStaticNatParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Ipaddressid.Valid {
		m["ipaddressid"] = fmt.Sprint(p.Ipaddressid.String)
	}
	if p.Networkid.Valid {
		m["networkid"] = fmt.Sprint(p.Networkid.String)
	}
	if p.Virtualmachineid.Valid {
		m["virtualmachineid"] = fmt.Sprint(p.Virtualmachineid.String)
	}
	if p.Vmguestip.Valid {
		m["vmguestip"] = fmt.Sprint(p.Vmguestip.String)
	}
	return m
}
func (c *Client) EnableStaticNat(p EnableStaticNatParameter) (Result, error) {
	var v map[string]json.RawMessage
	var ret Result
	b, err := c.Request("enableStaticNat", p.ToMap())
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
