package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type CreatePortForwardingRuleParameter struct {
	Cidrlist         []string
	Fordisplay       NullBool
	Ipaddressid      ID
	Networkid        ID
	Openfirewall     NullBool
	Privateendport   NullInt64
	Privateport      NullInt64
	Protocol         NullString
	Publicendport    NullInt64
	Publicport       NullInt64
	Virtualmachineid ID
	Vmguestip        NullString
}

func (p *CreatePortForwardingRuleParameter) SetCidrlist(l []string) {
	p.Cidrlist = l
}
func (p *CreatePortForwardingRuleParameter) SetFordisplay(b bool) {
	p.Fordisplay.Bool = b
	p.Fordisplay.Valid = true
}
func (p *CreatePortForwardingRuleParameter) SetIpaddressid(s string) {
	p.Ipaddressid.String = s
	p.Ipaddressid.Valid = true
}
func (p *CreatePortForwardingRuleParameter) SetNetworkid(s string) {
	p.Networkid.String = s
	p.Networkid.Valid = true
}
func (p *CreatePortForwardingRuleParameter) SetOpenfirewall(b bool) {
	p.Openfirewall.Bool = b
	p.Openfirewall.Valid = true
}
func (p *CreatePortForwardingRuleParameter) SetPrivateendport(n int64) {
	p.Privateendport.Int64 = n
	p.Privateendport.Valid = true
}
func (p *CreatePortForwardingRuleParameter) SetPrivateport(n int64) {
	p.Privateport.Int64 = n
	p.Privateport.Valid = true
}
func (p *CreatePortForwardingRuleParameter) SetProtocol(s string) {
	p.Protocol.String = s
	p.Protocol.Valid = true
}
func (p *CreatePortForwardingRuleParameter) SetPublicendport(n int64) {
	p.Publicendport.Int64 = n
	p.Publicendport.Valid = true
}
func (p *CreatePortForwardingRuleParameter) SetPublicport(n int64) {
	p.Publicport.Int64 = n
	p.Publicport.Valid = true
}
func (p *CreatePortForwardingRuleParameter) SetVirtualmachineid(s string) {
	p.Virtualmachineid.String = s
	p.Virtualmachineid.Valid = true
}
func (p *CreatePortForwardingRuleParameter) SetVmguestip(s string) {
	p.Vmguestip.String = s
	p.Vmguestip.Valid = true
}
func (p *CreatePortForwardingRuleParameter) ToMap() map[string]string {
	m := map[string]string{}
	if len(p.Cidrlist) > 0 {
		m["cidrlist"] = strings.Join(p.Cidrlist, ",")
	}
	if p.Fordisplay.Valid {
		m["fordisplay"] = fmt.Sprint(p.Fordisplay.Bool)
	}
	if p.Ipaddressid.Valid {
		m["ipaddressid"] = fmt.Sprint(p.Ipaddressid.String)
	}
	if p.Networkid.Valid {
		m["networkid"] = fmt.Sprint(p.Networkid.String)
	}
	if p.Openfirewall.Valid {
		m["openfirewall"] = fmt.Sprint(p.Openfirewall.Bool)
	}
	if p.Privateendport.Valid {
		m["privateendport"] = fmt.Sprint(p.Privateendport.Int64)
	}
	if p.Privateport.Valid {
		m["privateport"] = fmt.Sprint(p.Privateport.Int64)
	}
	if p.Protocol.Valid {
		m["protocol"] = fmt.Sprint(p.Protocol.String)
	}
	if p.Publicendport.Valid {
		m["publicendport"] = fmt.Sprint(p.Publicendport.Int64)
	}
	if p.Publicport.Valid {
		m["publicport"] = fmt.Sprint(p.Publicport.Int64)
	}
	if p.Virtualmachineid.Valid {
		m["virtualmachineid"] = fmt.Sprint(p.Virtualmachineid.String)
	}
	if p.Vmguestip.Valid {
		m["vmguestip"] = fmt.Sprint(p.Vmguestip.String)
	}
	return m
}
func (c *Client) CreatePortForwardingRule(p CreatePortForwardingRuleParameter) (Portforwardingrule, error) {
	var v map[string]json.RawMessage
	var ret Portforwardingrule
	b, err := c.Request("createPortForwardingRule", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["portforwardingrule"]
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
