package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type ListVirtualMachinesParameter struct {
	Account           NullString
	Affinitygroupid   ID
	Details           []string
	Displayvm         NullBool
	Domainid          ID
	Forvirtualnetwork NullBool
	Groupid           ID
	Hostid            ID
	Hypervisor        NullString
	Id                ID
	Ids               []string
	Isoid             ID
	Isrecursive       NullBool
	Keyword           NullString
	Listall           NullBool
	Name              NullString
	Networkid         ID
	Page              NullInt64
	Pagesize          NullInt64
	Podid             ID
	Projectid         ID
	Serviceofferingid ID
	State             NullString
	Storageid         ID
	Tags              map[string]string
	Templateid        ID
	Vpcid             ID
	Zoneid            ID
}

func (p *ListVirtualMachinesParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *ListVirtualMachinesParameter) SetAffinitygroupid(s string) {
	p.Affinitygroupid.String = s
	p.Affinitygroupid.Valid = true
}
func (p *ListVirtualMachinesParameter) SetDetails(l []string) {
	p.Details = l
}
func (p *ListVirtualMachinesParameter) SetDisplayvm(b bool) {
	p.Displayvm.Bool = b
	p.Displayvm.Valid = true
}
func (p *ListVirtualMachinesParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *ListVirtualMachinesParameter) SetForvirtualnetwork(b bool) {
	p.Forvirtualnetwork.Bool = b
	p.Forvirtualnetwork.Valid = true
}
func (p *ListVirtualMachinesParameter) SetGroupid(s string) {
	p.Groupid.String = s
	p.Groupid.Valid = true
}
func (p *ListVirtualMachinesParameter) SetHostid(s string) {
	p.Hostid.String = s
	p.Hostid.Valid = true
}
func (p *ListVirtualMachinesParameter) SetHypervisor(s string) {
	p.Hypervisor.String = s
	p.Hypervisor.Valid = true
}
func (p *ListVirtualMachinesParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *ListVirtualMachinesParameter) SetIds(l []string) {
	p.Ids = l
}
func (p *ListVirtualMachinesParameter) SetIsoid(s string) {
	p.Isoid.String = s
	p.Isoid.Valid = true
}
func (p *ListVirtualMachinesParameter) SetIsrecursive(b bool) {
	p.Isrecursive.Bool = b
	p.Isrecursive.Valid = true
}
func (p *ListVirtualMachinesParameter) SetKeyword(s string) {
	p.Keyword.String = s
	p.Keyword.Valid = true
}
func (p *ListVirtualMachinesParameter) SetListall(b bool) {
	p.Listall.Bool = b
	p.Listall.Valid = true
}
func (p *ListVirtualMachinesParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *ListVirtualMachinesParameter) SetNetworkid(s string) {
	p.Networkid.String = s
	p.Networkid.Valid = true
}
func (p *ListVirtualMachinesParameter) SetPage(n int64) {
	p.Page.Int64 = n
	p.Page.Valid = true
}
func (p *ListVirtualMachinesParameter) SetPagesize(n int64) {
	p.Pagesize.Int64 = n
	p.Pagesize.Valid = true
}
func (p *ListVirtualMachinesParameter) SetPodid(s string) {
	p.Podid.String = s
	p.Podid.Valid = true
}
func (p *ListVirtualMachinesParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *ListVirtualMachinesParameter) SetServiceofferingid(s string) {
	p.Serviceofferingid.String = s
	p.Serviceofferingid.Valid = true
}
func (p *ListVirtualMachinesParameter) SetState(s string) {
	p.State.String = s
	p.State.Valid = true
}
func (p *ListVirtualMachinesParameter) SetStorageid(s string) {
	p.Storageid.String = s
	p.Storageid.Valid = true
}
func (p *ListVirtualMachinesParameter) SetTags(m map[string]string) {
	p.Tags = m
}
func (p *ListVirtualMachinesParameter) SetTemplateid(s string) {
	p.Templateid.String = s
	p.Templateid.Valid = true
}
func (p *ListVirtualMachinesParameter) SetVpcid(s string) {
	p.Vpcid.String = s
	p.Vpcid.Valid = true
}
func (p *ListVirtualMachinesParameter) SetZoneid(s string) {
	p.Zoneid.String = s
	p.Zoneid.Valid = true
}
func (p *ListVirtualMachinesParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if p.Affinitygroupid.Valid {
		m["affinitygroupid"] = fmt.Sprint(p.Affinitygroupid.String)
	}
	if len(p.Details) > 0 {
		m["details"] = strings.Join(p.Details, ",")
	}
	if p.Displayvm.Valid {
		m["displayvm"] = fmt.Sprint(p.Displayvm.Bool)
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Forvirtualnetwork.Valid {
		m["forvirtualnetwork"] = fmt.Sprint(p.Forvirtualnetwork.Bool)
	}
	if p.Groupid.Valid {
		m["groupid"] = fmt.Sprint(p.Groupid.String)
	}
	if p.Hostid.Valid {
		m["hostid"] = fmt.Sprint(p.Hostid.String)
	}
	if p.Hypervisor.Valid {
		m["hypervisor"] = fmt.Sprint(p.Hypervisor.String)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	if len(p.Ids) > 0 {
		m["ids"] = strings.Join(p.Ids, ",")
	}
	if p.Isoid.Valid {
		m["isoid"] = fmt.Sprint(p.Isoid.String)
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
	if p.Networkid.Valid {
		m["networkid"] = fmt.Sprint(p.Networkid.String)
	}
	if p.Page.Valid {
		m["page"] = fmt.Sprint(p.Page.Int64)
	}
	if p.Pagesize.Valid {
		m["pagesize"] = fmt.Sprint(p.Pagesize.Int64)
	}
	if p.Podid.Valid {
		m["podid"] = fmt.Sprint(p.Podid.String)
	}
	if p.Projectid.Valid {
		m["projectid"] = fmt.Sprint(p.Projectid.String)
	}
	if p.Serviceofferingid.Valid {
		m["serviceofferingid"] = fmt.Sprint(p.Serviceofferingid.String)
	}
	if p.State.Valid {
		m["state"] = fmt.Sprint(p.State.String)
	}
	if p.Storageid.Valid {
		m["storageid"] = fmt.Sprint(p.Storageid.String)
	}
	if len(p.Tags) > 0 {
		i := 0
		for key, value := range m {
			m[fmt.Sprintf("tags[%d].key", i)] = key
			m[fmt.Sprintf("tags[%d].value", i)] = value
			i += 1
		}
	}
	if p.Templateid.Valid {
		m["templateid"] = fmt.Sprint(p.Templateid.String)
	}
	if p.Vpcid.Valid {
		m["vpcid"] = fmt.Sprint(p.Vpcid.String)
	}
	if p.Zoneid.Valid {
		m["zoneid"] = fmt.Sprint(p.Zoneid.String)
	}
	return m
}

type ListVirtualMachinesResponse struct {
	Count          float64 `json:"count"`
	Virtualmachine []struct {
		Account       NullString `json:"account"`
		Affinitygroup []struct {
			Account           NullString `json:"account"`
			Description       NullString `json:"description"`
			Domain            NullString `json:"domain"`
			Domainid          ID         `json:"domainid"`
			Id                ID         `json:"id"`
			Name              NullString `json:"name"`
			Type              NullString `json:"type"`
			VirtualmachineIds []ID       `json:"virtualmachineids"`
		} `json:"affinitygroup"`
		Cpunumber             NullInt64  `json:"cpunumber"`
		Cpuspeed              NullInt64  `json:"cpuspeed"`
		Cpuused               NullString `json:"cpuused"`
		Created               NullString `json:"created"`
		Details               NullString `json:"details"`
		Diskioread            NullInt64  `json:"diskioread"`
		Diskiowrite           NullInt64  `json:"diskiowrite"`
		Diskkbsread           NullInt64  `json:"diskkbsread"`
		Diskkbswrite          NullInt64  `json:"diskkbswrite"`
		Diskofferingid        ID         `json:"diskofferingid"`
		Diskofferingname      NullString `json:"diskofferingname"`
		Displayname           NullString `json:"displayname"`
		Displayvm             NullBool   `json:"displayvm"`
		Domain                NullString `json:"domain"`
		Domainid              ID         `json:"domainid"`
		Forvirtualnetwork     NullBool   `json:"forvirtualnetwork"`
		Group                 NullString `json:"group"`
		Groupid               ID         `json:"groupid"`
		Guestosid             ID         `json:"guestosid"`
		Haenable              NullBool   `json:"haenable"`
		Hostid                ID         `json:"hostid"`
		Hostname              NullString `json:"hostname"`
		Hypervisor            NullString `json:"hypervisor"`
		Id                    ID         `json:"id"`
		Instancename          NullString `json:"instancename"`
		Isdynamicallyscalable NullBool   `json:"isdynamicallyscalable"`
		Isodisplaytext        NullString `json:"isodisplaytext"`
		Isoid                 ID         `json:"isoid"`
		Isoname               NullString `json:"isoname"`
		Keypair               NullString `json:"keypair"`
		Memory                NullInt64  `json:"memory"`
		Name                  NullString `json:"name"`
		Networkkbsread        NullInt64  `json:"networkkbsread"`
		Networkkbswrite       NullInt64  `json:"networkkbswrite"`
		Nic                   []struct {
			Broadcasturi     NullString   `json:"broadcasturi"`
			Deviceid         ID           `json:"deviceid"`
			Gateway          NullString   `json:"gateway"`
			Id               ID           `json:"id"`
			Ip6address       NullString   `json:"ip6address"`
			Ip6cidr          NullString   `json:"ip6cidr"`
			Ip6gateway       NullString   `json:"ip6gateway"`
			Ipaddress        NullString   `json:"ipaddress"`
			Isdefault        NullBool     `json:"isdefault"`
			Isolationuri     NullString   `json:"isolationuri"`
			Macaddress       NullString   `json:"macaddress"`
			Netmask          NullString   `json:"netmask"`
			Networkid        ID           `json:"networkid"`
			Networkname      NullString   `json:"networkname"`
			Secondaryip      []NullString `json:"secondaryip"`
			Traffictype      NullString   `json:"traffictype"`
			Type             NullString   `json:"type"`
			Virtualmachineid ID           `json:"virtualmachineid"`
		} `json:"nic"`
		Ostypeid        ID         `json:"ostypeid"`
		Password        NullString `json:"password"`
		Passwordenabled NullBool   `json:"passwordenabled"`
		Project         NullString `json:"project"`
		Projectid       ID         `json:"projectid"`
		Publicip        NullString `json:"publicip"`
		Publicipid      ID         `json:"publicipid"`
		Rootdeviceid    ID         `json:"rootdeviceid"`
		Rootdevicetype  NullString `json:"rootdevicetype"`
		Securitygroup   []struct {
			Account     NullString `json:"account"`
			Description NullString `json:"description"`
			Domain      NullString `json:"domain"`
			Domainid    ID         `json:"domainid"`
			Egressrule  []struct {
				Account           NullString `json:"account"`
				Cidr              NullString `json:"cidr"`
				Endport           NullInt64  `json:"endport"`
				Icmpcode          NullInt64  `json:"icmpcode"`
				Icmptype          NullInt64  `json:"icmptype"`
				Protocol          NullString `json:"protocol"`
				Ruleid            ID         `json:"ruleid"`
				Securitygroupname NullString `json:"securitygroupname"`
				Startport         NullInt64  `json:"startport"`
			} `json:"egressrule"`
			Id          ID `json:"id"`
			Ingressrule []struct {
				Account           NullString `json:"account"`
				Cidr              NullString `json:"cidr"`
				Endport           NullInt64  `json:"endport"`
				Icmpcode          NullInt64  `json:"icmpcode"`
				Icmptype          NullInt64  `json:"icmptype"`
				Protocol          NullString `json:"protocol"`
				Ruleid            ID         `json:"ruleid"`
				Securitygroupname NullString `json:"securitygroupname"`
				Startport         NullInt64  `json:"startport"`
			} `json:"ingressrule"`
			Name      NullString `json:"name"`
			Project   NullString `json:"project"`
			Projectid ID         `json:"projectid"`
			Tags      []struct {
				Account      NullString `json:"account"`
				Customer     NullString `json:"customer"`
				Domain       NullString `json:"domain"`
				Domainid     ID         `json:"domainid"`
				Key          NullString `json:"key"`
				Project      NullString `json:"project"`
				Projectid    ID         `json:"projectid"`
				Resourceid   ID         `json:"resourceid"`
				Resourcetype NullString `json:"resourcetype"`
				Value        NullString `json:"value"`
			} `json:"tags"`
		} `json:"securitygroup"`
		Serviceofferingid   ID         `json:"serviceofferingid"`
		Serviceofferingname NullString `json:"serviceofferingname"`
		Servicestate        NullString `json:"servicestate"`
		State               NullString `json:"state"`
		Tags                []struct {
			Account      NullString `json:"account"`
			Customer     NullString `json:"customer"`
			Domain       NullString `json:"domain"`
			Domainid     ID         `json:"domainid"`
			Key          NullString `json:"key"`
			Project      NullString `json:"project"`
			Projectid    ID         `json:"projectid"`
			Resourceid   ID         `json:"resourceid"`
			Resourcetype NullString `json:"resourcetype"`
			Value        NullString `json:"value"`
		} `json:"tags"`
		Templatedisplaytext NullString `json:"templatedisplaytext"`
		Templateid          ID         `json:"templateid"`
		Templatename        NullString `json:"templatename"`
		Vgpu                NullString `json:"vgpu"`
		Zoneid              ID         `json:"zoneid"`
		Zonename            NullString `json:"zonename"`
	} `json:"virtualmachine"`
}

func (c *Client) ListVirtualMachines(p ListVirtualMachinesParameter) (*ListVirtualMachinesResponse, error) {
	resp := new(ListVirtualMachinesResponse)
	b, err := c.Request("listVirtualMachines", p.ToMap())
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
