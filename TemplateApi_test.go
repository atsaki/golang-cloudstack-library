package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListTemplates(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listtemplatesresponse": {
        "count": 2,
        "template": [
            {
                "account": "system",
                "checksum": "",
                "created": "2016-07-25T15:26:19+0000",
                "crossZones": true,
                "displaytext": "SystemVM Template (simulator)",
                "domain": "ROOT",
                "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "format": "VHD",
                "hypervisor": "Simulator",
                "id": "718d8b7b-5f5a-11e5-bc86-0242ac11180a",
                "isdynamicallyscalable": false,
                "isextractable": false,
                "isfeatured": false,
                "ispublic": false,
                "isready": true,
                "name": "SystemVM Template (simulator)",
                "ostypeid": "629901ea-5f5a-11e5-bc86-0242ac11180a",
                "ostypename": "Debian GNU/Linux 5.0 (64-bit)",
                "passwordenabled": false,
                "sshkeyenabled": false,
                "status": "Download Complete",
                "tags": [],
                "templatetype": "SYSTEM",
                "zoneid": "f902785c-9dda-45c0-afc3-f1a29d202b2e",
                "zonename": "Sandbox-simulator"
            },
            {
                "account": "system",
                "checksum": "",
                "crossZones": true,
                "displaytext": "CentOS 5.3(64-bit) no GUI (Simulator)",
                "domain": "ROOT",
                "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "format": "VHD",
                "hypervisor": "Simulator",
                "id": "718d9c8f-5f5a-11e5-bc86-0242ac11180a",
                "isdynamicallyscalable": false,
                "isextractable": false,
                "isfeatured": true,
                "ispublic": true,
                "isready": false,
                "name": "CentOS 5.3(64-bit) no GUI (Simulator)",
                "ostypeid": "6298c0f8-5f5a-11e5-bc86-0242ac11180a",
                "ostypename": "CentOS 5.3 (32-bit)",
                "passwordenabled": false,
                "sshkeyenabled": false,
                "tags": [],
                "templatetype": "BUILTIN",
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
	p := NewListTemplatesParameter("all")
	templates, _ := client.ListTemplates(p)

	if len(templates) != 2 {
		t.Errorf("length: actual %d, expected 2", len(templates))
	}

	if templates[1].Name.String() != "CentOS 5.3(64-bit) no GUI (Simulator)" {
		t.Errorf("name: actual %s, expected \"CentOS 5.3(64-bit) no GUI (Simulator)\"", templates[0].Name.String())
	}
}

func TestListTemplatePermissions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listtemplatepermissionsresponse": {
        "templatepermission": {
            "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
            "id": "718d9c8f-5f5a-11e5-bc86-0242ac11180a",
            "ispublic": true
        }
    }
}
		`
		fmt.Fprintln(w, resp)
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewListTemplatePermissionsParameter("718d9c8f-5f5a-11e5-bc86-0242ac11180a")
	permission, _ := client.ListTemplatePermissions(p)

	if !permission.IsPublic.Bool() {
		t.Errorf("ispublic: actual false, expected true", permission.IsPublic.Bool())
	}
}
