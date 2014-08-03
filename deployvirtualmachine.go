package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type DeployVirtualMachineParameter struct {
	Account            NullString
	Affinitygroupids   []string
	Affinitygroupnames []string
	Customid           NullString
	Deploymentplanner  NullString
	Details            map[string]string
	Diskofferingid     ID
	Displayname        NullString
	Displayvm          NullBool
	Domainid           ID
	Group              NullString
	Hostid             ID
	Hypervisor         NullString
	Ip6address         NullString
	Ipaddress          NullString
	Iptonetworklist    map[string]string
	Keyboard           NullString
	Keypair            NullString
	Name               NullString
	Networkids         []string
	Projectid          ID
	Rootdisksize       NullInt64
	Securitygroupids   []string
	Securitygroupnames []string
	Serviceofferingid  ID
	Size               NullInt64
	Startvm            NullBool
	Templateid         ID
	Userdata           NullString
	Zoneid             ID
}

func (p *DeployVirtualMachineParameter) SetAccount(s string) {
	p.Account.String = s
	p.Account.Valid = true
}
func (p *DeployVirtualMachineParameter) SetAffinitygroupids(l []string) {
	p.Affinitygroupids = l
}
func (p *DeployVirtualMachineParameter) SetAffinitygroupnames(l []string) {
	p.Affinitygroupnames = l
}
func (p *DeployVirtualMachineParameter) SetCustomid(s string) {
	p.Customid.String = s
	p.Customid.Valid = true
}
func (p *DeployVirtualMachineParameter) SetDeploymentplanner(s string) {
	p.Deploymentplanner.String = s
	p.Deploymentplanner.Valid = true
}
func (p *DeployVirtualMachineParameter) SetDetails(m map[string]string) {
	p.Details = m
}
func (p *DeployVirtualMachineParameter) SetDiskofferingid(s string) {
	p.Diskofferingid.String = s
	p.Diskofferingid.Valid = true
}
func (p *DeployVirtualMachineParameter) SetDisplayname(s string) {
	p.Displayname.String = s
	p.Displayname.Valid = true
}
func (p *DeployVirtualMachineParameter) SetDisplayvm(b bool) {
	p.Displayvm.Bool = b
	p.Displayvm.Valid = true
}
func (p *DeployVirtualMachineParameter) SetDomainid(s string) {
	p.Domainid.String = s
	p.Domainid.Valid = true
}
func (p *DeployVirtualMachineParameter) SetGroup(s string) {
	p.Group.String = s
	p.Group.Valid = true
}
func (p *DeployVirtualMachineParameter) SetHostid(s string) {
	p.Hostid.String = s
	p.Hostid.Valid = true
}
func (p *DeployVirtualMachineParameter) SetHypervisor(s string) {
	p.Hypervisor.String = s
	p.Hypervisor.Valid = true
}
func (p *DeployVirtualMachineParameter) SetIp6address(s string) {
	p.Ip6address.String = s
	p.Ip6address.Valid = true
}
func (p *DeployVirtualMachineParameter) SetIpaddress(s string) {
	p.Ipaddress.String = s
	p.Ipaddress.Valid = true
}
func (p *DeployVirtualMachineParameter) SetIptonetworklist(m map[string]string) {
	p.Iptonetworklist = m
}
func (p *DeployVirtualMachineParameter) SetKeyboard(s string) {
	p.Keyboard.String = s
	p.Keyboard.Valid = true
}
func (p *DeployVirtualMachineParameter) SetKeypair(s string) {
	p.Keypair.String = s
	p.Keypair.Valid = true
}
func (p *DeployVirtualMachineParameter) SetName(s string) {
	p.Name.String = s
	p.Name.Valid = true
}
func (p *DeployVirtualMachineParameter) SetNetworkids(l []string) {
	p.Networkids = l
}
func (p *DeployVirtualMachineParameter) SetProjectid(s string) {
	p.Projectid.String = s
	p.Projectid.Valid = true
}
func (p *DeployVirtualMachineParameter) SetRootdisksize(n int64) {
	p.Rootdisksize.Int64 = n
	p.Rootdisksize.Valid = true
}
func (p *DeployVirtualMachineParameter) SetSecuritygroupids(l []string) {
	p.Securitygroupids = l
}
func (p *DeployVirtualMachineParameter) SetSecuritygroupnames(l []string) {
	p.Securitygroupnames = l
}
func (p *DeployVirtualMachineParameter) SetServiceofferingid(s string) {
	p.Serviceofferingid.String = s
	p.Serviceofferingid.Valid = true
}
func (p *DeployVirtualMachineParameter) SetSize(n int64) {
	p.Size.Int64 = n
	p.Size.Valid = true
}
func (p *DeployVirtualMachineParameter) SetStartvm(b bool) {
	p.Startvm.Bool = b
	p.Startvm.Valid = true
}
func (p *DeployVirtualMachineParameter) SetTemplateid(s string) {
	p.Templateid.String = s
	p.Templateid.Valid = true
}
func (p *DeployVirtualMachineParameter) SetUserdata(s string) {
	p.Userdata.String = s
	p.Userdata.Valid = true
}
func (p *DeployVirtualMachineParameter) SetZoneid(s string) {
	p.Zoneid.String = s
	p.Zoneid.Valid = true
}
func (p *DeployVirtualMachineParameter) ToMap() map[string]string {
	m := map[string]string{}
	if p.Account.Valid {
		m["account"] = fmt.Sprint(p.Account.String)
	}
	if len(p.Affinitygroupids) > 0 {
		m["affinitygroupids"] = strings.Join(p.Affinitygroupids, ",")
	}
	if len(p.Affinitygroupnames) > 0 {
		m["affinitygroupnames"] = strings.Join(p.Affinitygroupnames, ",")
	}
	if p.Customid.Valid {
		m["customid"] = fmt.Sprint(p.Customid.String)
	}
	if p.Deploymentplanner.Valid {
		m["deploymentplanner"] = fmt.Sprint(p.Deploymentplanner.String)
	}
	if len(p.Details) > 0 {
		i := 0
		for key, value := range m {
			m[fmt.Sprintf("details[%d].key", i)] = key
			m[fmt.Sprintf("details[%d].value", i)] = value
			i += 1
		}
	}
	if p.Diskofferingid.Valid {
		m["diskofferingid"] = fmt.Sprint(p.Diskofferingid.String)
	}
	if p.Displayname.Valid {
		m["displayname"] = fmt.Sprint(p.Displayname.String)
	}
	if p.Displayvm.Valid {
		m["displayvm"] = fmt.Sprint(p.Displayvm.Bool)
	}
	if p.Domainid.Valid {
		m["domainid"] = fmt.Sprint(p.Domainid.String)
	}
	if p.Group.Valid {
		m["group"] = fmt.Sprint(p.Group.String)
	}
	if p.Hostid.Valid {
		m["hostid"] = fmt.Sprint(p.Hostid.String)
	}
	if p.Hypervisor.Valid {
		m["hypervisor"] = fmt.Sprint(p.Hypervisor.String)
	}
	if p.Ip6address.Valid {
		m["ip6address"] = fmt.Sprint(p.Ip6address.String)
	}
	if p.Ipaddress.Valid {
		m["ipaddress"] = fmt.Sprint(p.Ipaddress.String)
	}
	if len(p.Iptonetworklist) > 0 {
		i := 0
		for key, value := range m {
			m[fmt.Sprintf("iptonetworklist[%d].key", i)] = key
			m[fmt.Sprintf("iptonetworklist[%d].value", i)] = value
			i += 1
		}
	}
	if p.Keyboard.Valid {
		m["keyboard"] = fmt.Sprint(p.Keyboard.String)
	}
	if p.Keypair.Valid {
		m["keypair"] = fmt.Sprint(p.Keypair.String)
	}
	if p.Name.Valid {
		m["name"] = fmt.Sprint(p.Name.String)
	}
	if len(p.Networkids) > 0 {
		m["networkids"] = strings.Join(p.Networkids, ",")
	}
	if p.Projectid.Valid {
		m["projectid"] = fmt.Sprint(p.Projectid.String)
	}
	if p.Rootdisksize.Valid {
		m["rootdisksize"] = fmt.Sprint(p.Rootdisksize.Int64)
	}
	if len(p.Securitygroupids) > 0 {
		m["securitygroupids"] = strings.Join(p.Securitygroupids, ",")
	}
	if len(p.Securitygroupnames) > 0 {
		m["securitygroupnames"] = strings.Join(p.Securitygroupnames, ",")
	}
	if p.Serviceofferingid.Valid {
		m["serviceofferingid"] = fmt.Sprint(p.Serviceofferingid.String)
	}
	if p.Size.Valid {
		m["size"] = fmt.Sprint(p.Size.Int64)
	}
	if p.Startvm.Valid {
		m["startvm"] = fmt.Sprint(p.Startvm.Bool)
	}
	if p.Templateid.Valid {
		m["templateid"] = fmt.Sprint(p.Templateid.String)
	}
	if p.Userdata.Valid {
		m["userdata"] = fmt.Sprint(p.Userdata.String)
	}
	if p.Zoneid.Valid {
		m["zoneid"] = fmt.Sprint(p.Zoneid.String)
	}
	return m
}

type DeployVirtualMachineResponse struct {
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

func (c *Client) DeployVirtualMachine(p DeployVirtualMachineParameter) (*DeployVirtualMachineResponse, error) {
	resp := new(DeployVirtualMachineResponse)
	b, err := c.Request("deployVirtualMachine", p.ToMap())
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
