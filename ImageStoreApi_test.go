package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListImageStores(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listimagestoresresponse": {
        "count": 1,
        "imagestore": [
            {
                "details": [],
                "id": "dd11f2c4-a262-49f9-8658-3fd4cf0fa1ba",
                "name": "nfs://10.147.28.6:/export/home/sandbox/secondary",
                "protocol": "nfs",
                "providername": "NFS",
                "scope": "ZONE",
                "url": "nfs://10.147.28.6:/export/home/sandbox/secondary",
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
	p := NewListImageStoresParameter()
	stores, _ := client.ListImageStores(p)

	if len(stores) != 1 {
		t.Errorf("length: actual %d, expected 1", len(stores))
	}

	if stores[0].Name.String() != "nfs://10.147.28.6:/export/home/sandbox/secondary" {
		t.Errorf("name: actual %s, expected nfs://10.147.28.6:/export/home/sandbox/secondary", stores[0].Name.String())
	}
}

func TestAddImageStore(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "addimagestoreresponse": {
        "imagestore": {
            "details": [],
            "id": "bfaf15d1-782d-4757-9097-d36466b7ee11",
            "name": "nfs://10.147.28.6/export/home/sandbox/secondary",
            "protocol": "nfs",
            "providername": "NFS",
            "scope": "ZONE",
            "url": "nfs://10.147.28.6/export/home/sandbox/secondary",
            "zoneid": "f902785c-9dda-45c0-afc3-f1a29d202b2e",
            "zonename": "Sandbox-simulator"
        }
    }
}
		`
		fmt.Fprintln(w, resp)
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	url := "nfs://10.147.28.6/export/home/sandbox/secondary"
	zoneid := "f902785c-9dda-45c0-afc3-f1a29d202b2e"
	p := NewAddImageStoreParameter("nfs")
	p.Url.Set(url)
	p.ZoneId.Set(zoneid)
	store, _ := client.AddImageStore(p)

	if store.Name.String() != "nfs://10.147.28.6/export/home/sandbox/secondary" {
		t.Errorf("name: actual %s, expected nfs://10.147.28.6/export/home/sandbox/secondary", store.Name.String())
	}
}

func TestAddSecondaryStorage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "addsecondarystorageresponse": {
        "secondarystorage": {
            "details": [],
            "id": "c4ddf9c4-901a-406b-96f4-fd6d8d89731e",
            "name": "nfs://10.147.28.6/export/home/sandbox/secondary",
            "protocol": "nfs",
            "providername": "NFS",
            "scope": "ZONE",
            "url": "nfs://10.147.28.6/export/home/sandbox/secondary",
            "zoneid": "f902785c-9dda-45c0-afc3-f1a29d202b2e",
            "zonename": "Sandbox-simulator"
        }
    }
}
		`
		fmt.Fprintln(w, resp)
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	url := "nfs://10.147.28.6/export/home/sandbox/secondary"
	zoneid := "f902785c-9dda-45c0-afc3-f1a29d202b2e"
	p := NewAddSecondaryStorageParameter(url)
	p.ZoneId.Set(zoneid)
	store, _ := client.AddSecondaryStorage(p)

	if store.Name.String() != "nfs://10.147.28.6/export/home/sandbox/secondary" {
		t.Errorf("name: actual %s, expected nfs://10.147.28.6/export/home/sandbox/secondary", store.Name.String())
	}
}

func TestListSecondaryStagingStores(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listsecondarystagingstoreresponse": {
        "count": 1,
        "imagestore": [
            {
                "details": [],
                "id": "5b9d058b-492f-4959-8ac6-43cfbb33d267",
                "name": "nfs://10.147.28.6/export/home/sandbox/secondary",
                "protocol": "nfs",
                "providername": "NFS",
                "scope": "ZONE",
                "url": "nfs://10.147.28.6/export/home/sandbox/secondary",
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
	p := NewListSecondaryStagingStoresParameter()
	stores, _ := client.ListSecondaryStagingStores(p)

	if len(stores) != 1 {
		t.Errorf("length: actual %d, expected 1", len(stores))
	}

	if stores[0].Name.String() != "nfs://10.147.28.6/export/home/sandbox/secondary" {
		t.Errorf("name: actual %s, expected nfs://10.147.28.6/export/home/sandbox/secondary", stores[0].Name.String())
	}
}
