package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListTags(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listtagsresponse": {
        "count": 2,
        "tag": [
            {
                "account": "admin",
                "domain": "ROOT",
                "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "key": "key0",
                "resourceid": "869cad78-2232-4c4e-8018-f38462cbc1df",
                "resourcetype": "UserVm",
                "value": "value0"
            },
            {
                "account": "admin",
                "domain": "ROOT",
                "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "key": "key1",
                "resourceid": "869cad78-2232-4c4e-8018-f38462cbc1df",
                "resourcetype": "UserVm",
                "value": "value1"
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
	p := NewListTagsParameter()
	tags, _ := client.ListTags(p)

	if len(tags) != 2 {
		t.Errorf("length: actual %d, expected 2", len(tags))
	}

	if tags[0].Key.String() != "key0" {
		t.Errorf("key: actual %s, expected key0", tags[0].Key.String())
	}

	if tags[0].Value.String() != "value0" {
		t.Errorf("value: actual %s, expected value0", tags[0].Value.String())
	}
}

func TestListStorageTags(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "liststoragetagsresponse": {
        "count": 1,
        "storagetag": [
            {
                "name": "tag01",
                "poolid": 1
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
	p := NewListStorageTagsParameter()
	tags, _ := client.ListStorageTags(p)

	if len(tags) != 1 {
		t.Errorf("length: actual %d, expected 1", len(tags))
	}

	if tags[0].Name.String() != "tag01" {
		t.Errorf("key: actual %s, expected \"tag01\"", tags[0].Name.String())
	}
}
