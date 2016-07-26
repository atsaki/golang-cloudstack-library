package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListServiceOfferings(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listserviceofferingsresponse": {
        "count": 2,
        "serviceoffering": [
            {
                "cpunumber": 1,
                "cpuspeed": 500,
                "created": "2015-09-20T06:22:03+0000",
                "defaultuse": false,
                "displaytext": "Small Instance",
                "id": "6ac77bbf-1bf1-4625-83aa-5905d9fe9781",
                "iscustomized": false,
                "issystem": false,
                "isvolatile": false,
                "limitcpuuse": false,
                "memory": 512,
                "name": "Small Instance",
                "offerha": false,
                "provisioningtype": "thin",
                "storagetype": "shared"
            },
            {
                "cpunumber": 1,
                "cpuspeed": 1000,
                "created": "2015-09-20T06:22:03+0000",
                "defaultuse": false,
                "displaytext": "Medium Instance",
                "id": "e53c604d-558a-4b0e-ad62-6a72f83fcd8e",
                "iscustomized": false,
                "issystem": false,
                "isvolatile": false,
                "limitcpuuse": false,
                "memory": 1024,
                "name": "Medium Instance",
                "offerha": false,
                "provisioningtype": "thin",
                "storagetype": "shared"
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
	p := NewListServiceOfferingsParameter()
	offerings, _ := client.ListServiceOfferings(p)

	if len(offerings) != 2 {
		t.Errorf("length: actual %d, expected 2", len(offerings))
	}

	if offerings[0].Name.String() != "Small Instance" {
		t.Errorf("name: actual %s, expected \"Small Instance\"", offerings[0].Name.String())
	}
}
