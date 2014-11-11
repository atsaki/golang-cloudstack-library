package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type ListFirewallRulesParameter struct {
	Account     NullString
	Domainid    ID
	Fordisplay  NullBool
	Id          ID
	Ipaddressid ID
	Isrecursive NullBool
	Keyword     NullString
	Listall     NullBool
	Networkid   ID
	Page        NullInt64
	Pagesize    NullInt64
	Projectid   ID
	Tags        map[string]string
}

func (p *ListFirewallRulesParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *ListFirewallRulesParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *ListFirewallRulesParameter) SetFordisplay(b bool) {
	p.Fordisplay.Bool = b
	p.Fordisplay.Valid = true
}
func (p *ListFirewallRulesParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *ListFirewallRulesParameter) SetIpaddressid(s string) {
	p.Ipaddressid.String = s
	p.Ipaddressid.Valid = true
}
func (p *ListFirewallRulesParameter) SetIsrecursive(b bool) {
	p.Isrecursive.Bool = b
	p.Isrecursive.Valid = true
}
func (p *ListFirewallRulesParameter) SetKeyword(s string) {
	p.Keyword.String = s
	p.Keyword.Valid = true
}
func (p *ListFirewallRulesParameter) SetListall(b bool) {
	p.Listall.Bool = b
	p.Listall.Valid = true
}
func (p *ListFirewallRulesParameter) SetNetworkid(s string) {
	p.Networkid.String = s
	p.Networkid.Valid = true
}
func (p *ListFirewallRulesParameter) SetPage(n int64) {
	p.Page.Int64 = n
	p.Page.Valid = true
}
func (p *ListFirewallRulesParameter) SetPagesize(n int64) {
	p.Pagesize.Int64 = n
	p.Pagesize.Valid = true
}
func (p *ListFirewallRulesParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *ListFirewallRulesParameter) SetTags(m map[string]string) {
	p.Tags = m
}
func (p *ListFirewallRulesParameter) ToMap() map[string]string {
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
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Ipaddressid.Valid {
		m["ipaddressid"] = fmt.Sprint(p.Ipaddressid.String)
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
	if p.Networkid.Valid {
		m["networkid"] = fmt.Sprint(p.Networkid.String)
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
	if len(p.Tags) > 0 {
		i := 0
		for key, value := range m {
			m[fmt.Sprintf("tags[%d].key", i)] = key
			m[fmt.Sprintf("tags[%d].value", i)] = value
			i += 1
		}
	}
	return m
}
func (c *Client) ListFirewallRules(p ListFirewallRulesParameter) ([]Firewallrule, error) {
	var v map[string]json.RawMessage
	var ret []Firewallrule
	b, err := c.Request("listFirewallRules", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["firewallrule"]
	if !ok {
		if len(v) == 0 {
			return ret, nil
		}
		errortext, _ := v["errortext"]
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
