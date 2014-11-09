package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type DeleteSecurityGroupParameter struct {
	Account   NullString
	Domainid  ID
	Id        ID
	Name      NullString
	Projectid ID
}

func (p *DeleteSecurityGroupParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *DeleteSecurityGroupParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *DeleteSecurityGroupParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *DeleteSecurityGroupParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *DeleteSecurityGroupParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *DeleteSecurityGroupParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Name.Valid {
		m["name"] = fmt.Sprint(p.Name.String)
	}
	if p.Projectid.Valid {
		m["projectid"] = fmt.Sprint(p.Projectid.String)
	}
	return m
}
func (c *Client) DeleteSecurityGroup(p DeleteSecurityGroupParameter) (Result, error) {
	var ret Result
	b, err := c.Request("deleteSecurityGroup", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	return ret, nil
}
