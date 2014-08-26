package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
)

type StopVirtualMachineParameter struct {
	Forced NullBool
	Id     ID
}

func (p *StopVirtualMachineParameter) SetForced(b bool) {
	p.Forced.Bool = b
	p.Forced.Valid = true
}
func (p *StopVirtualMachineParameter) SetId(s string) {
	p.Id.String = s
	p.Id.Valid = true
}
func (p *StopVirtualMachineParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Forced.Valid {
		m["forced"] = fmt.Sprint(p.Forced.Bool)
	}
	if p.Id.Valid {
		m["id"] = fmt.Sprint(p.Id.String)
	}
	return m
}

type StopVirtualMachineResponse struct {
	Virtualmachine struct {
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

func (c *Client) StopVirtualMachine(p StopVirtualMachineParameter) (*StopVirtualMachineResponse, error) {
	resp := new(StopVirtualMachineResponse)
	b, err := c.Request("stopVirtualMachine", p.ToMap())
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
