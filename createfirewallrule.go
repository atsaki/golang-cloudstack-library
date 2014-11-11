package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type CreateFirewallRuleParameter struct {
	Cidrlist    []string
	Endport     NullInt64
	Fordisplay  NullBool
	Icmpcode    NullInt64
	Icmptype    NullInt64
	Ipaddressid ID
	Protocol    NullString
	Startport   NullInt64
	Type        NullString
}

func (p *CreateFirewallRuleParameter) SetCidrlist(l []string) {
	p.Cidrlist = l
}
func (p *CreateFirewallRuleParameter) SetEndport(n int64) {
	p.Endport.Int64 = n
	p.Endport.Valid = true
}
func (p *CreateFirewallRuleParameter) SetFordisplay(b bool) {
	p.Fordisplay.Bool = b
	p.Fordisplay.Valid = true
}
func (p *CreateFirewallRuleParameter) SetIcmpcode(n int64) {
	p.Icmpcode.Int64 = n
	p.Icmpcode.Valid = true
}
func (p *CreateFirewallRuleParameter) SetIcmptype(n int64) {
	p.Icmptype.Int64 = n
	p.Icmptype.Valid = true
}
func (p *CreateFirewallRuleParameter) SetIpaddressid(s string) {
	p.Ipaddressid.String = s
	p.Ipaddressid.Valid = true
}
func (p *CreateFirewallRuleParameter) SetProtocol(s string) {
	p.Protocol.String = s
	p.Protocol.Valid = true
}
func (p *CreateFirewallRuleParameter) SetStartport(n int64) {
	p.Startport.Int64 = n
	p.Startport.Valid = true
}
func (p *CreateFirewallRuleParameter) SetType(s string) {
	p.Type.String = s
	p.Type.Valid = true
}
func (p *CreateFirewallRuleParameter) ToMap() map[string]string {
	m := map[string]string{}
	if len(p.Cidrlist) > 0 {
		m["cidrlist"] = strings.Join(p.Cidrlist, ",")
	}
	if p.Endport.Valid {
		m["endport"] = fmt.Sprint(p.Endport.Int64)
	}
	if p.Fordisplay.Valid {
		m["fordisplay"] = fmt.Sprint(p.Fordisplay.Bool)
	}
	if p.Icmpcode.Valid {
		m["icmpcode"] = fmt.Sprint(p.Icmpcode.Int64)
	}
	if p.Icmptype.Valid {
		m["icmptype"] = fmt.Sprint(p.Icmptype.Int64)
	}
	if p.Ipaddressid.Valid {
		m["ipaddressid"] = fmt.Sprint(p.Ipaddressid.String)
	}
	if p.Protocol.Valid {
		m["protocol"] = fmt.Sprint(p.Protocol.String)
	}
	if p.Startport.Valid {
		m["startport"] = fmt.Sprint(p.Startport.Int64)
	}
	if p.Type.Valid {
		m["type"] = fmt.Sprint(p.Type.String)
	}
	return m
}
func (c *Client) CreateFirewallRule(p CreateFirewallRuleParameter) (Firewallrule, error) {
	var v map[string]json.RawMessage
	var ret Firewallrule
	b, err := c.Request("createFirewallRule", p.ToMap())
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
