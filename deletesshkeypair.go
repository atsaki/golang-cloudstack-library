package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type DeleteSSHKeyPairParameter struct {
	Account   NullString
	Domainid  ID
	Name      NullString
	Projectid ID
}

func (p *DeleteSSHKeyPairParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *DeleteSSHKeyPairParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *DeleteSSHKeyPairParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *DeleteSSHKeyPairParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *DeleteSSHKeyPairParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Name.Valid {
		m["name"] = fmt.Sprint(p.Name.String)
	}
	if p.Projectid.Valid {
		m["projectid"] = fmt.Sprint(p.Projectid.String)
	}
	return m
}
func (c *Client) DeleteSSHKeyPair(p DeleteSSHKeyPairParameter) (Result, error) {
	var v map[string]json.RawMessage
	var ret Result
	b, err := c.Request("deleteSSHKeyPair", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["result"]
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
