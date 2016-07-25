package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListStoragePools(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "liststoragepoolsresponse": {
        "count": 2,
        "storagepool": [
            {
                "clusterid": "b980b10b-92f4-45f5-98c8-68844d4d8287",
                "clustername": "C0",
                "created": "2015-09-20T06:27:43+0000",
                "disksizeallocated": 200,
                "disksizetotal": 1099511627776,
                "disksizeused": 0,
                "id": "cd360613-e1cf-3a32-b8ef-d07bad8dd92b",
                "ipaddress": "10.147.28.6",
                "name": "PS1",
                "overprovisionfactor": "2.0",
                "path": "/export/home/sandbox/primary1",
                "podid": "64ca09fe-9269-462d-8b4e-e38ef543986e",
                "podname": "POD0",
                "scope": "CLUSTER",
                "state": "Up",
                "storagecapabilities": {
                    "VOLUME_SNAPSHOT_QUIESCEVM": "false"
                },
                "type": "NetworkFilesystem",
                "zoneid": "f902785c-9dda-45c0-afc3-f1a29d202b2e",
                "zonename": "Sandbox-simulator"
            },
            {
                "clusterid": "b980b10b-92f4-45f5-98c8-68844d4d8287",
                "clustername": "C0",
                "created": "2015-09-20T06:27:43+0000",
                "disksizeallocated": 100,
                "disksizetotal": 1099511627776,
                "disksizeused": 0,
                "id": "7c07ec9b-a3c6-3466-ab5a-f5669ead0b22",
                "ipaddress": "10.147.28.6",
                "name": "PS0",
                "overprovisionfactor": "2.0",
                "path": "/export/home/sandbox/primary0",
                "podid": "64ca09fe-9269-462d-8b4e-e38ef543986e",
                "podname": "POD0",
                "scope": "CLUSTER",
                "state": "Up",
                "storagecapabilities": {
                    "VOLUME_SNAPSHOT_QUIESCEVM": "false"
                },
                "type": "NetworkFilesystem",
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
	p := NewListStoragePoolsParameter()
	p.ClusterId.Set("b980b10b-92f4-45f5-98c8-68844d4d8287")
	pools, _ := client.ListStoragePools(p)

	if len(pools) != 2 {
		t.Errorf("length: actual %d, expected 2", len(pools))
	}

	if pools[0].Name.String() != "PS1" {
		t.Errorf("name: actual %s, expected PS1", pools[0].Name.String())
	}

	if pools[0].StorageCapabilities["VOLUME_SNAPSHOT_QUIESCEVM"] != "false" {
		t.Errorf("VOLUME_SNAPSHOT_QUIESCEVM: actual %s, expected false", pools[0].StorageCapabilities["VOLUME_SNAPSHOT_QUIESCEVM"])
	}
}
