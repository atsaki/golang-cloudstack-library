package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type RegisterSSHKeyPairParameter struct {
	Account   NullString
	Domainid  ID
	Name      NullString
	Projectid ID
	Publickey NullString
}

func (p *RegisterSSHKeyPairParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *RegisterSSHKeyPairParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *RegisterSSHKeyPairParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *RegisterSSHKeyPairParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *RegisterSSHKeyPairParameter) SetPublickey(s string) {
	p.Publickey.String = s
	p.Publickey.Valid = true
}
func (p *RegisterSSHKeyPairParameter) ToMap() map[string]string {
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
	if p.Publickey.Valid {
		m["publickey"] = fmt.Sprint(p.Publickey.String)
	}
	return m
}
func (c *Client) RegisterSSHKeyPair(p RegisterSSHKeyPairParameter) (Sshkeypair, error) {
	var v map[string]json.RawMessage
	var ret Sshkeypair
	b, err := c.Request("registerSSHKeyPair", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Println("failed:", err)
	}
	content, ok := v["sshkeypair"]
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
