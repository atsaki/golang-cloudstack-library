package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestVirtualMachineMethods(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		responses := map[string]string{
			"listVirtualMachines": `

{
    "listvirtualmachinesresponse": {
        "count": 1, 
        "virtualmachine": [
            {
                "account": "account1", 
                "affinitygroup": [], 
                "cpunumber": 1, 
                "cpuspeed": 800, 
                "cpuused": "0%", 
                "created": "2014-11-24T12:13:45+0900", 
                "displayname": "vm01", 
                "displayvm": true, 
                "domain": "domain1", 
                "domainid": "00000000-0000-0000-0000-000000000000", 
                "guestosid": "00000000-0000-0000-0000-000000000000", 
                "haenable": false, 
                "hypervisor": "VMware", 
                "id": "00000000-0000-0000-0000-000000000000", 
                "isdynamicallyscalable": true, 
                "keypair": "sshkey", 
                "memory": 1024, 
                "name": "vm01", 
                "networkkbsread": 0, 
                "networkkbswrite": 0, 
                "nic": [
                    {
                        "broadcasturi": "vlan://10", 
                        "gateway": "10.3.0.1", 
                        "id": "00000000-0000-0000-0000-000000000000", 
                        "ipaddress": "10.3.0.2", 
                        "isdefault": true, 
                        "isolationuri": "vlan://10", 
                        "macaddress": "00:00:00:00:00:00", 
                        "netmask": "255.255.252.0", 
                        "networkid": "00000000-0000-0000-0000-000000000000", 
                        "networkname": "network1", 
                        "traffictype": "Guest", 
                        "type": "Isolated"
                    }
                ], 
                "passwordenabled": true, 
                "publicip": "1.1.1.1", 
                "publicipid": "00000000-0000-0000-0000-000000000000", 
                "rootdeviceid": 0, 
                "rootdevicetype": "ROOT", 
                "securitygroup": [], 
                "serviceofferingid": "00000000-0000-0000-0000-000000000000", 
                "serviceofferingname": "t1.micro", 
                "state": "Running", 
                "tags": [], 
                "templatedisplaytext": "CentOS65", 
                "templateid": "00000000-0000-0000-0000-000000000000", 
                "templatename": "CentOS 6.5 64-bit", 
                "zoneid": "00000000-0000-0000-0000-000000000000", 
                "zonename": "zone1"
            }
        ]
    }
}
			`,
			"destroyVirtualMachine": `
{
    "destroyvirtualmachineresponse": {
        "jobid": "94065e9f-680c-484f-9586-f69ea1a92818"
    }
}
`,
			"queryAsyncJobResult": `
{
    "queryasyncjobresultresponse": {
        "accountid": "00000000-0000-0000-0000-000000000000", 
        "cmd": "org.apache.cloudstack.api.command.user.vm.DestroyVMCmd", 
        "created": "2014-11-29T17:16:05+0900", 
        "jobid": "94065e9f-680c-484f-9586-f69ea1a92818", 
        "jobprocstatus": 0, 
        "jobresult": {
            "null": {
                "affinitygroup": [], 
                "nic": [], 
                "securitygroup": [], 
                "tags": []
            }
        }, 
        "jobresultcode": 0, 
        "jobresulttype": "object", 
        "jobstatus": 1, 
        "userid": "00000000-0000-0000-0000-000000000000"
    }
}
`,
		}

		fmt.Fprintln(w, responses[r.FormValue("command")])
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewListVirtualMachinesParameter()
	vms, err := client.ListVirtualMachines(p)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(vms) != 1 {
		t.Errorf("length: actual %d, expected 1", len(vms))
	}

	// Check VirtualMachine implements Resource
	var r Resource
	r = vms[0]

	vm := new(VirtualMachine)
	*vm = *(r.(*VirtualMachine))

	// modify original value
	vm.Nic[0].IpAddress.Set("8.8.8.8")
	if vm.Nic[0].IpAddress.String() != "8.8.8.8" {
		t.Errorf("ipaddress: actual %s, expected 8.8.8.8", vm.Nic[0].IpAddress.String())
	}

	if err = vm.Refresh(); err != nil {
		t.Error(err)
	}

	if len(vm.Nic) != 1 {
		t.Errorf("nic length: actual %d, expected 1", len(vms))
	}
	// check original value is restored
	if vm.Nic[0].IpAddress.String() != "10.3.0.2" {
		t.Errorf("ipaddress: actual %s, expected 10.3.0.2", vm.Nic[0].IpAddress.String())
	}
	if vm.client.EndPoint != endpoint {
		t.Errorf("endpoint: actual %v, expected %v", vms[0].client.EndPoint, endpoint)
	}

	if err = vm.Delete(); err != nil {
		t.Error(err)
	}
}
