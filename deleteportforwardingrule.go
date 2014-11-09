package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type DeletePortForwardingRuleParameter struct {
	Id ID
}

func (p *DeletePortForwardingRuleParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *DeletePortForwardingRuleParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	return m
}
func (c *Client) DeletePortForwardingRule(p DeletePortForwardingRuleParameter) (Result, error) {
	var v map[string]json.RawMessage
	var ret Result
	b, err := c.Request("deletePortForwardingRule", p.ToMap())
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
