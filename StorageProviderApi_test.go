package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListStorageProviders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "liststorageprovidersresponse": {
        "count": 4,
        "dataStoreProvider": [
            {
                "name": "DefaultPrimary",
                "type": "PRIMARY"
            },
            {
                "name": "SolidFire",
                "type": "PRIMARY"
            },
            {
                "name": "CloudByte",
                "type": "PRIMARY"
            },
            {
                "name": "SolidFireShared",
                "type": "PRIMARY"
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
	p := NewListStorageProvidersParameter("primary")
	providers, _ := client.ListStorageProviders(p)

	if len(providers) != 4 {
		t.Errorf("length: actual %d, expected 4", len(providers))
	}

	if providers[0].Name.String() != "DefaultPrimary" {
		t.Errorf("name: actual %s, expected DefaultPrimary", providers[0].Name.String())
	}

	if providers[0].Type.String() != "PRIMARY" {
		t.Errorf("type: actual %s, expected PRIMARY", providers[0].Type.String())
	}
}
