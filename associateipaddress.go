package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type AssociateIpAddressParameter struct {
	Account    NullString
	Domainid   ID
	Fordisplay NullBool
	Isportable NullBool
	Networkid  ID
	Projectid  ID
	Regionid   NullInt64
	Vpcid      ID
	Zoneid     ID
}

func (p *AssociateIpAddressParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *AssociateIpAddressParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *AssociateIpAddressParameter) SetFordisplay(b bool) {
	p.Fordisplay.Bool = b
	p.Fordisplay.Valid = true
}
func (p *AssociateIpAddressParameter) SetIsportable(b bool) {
	p.Isportable.Bool = b
	p.Isportable.Valid = true
}
func (p *AssociateIpAddressParameter) SetNetworkid(s string) {
	p.Networkid.String = s
	p.Networkid.Valid = true
}
func (p *AssociateIpAddressParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *AssociateIpAddressParameter) SetRegionid(n int64) {
	p.Regionid.Int64 = n
	p.Regionid.Valid = true
}
func (p *AssociateIpAddressParameter) SetVpcid(s string) {
	p.Vpcid.String = s
	p.Vpcid.Valid = true
}
func (p *AssociateIpAddressParameter) SetZoneid(s string) {
	p.Zoneid.String = s
	p.Zoneid.Valid = true
}
func (p *AssociateIpAddressParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Fordisplay.Valid {
		m["fordisplay"] = fmt.Sprint(p.Fordisplay.Bool)
	}
	if p.Isportable.Valid {
		m["isportable"] = fmt.Sprint(p.Isportable.Bool)
	}
	if p.Networkid.Valid {
		m["networkid"] = fmt.Sprint(p.Networkid.String)
	}
	if p.Projectid.Valid {
		m["projectid"] = fmt.Sprint(p.Projectid.String)
	}
	if p.Regionid.Valid {
		m["regionid"] = fmt.Sprint(p.Regionid.Int64)
	}
	if p.Vpcid.Valid {
		m["vpcid"] = fmt.Sprint(p.Vpcid.String)
	}
	if p.Zoneid.Valid {
		m["zoneid"] = fmt.Sprint(p.Zoneid.String)
	}
	return m
}
func (c *Client) AssociateIpAddress(p AssociateIpAddressParameter) (Publicipaddress, error) {
	var v map[string]json.RawMessage
	var ret Publicipaddress
	b, err := c.Request("associateIpAddress", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["ipaddress"]
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
