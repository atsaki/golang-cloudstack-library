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

		resp := `
{
    "listdiskofferingsresponse": {
        "count": 4,
        "diskoffering": [
            {
                "created": "2015-09-20T06:22:03+0000",
                "disksize": 5,
                "displayoffering": true,
                "displaytext": "Small Disk, 5 GB",
                "id": "018438d9-be77-4efd-a81e-ef8c73951acc",
                "iscustomized": false,
                "name": "Small",
                "provisioningtype": "thin",
                "storagetype": "shared"
            },
            {
                "created": "2015-09-20T06:22:03+0000",
                "disksize": 20,
                "displayoffering": true,
                "displaytext": "Medium Disk, 20 GB",
                "id": "61eca983-42b5-49d3-9230-9a771a12ebaf",
                "iscustomized": false,
                "name": "Medium",
                "provisioningtype": "thin",
                "storagetype": "shared"
            },
            {
                "created": "2015-09-20T06:22:03+0000",
                "disksize": 100,
                "displayoffering": true,
                "displaytext": "Large Disk, 100 GB",
                "id": "f58be420-a967-45c5-9531-2ed56d24e893",
                "iscustomized": false,
                "name": "Large",
                "provisioningtype": "thin",
                "storagetype": "shared"
            },
            {
                "created": "2015-09-20T06:22:03+0000",
                "disksize": 0,
                "displayoffering": true,
                "displaytext": "Custom Disk",
                "id": "7fc4314a-ca54-495b-83d6-973ed6ea06f6",
                "iscustomized": true,
                "name": "Custom",
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
	p := NewListDiskOfferingsParameter()
	dos, err := client.ListDiskOfferings(p)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(dos) != 4 {
		t.Errorf("length: actual %d, expected 4", len(dos))
	}

	if dos[0].Name.String() != "Small" {
		t.Errorf("name: actual %s, expected \"Small\"", dos[0].Name.String())
	}

	if dos[0].ProvisioningType.String() != "thin" {
		t.Errorf("provisioningtype: actual %s, expected \"thin\"", dos[0].ProvisioningType.String())
	}
}
