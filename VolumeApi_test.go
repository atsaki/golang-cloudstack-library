package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListVolumes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listvolumesresponse": {
        "count": 1,
        "volume": [
            {
                "account": "admin",
                "created": "2016-07-21T13:32:53+0000",
                "destroyed": false,
                "deviceid": 0,
                "displayvolume": true,
                "domain": "ROOT",
                "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "hypervisor": "Simulator",
                "id": "9aa5166e-5965-4788-9238-68a7256d3937",
                "isextractable": false,
                "name": "ROOT-4",
                "path": "f5fca5ba-14d9-41d1-8f52-634ff4495ea9",
                "provisioningtype": "thin",
                "quiescevm": false,
                "serviceofferingdisplaytext": "Small Instance",
                "serviceofferingid": "6ac77bbf-1bf1-4625-83aa-5905d9fe9781",
                "serviceofferingname": "Small Instance",
                "size": 100,
                "state": "Ready",
                "storage": "PS2",
                "storageid": "0bbd54fe-b348-3e3d-9b91-b453a21bc43a",
                "storagetype": "shared",
                "tags": [],
                "templatedisplaytext": "CentOS 5.3(64-bit) no GUI (Simulator)",
                "templateid": "718d9c8f-5f5a-11e5-bc86-0242ac11180a",
                "templatename": "CentOS 5.3(64-bit) no GUI (Simulator)",
                "type": "ROOT",
                "virtualmachineid": "869cad78-2232-4c4e-8018-f38462cbc1df",
                "vmdisplayname": "vm01",
                "vmname": "vm01",
                "vmstate": "Running",
                "zoneid": "f902785c-9dda-45c0-afc3-f1a29d202b2e",
                "zonename": "Sandbox-simulator"
            }
        ]
    }
}
		`
		fmt.Fprintln(w, resp)
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewListVolumesParameter()
	volumes, _ := client.ListVolumes(p)

	if len(volumes) != 1 {
		t.Errorf("length: actual %d, expected 1", len(volumes))
	}

	if volumes[0].Name.String() != "ROOT-4" {
		t.Errorf("name: actual %s, expected \"ROOT-4\"", volumes[0].Name.String())
	}

	if volumes[0].ProvisioningType.String() != "thin" {
		t.Errorf("provisioningtype: actual %s, expected \"thin\"", volumes[0].ProvisioningType.String())
	}
}

func TestDeleteVolume(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		responses := map[string]string{
			"deleteVolume": `{"deletevolumeresponse":{"success":"true"}}`,
		}

		fmt.Fprintln(w, responses[r.FormValue("command")])
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewDeleteVolumeParameter("0728020b-cdee-4f63-bb96-de831c384a73")
	result, err := client.DeleteVolume(p)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !result.Success.Bool() {
		t.Errorf("success: actual true, expected true")
	}

}
