package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListDomains(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listdomainsresponse": {
        "count": 1,
        "domain": [
            {
                "haschild": false,
                "id": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "level": 0,
                "name": "ROOT",
                "path": "ROOT"
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
	p := NewListDomainsParameter()
	domains, _ := client.ListDomains(p)

	if len(domains) != 1 {
		t.Errorf("length: actual %d, expected 1", len(domains))
	}

	if domains[0].Name.String() != "ROOT" {
		t.Errorf("name: actual %s, expected ROOT", domains[0].Name.String())
	}
}

func TestListDomainChildren(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listdomainchildrenresponse": {
        "count": 3,
        "domain": [
            {
                "haschild": true,
                "id": "81c79460-4498-4f95-a2fa-30bfc363ea25",
                "level": 1,
                "name": "A",
                "parentdomainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "parentdomainname": "ROOT",
                "path": "ROOT/A"
            },
            {
                "haschild": false,
                "id": "322aa917-7e92-4e3d-9d6a-aa635c738b3a",
                "level": 1,
                "name": "B",
                "parentdomainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "parentdomainname": "ROOT",
                "path": "ROOT/B"
            },
            {
                "haschild": false,
                "id": "70790c51-a745-4cc9-84af-10aac0f55179",
                "level": 2,
                "name": "C",
                "parentdomainid": "81c79460-4498-4f95-a2fa-30bfc363ea25",
                "parentdomainname": "A",
                "path": "ROOT/A/C"
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
	p := NewListDomainChildrenParameter()
	p.IsRecursive.Set(true)
	domains, _ := client.ListDomainChildren(p)

	var level int64

	if len(domains) != 3 {
		t.Errorf("length: actual %d, expected 3", len(domains))
	}

	if domains[0].Path.String() != "ROOT/A" {
		t.Errorf("name: actual %s, expected ROOT/A", domains[0].Name.String())
	}

	if domains[0].ParentDomainName.String() != "ROOT" {
		t.Errorf("parentdomainname: actual %s, expected ROOT", domains[0].ParentDomainName.String())
	}

	level, _ = domains[0].Level.Int64()
	if level != 1 {
		t.Errorf("level: actual %s, expected 1", level)
	}

	if domains[2].Path.String() != "ROOT/A/C" {
		t.Errorf("name: actual %s, expected ROOT/A/C", domains[2].Name.String())
	}

	if domains[2].ParentDomainName.String() != "A" {
		t.Errorf("parentdomainname: actual %s, expected A", domains[2].ParentDomainName.String())
	}

	level, _ = domains[2].Level.Int64()
	if level != 2 {
		t.Errorf("level: actual %s, expected 2", level)
	}
}
