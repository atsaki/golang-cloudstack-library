package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListDiskOfferings(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		responses := map[string]string{
			"listDiskOfferings": `
{
    "listdiskofferingsresponse": {
        "count": 2, 
        "diskoffering": [
            {
                "created": "2014-10-06T16:13:52+0900", 
                "disksize": 0, 
                "displayoffering": true, 
                "displaytext": "Custom Disk", 
                "id": "00000000-0000-0000-0000-000000000000", 
                "iscustomized": true, 
                "name": "Custom Disk", 
                "storagetype": "shared", 
                "tags": "DATA"
            }, 
            {
                "created": "2014-11-18T18:23:50+0900", 
                "disksize": 800, 
                "displayoffering": true, 
                "displaytext": "DATADISK-LOCAL", 
                "id": "11111111-1111-1111-1111-111111111111", 
                "iscustomized": false, 
                "name": "DATADISK-LOCAL", 
                "storagetype": "local", 
                "tags": "local"
            }
        ]
    }
}
`,
		}

		fmt.Fprintln(w, responses[r.FormValue("command")])
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewListDiskOfferingsParameter()
	dos, err := client.ListDiskOfferings(p)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(dos) != 2 {
		t.Errorf("length: actual %d, expected 2", len(dos))
	}
	if dos[0].Name.String() != "Custom Disk" {
		t.Errorf("name: actual %s, expected \"Custom Disk\"", dos[0].Name.String())
	}
}
