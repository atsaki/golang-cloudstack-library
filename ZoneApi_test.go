package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListZones(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listzonesresponse": {
        "count": 1, 
        "zone": [
            {
                "allocationstate": "Enabled", 
                "dhcpprovider": "VirtualRouter", 
                "id": "0ae80649-3a41-4fcd-8412-5413fc48f398", 
                "localstorageenabled": true, 
                "name": "zone1", 
                "networktype": "Advanced", 
                "securitygroupsenabled": false, 
                "tags": [], 
                "zonetoken": "412835d8-ad99-4dc1-a123-bc181aa152f0"
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
	p := NewListZonesParameter()
	zones, _ := client.ListZones(p)

	if len(zones) != 1 {
		t.Errorf("length: actual %d, expected 1", len(zones))
	}

	if zones[0].Name.String() != "zone1" {
		t.Errorf("name: actual %s, expected zone1", zones[0].Name.String())
	}

	if zones[0].Id.String() != "0ae80649-3a41-4fcd-8412-5413fc48f398" {
		t.Errorf("Unexpected length: actual %s, expected 0ae80649-3a41-4fcd-8412-5413fc48f398",
			zones[0].Id.String())
	}
}
