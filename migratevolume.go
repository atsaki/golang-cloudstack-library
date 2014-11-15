package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type MigrateVolumeParameter struct {
	Livemigrate NullBool
	Storageid   ID
	Volumeid    ID
}

func (p *MigrateVolumeParameter) SetLivemigrate(b bool) {
	p.Livemigrate.Bool = b
	p.Livemigrate.Valid = true
}
func (p *MigrateVolumeParameter) SetStorageid(s string) {
	p.Storageid.String = s
	p.Storageid.Valid = true
}
func (p *MigrateVolumeParameter) SetVolumeid(s string) {
	p.Volumeid.String = s
	p.Volumeid.Valid = true
}
func (p *MigrateVolumeParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Livemigrate.Valid {
		m["livemigrate"] = fmt.Sprint(p.Livemigrate.Bool)
	}
	if p.Storageid.Valid {
		m["storageid"] = fmt.Sprint(p.Storageid.String)
	}
	if p.Volumeid.Valid {
		m["volumeid"] = fmt.Sprint(p.Volumeid.String)
	}
	return m
}
func (c *Client) MigrateVolume(p MigrateVolumeParameter) (Volume, error) {
	var v map[string]json.RawMessage
	var ret Volume
	b, err := c.Request("migrateVolume", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return ret, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return ret, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}
	content, ok := v["volume"]
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
