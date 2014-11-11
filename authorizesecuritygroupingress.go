package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type AuthorizeSecurityGroupIngressParameter struct {
	Account               NullString
	Cidrlist              []string
	Domainid              ID
	Endport               NullInt64
	Icmpcode              NullInt64
	Icmptype              NullInt64
	Projectid             ID
	Protocol              NullString
	Securitygroupid       ID
	Securitygroupname     NullString
	Startport             NullInt64
	Usersecuritygrouplist map[string]string
}

func (p *AuthorizeSecurityGroupIngressParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *AuthorizeSecurityGroupIngressParameter) SetCidrlist(l []string) {
	p.Cidrlist = l
}
func (p *AuthorizeSecurityGroupIngressParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *AuthorizeSecurityGroupIngressParameter) SetEndport(n int64) {
	p.Endport.Int64 = n
	p.Endport.Valid = true
}
func (p *AuthorizeSecurityGroupIngressParameter) SetIcmpcode(n int64) {
	p.Icmpcode.Int64 = n
	p.Icmpcode.Valid = true
}
func (p *AuthorizeSecurityGroupIngressParameter) SetIcmptype(n int64) {
	p.Icmptype.Int64 = n
	p.Icmptype.Valid = true
}
func (p *AuthorizeSecurityGroupIngressParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *AuthorizeSecurityGroupIngressParameter) SetProtocol(s string) {
	p.Protocol.String = s
	p.Protocol.Valid = true
}
func (p *AuthorizeSecurityGroupIngressParameter) SetSecuritygroupid(s string) {
	p.Securitygroupid.String = s
	p.Securitygroupid.Valid = true
}
func (p *AuthorizeSecurityGroupIngressParameter) SetSecuritygroupname(s string) {
	p.Securitygroupname.String = s
	p.Securitygroupname.Valid = true
}
func (p *AuthorizeSecurityGroupIngressParameter) SetStartport(n int64) {
	p.Startport.Int64 = n
	p.Startport.Valid = true
}
func (p *AuthorizeSecurityGroupIngressParameter) SetUsersecuritygrouplist(m map[string]string) {
	p.Usersecuritygrouplist = m
}
func (p *AuthorizeSecurityGroupIngressParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if len(p.Cidrlist) > 0 {
		m["cidrlist"] = strings.Join(p.Cidrlist, ",")
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Endport.Valid {
		m["endport"] = fmt.Sprint(p.Endport.Int64)
	}
	if p.Icmpcode.Valid {
		m["icmpcode"] = fmt.Sprint(p.Icmpcode.Int64)
	}
	if p.Icmptype.Valid {
		m["icmptype"] = fmt.Sprint(p.Icmptype.Int64)
	}
	if p.Projectid.Valid {
		m["projectid"] = fmt.Sprint(p.Projectid.String)
	}
	if p.Protocol.Valid {
		m["protocol"] = fmt.Sprint(p.Protocol.String)
	}
	if p.Securitygroupid.Valid {
		m["securitygroupid"] = fmt.Sprint(p.Securitygroupid.String)
	}
	if p.Securitygroupname.Valid {
		m["securitygroupname"] = fmt.Sprint(p.Securitygroupname.String)
	}
	if p.Startport.Valid {
		m["startport"] = fmt.Sprint(p.Startport.Int64)
	}
	if len(p.Usersecuritygrouplist) > 0 {
		i := 0
		for key, value := range m {
			m[fmt.Sprintf("usersecuritygrouplist[%d].key", i)] = key
			m[fmt.Sprintf("usersecuritygrouplist[%d].value", i)] = value
			i += 1
		}
	}
	return m
}
func (c *Client) AuthorizeSecurityGroupIngress(p AuthorizeSecurityGroupIngressParameter) (Securitygroupingress, error) {
	var v map[string]json.RawMessage
	var ret Securitygroupingress
	b, err := c.Request("authorizeSecurityGroupIngress", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["securitygroup"]
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
