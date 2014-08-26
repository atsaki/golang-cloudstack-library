package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type ListServiceOfferingsParameter struct {
	Domainid         ID
	Id               ID
	Issystem         NullBool
	Keyword          NullString
	Name             NullString
	Page             NullInt64
	Pagesize         NullInt64
	Systemvmtype     NullString
	Virtualmachineid ID
}

func (p *ListServiceOfferingsParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *ListServiceOfferingsParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *ListServiceOfferingsParameter) SetIssystem(b bool) {
	p.Issystem.Bool = b
	p.Issystem.Valid = true
}
func (p *ListServiceOfferingsParameter) SetKeyword(s string) {
	p.Keyword.String = s
	p.Keyword.Valid = true
}
func (p *ListServiceOfferingsParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *ListServiceOfferingsParameter) SetPage(n int64) {
	p.Page.Int64 = n
	p.Page.Valid = true
}
func (p *ListServiceOfferingsParameter) SetPagesize(n int64) {
	p.Pagesize.Int64 = n
	p.Pagesize.Valid = true
}
func (p *ListServiceOfferingsParameter) SetSystemvmtype(s string) {
	p.Systemvmtype.String = s
	p.Systemvmtype.Valid = true
}
func (p *ListServiceOfferingsParameter) SetVirtualmachineid(s string) {
	p.Virtualmachineid.String = s
	p.Virtualmachineid.Valid = true
}
func (p *ListServiceOfferingsParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if p.Issystem.Valid {
		m["issystem"] = fmt.Sprint(p.Issystem.Bool)
	}
	if p.Keyword.Valid {
		m["keyword"] = fmt.Sprint(p.Keyword.String)
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
	if p.Systemvmtype.Valid {
		m["systemvmtype"] = fmt.Sprint(p.Systemvmtype.String)
	}
	if p.Virtualmachineid.Valid {
		m["virtualmachineid"] = fmt.Sprint(p.Virtualmachineid.String)
	}
	return m
}

type ListServiceOfferingsResponse struct {
	Count           float64 `json:"count"`
	Serviceoffering []struct {
		Cpunumber                 NullInt64  `json:"cpunumber"`
		Cpuspeed                  NullInt64  `json:"cpuspeed"`
		Created                   NullString `json:"created"`
		Defaultuse                NullBool   `json:"defaultuse"`
		Deploymentplanner         NullString `json:"deploymentplanner"`
		DiskBytesReadRate         NullInt64  `json:"diskbytesreadrate"`
		DiskBytesWriteRate        NullInt64  `json:"diskbyteswriterate"`
		DiskIopsReadRate          NullInt64  `json:"diskiopsreadrate"`
		DiskIopsWriteRate         NullInt64  `json:"diskiopswriterate"`
		Displaytext               NullString `json:"displaytext"`
		Domain                    NullString `json:"domain"`
		Domainid                  ID         `json:"domainid"`
		Hosttags                  NullString `json:"hosttags"`
		Hypervisorsnapshotreserve NullInt64  `json:"hypervisorsnapshotreserve"`
		Id                        ID         `json:"id"`
		Iscustomized              NullBool   `json:"iscustomized"`
		Iscustomizediops          NullBool   `json:"iscustomizediops"`
		Issystem                  NullBool   `json:"issystem"`
		Isvolatile                NullBool   `json:"isvolatile"`
		Limitcpuuse               NullBool   `json:"limitcpuuse"`
		Maxiops                   NullInt64  `json:"maxiops"`
		Memory                    NullInt64  `json:"memory"`
		Miniops                   NullInt64  `json:"miniops"`
		Name                      NullString `json:"name"`
		Networkrate               NullInt64  `json:"networkrate"`
		Offerha                   NullBool   `json:"offerha"`
		Serviceofferingdetails    NullString `json:"serviceofferingdetails"`
		Storagetype               NullString `json:"storagetype"`
		Systemvmtype              NullString `json:"systemvmtype"`
		Tags                      NullString `json:"tags"`
	} `json:"serviceoffering"`
}

func (c *Client) ListServiceOfferings(p ListServiceOfferingsParameter) (*ListServiceOfferingsResponse, error) {
	resp := new(ListServiceOfferingsResponse)
	b, err := c.Request("listServiceOfferings", p.ToMap())
	if err != nil {
		log.Println("Request failed:", err)
		return nil, err
	}
	err = json.Unmarshal(b, resp)
	if err != nil {
		log.Println("json.Unmarshal failed:", err)
	}
	return resp, err
}
