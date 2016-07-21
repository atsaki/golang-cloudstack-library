package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListClusters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listclustersresponse": {
        "cluster": [
            {
                "allocationstate": "Enabled",
                "clustertype": "CloudManaged",
                "cpuovercommitratio": "1.0",
                "hypervisortype": "Simulator",
                "id": "b980b10b-92f4-45f5-98c8-68844d4d8287",
                "managedstate": "Managed",
                "memoryovercommitratio": "1.0",
                "name": "C0",
                "podid": "64ca09fe-9269-462d-8b4e-e38ef543986e",
                "podname": "POD0",
                "zoneid": "f902785c-9dda-45c0-afc3-f1a29d202b2e",
                "zonename": "Sandbox-simulator"
            },
            {
                "allocationstate": "Enabled",
                "clustertype": "CloudManaged",
                "cpuovercommitratio": "1.0",
                "hypervisortype": "Simulator",
                "id": "11370fb1-a9b6-4c3a-b8d0-c7b8ae378c79",
                "managedstate": "Managed",
                "memoryovercommitratio": "1.0",
                "name": "C1",
                "podid": "64ca09fe-9269-462d-8b4e-e38ef543986e",
                "podname": "POD0",
                "zoneid": "f902785c-9dda-45c0-afc3-f1a29d202b2e",
                "zonename": "Sandbox-simulator"
            }
        ],
        "count": 2
    }
}
		`
		fmt.Fprintln(w, resp)
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewListClustersParameter()
	clusters, _ := client.ListClusters(p)

	if len(clusters) != 2 {
		t.Errorf("length: actual %d, expected 2", len(clusters))
	}

	if clusters[0].Name.String() != "C0" {
		t.Errorf("name: actual %s, expected C0", clusters[0].Name.String())
	}

	if clusters[0].PodName.String() != "POD0" {
		t.Errorf("podname: actual %s, expected POD0", clusters[0].PodName.String())
	}

	if clusters[1].ZoneName.String() != "Sandbox-simulator" {
		t.Errorf("zonename: actual %s, expected Sandbox-simulator", clusters[1].ZoneName.String())
	}
}

func TestListDedicatedClusters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listdedicatedclustersresponse": {
        "count": 1,
        "dedicatedcluster": [
            {
                "affinitygroupid": "22058c36-b875-4c2d-b82f-6269bb85e250",
                "clusterid": "b980b10b-92f4-45f5-98c8-68844d4d8287",
                "clustername": "C0",
                "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "id": "b3300b69-7cb5-44d0-a903-d106e0e32ea2"
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
	p := NewListDedicatedClustersParameter()
	clusters, _ := client.ListDedicatedClusters(p)

	if len(clusters) != 1 {
		t.Errorf("length: actual %d, expected 1", len(clusters))
	}

	if clusters[0].Id.String() != "b3300b69-7cb5-44d0-a903-d106e0e32ea2" {
		t.Errorf("id: actual %d, expected b3300b69-7cb5-44d0-a903-d106e0e32ea2", clusters[0].Id)
	}
}
