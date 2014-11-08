package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type ListSSHKeyPairsParameter struct {
	Account     NullString
	Domainid    ID
	Fingerprint NullString
	Isrecursive NullBool
	Keyword     NullString
	Listall     NullBool
	Name        NullString
	Page        NullInt64
	Pagesize    NullInt64
	Projectid   ID
}

func (p *ListSSHKeyPairsParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *ListSSHKeyPairsParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *ListSSHKeyPairsParameter) SetFingerprint(s string) {
	p.Fingerprint.String = s
	p.Fingerprint.Valid = true
}
func (p *ListSSHKeyPairsParameter) SetIsrecursive(b bool) {
	p.Isrecursive.Bool = b
	p.Isrecursive.Valid = true
}
func (p *ListSSHKeyPairsParameter) SetKeyword(s string) {
	p.Keyword.String = s
	p.Keyword.Valid = true
}
func (p *ListSSHKeyPairsParameter) SetListall(b bool) {
	p.Listall.Bool = b
	p.Listall.Valid = true
}
func (p *ListSSHKeyPairsParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *ListSSHKeyPairsParameter) SetPage(n int64) {
	p.Page.Int64 = n
	p.Page.Valid = true
}
func (p *ListSSHKeyPairsParameter) SetPagesize(n int64) {
	p.Pagesize.Int64 = n
	p.Pagesize.Valid = true
}
func (p *ListSSHKeyPairsParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *ListSSHKeyPairsParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Fingerprint.Valid {
		m["fingerprint"] = fmt.Sprint(p.Fingerprint.String)
	}
	if p.Isrecursive.Valid {
		m["isrecursive"] = fmt.Sprint(p.Isrecursive.Bool)
	}
	if p.Keyword.Valid {
		m["keyword"] = fmt.Sprint(p.Keyword.String)
	}
	if p.Listall.Valid {
		m["listall"] = fmt.Sprint(p.Listall.Bool)
	}
	if p.Name.Valid {
		m["name"] = fmt.Sprint(p.Name.String)
	}
	if p.Page.Valid {
		m["page"] = fmt.Sprint(p.Page.Int64)
	}
	if p.Pagesize.Valid {
		m["pagesize"] = fmt.Sprint(p.Pagesize.Int64)
	}
	if p.Projectid.Valid {
		m["projectid"] = fmt.Sprint(p.Projectid.String)
	}
	return m
}
func (c *Client) ListSSHKeyPairs(p ListSSHKeyPairsParameter) ([]Sshkeypair, error) {
	var v map[string]json.RawMessage
	var ret []Sshkeypair
	b, err := c.Request("listSSHKeyPairs", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		log.Println("json.Unmarshal failed:", err)
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
