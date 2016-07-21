package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListPods(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listpodsresponse": {
        "count": 1,
        "pod": [
            {
                "allocationstate": "Enabled",
                "endip": "172.16.15.200",
                "gateway": "172.16.15.1",
                "id": "64ca09fe-9269-462d-8b4e-e38ef543986e",
                "name": "POD0",
                "netmask": "255.255.255.0",
                "startip": "172.16.15.2",
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
	p := NewListPodsParameter()
	pods, _ := client.ListPods(p)

	if len(pods) != 1 {
		t.Errorf("length: actual %d, expected 1", len(pods))
	}

	if pods[0].Name.String() != "POD0" {
		t.Errorf("name: actual %s, expected POD0", pods[0].Name.String())
	}

	if pods[0].StartIp.String() != "172.16.15.2" {
		t.Errorf("startip: actual %s, expected 172.16.15.2", pods[0].StartIp.String())
	}
}

func TestListDedicatedPods(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listdedicatedpodsresponse": {
        "count": 1,
        "dedicatedpod": [
            {
                "affinitygroupid": "26ff4a0a-d7a2-4b4e-9fda-83d24ab61c14",
                "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "id": "49f69eeb-9f27-47a1-8c44-9609978af5f9",
                "podid": "64ca09fe-9269-462d-8b4e-e38ef543986e",
                "podname": "POD0"
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
	p := NewListDedicatedPodsParameter()
	pods, _ := client.ListDedicatedPods(p)

	if len(pods) != 1 {
		t.Errorf("length: actual %d, expected 1", len(pods))
	}

	if pods[0].Id.String() != "49f69eeb-9f27-47a1-8c44-9609978af5f9" {
		t.Errorf("id: actual %d, expected 49f69eeb-9f27-47a1-8c44-9609978af5f9", pods[0].Id)
	}
}
