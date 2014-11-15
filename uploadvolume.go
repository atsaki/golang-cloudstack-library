package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type UploadVolumeParameter struct {
	Account        NullString
	Checksum       NullString
	Domainid       ID
	Format         NullString
	Imagestoreuuid NullString
	Name           NullString
	Projectid      ID
	Url            NullString
	Zoneid         ID
}

func (p *UploadVolumeParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *UploadVolumeParameter) SetChecksum(s string) {
	p.Checksum.String = s
	p.Checksum.Valid = true
}
func (p *UploadVolumeParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *UploadVolumeParameter) SetFormat(s string) {
	p.Format.String = s
	p.Format.Valid = true
}
func (p *UploadVolumeParameter) SetImagestoreuuid(s string) {
	p.Imagestoreuuid.String = s
	p.Imagestoreuuid.Valid = true
}
func (p *UploadVolumeParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *UploadVolumeParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *UploadVolumeParameter) SetUrl(s string) {
	p.Url.String = s
	p.Url.Valid = true
}
func (p *UploadVolumeParameter) SetZoneid(s string) {
	p.Zoneid.String = s
	p.Zoneid.Valid = true
}
func (p *UploadVolumeParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if p.Checksum.Valid {
		m["checksum"] = fmt.Sprint(p.Checksum.String)
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Format.Valid {
		m["format"] = fmt.Sprint(p.Format.String)
	}
	if p.Imagestoreuuid.Valid {
		m["imagestoreuuid"] = fmt.Sprint(p.Imagestoreuuid.String)
	}
	if p.Name.Valid {
		m["name"] = fmt.Sprint(p.Name.String)
	}
	if p.Projectid.Valid {
		m["projectid"] = fmt.Sprint(p.Projectid.String)
	}
	if p.Url.Valid {
		m["url"] = fmt.Sprint(p.Url.String)
	}
	if p.Zoneid.Valid {
		m["zoneid"] = fmt.Sprint(p.Zoneid.String)
	}
	return m
}
func (c *Client) UploadVolume(p UploadVolumeParameter) (Volume, error) {
	var v map[string]json.RawMessage
	var ret Volume
	b, err := c.Request("uploadVolume", p.ToMap())
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
