package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type DeleteFirewallRuleParameter struct {
	Id ID
}

func (p *DeleteFirewallRuleParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *DeleteFirewallRuleParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	return m
}
func (c *Client) DeleteFirewallRule(p DeleteFirewallRuleParameter) (Result, error) {
	var ret Result
	b, err := c.Request("deleteFirewallRule", p.ToMap())
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
