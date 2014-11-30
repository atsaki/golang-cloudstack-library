package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestAssociateIpAddress(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responses := map[string]string{
			"associateIpAddress": `
		{
		    "associateipaddressresponse": {
		        "id": "2c7fb564-38e0-44db-821f-6486a544e423",
		        "jobid": "7bc6769c-bbeb-4de2-ace2-e57f702af44a"
		    }
		}
				`,
			"queryAsyncJobResult": `
		{
		    "queryasyncjobresultresponse": {
		        "accountid": "6c9ce4be-ed0b-477b-aa76-f0cb1b7a200b",
		        "cmd": "org.apache.cloudstack.api.command.user.address.AssociateIPAddrCmd",
		        "created": "2014-11-23T15:42:26+0900",
		        "jobid": "7bc6769c-bbeb-4de2-ace2-e57f702af44a",
		        "jobprocstatus": 0,
		        "jobresult": {
		            "ipaddress": {
		                "account": "account1",
		                "allocated": "2014-11-23T15:42:26+0900",
		                "associatednetworkid": "7eef65ed-b952-46c5-bea4-25d106c37a3b",
		                "associatednetworkname": "network1",
		                "domain": "domain1",
		                "domainid": "9d9dddc3-ce38-494b-973e-b80519d76b22",
		                "forvirtualnetwork": true,
		                "id": "2c7fb564-38e0-44db-821f-6486a544e423",
		                "ipaddress": "1.1.1.1",
		                "isportable": false,
		                "issourcenat": false,
		                "isstaticnat": false,
		                "issystem": false,
		                "networkid": "79132c74-fe77-4bd5-9915-ce7c577fb95f",
		                "physicalnetworkid": "4a00ce42-6a30-4494-afdd-3531d883237b",
		                "state": "Allocating",
		                "tags": [],
		                "zoneid": "a117e75f-d02e-4074-806d-889c61261394",
		                "zonename": "tesla"
		            }
		        },
		        "jobresultcode": 0,
		        "jobresulttype": "object",
		        "jobstatus": 1,
		        "userid": "c8be3f37-1175-475b-8f67-08bb33d6f6ea"
		    }
		}
		`,
		}

		fmt.Fprintln(w, responses[r.FormValue("command")])
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewAssociateIpAddressParameter()
	p.ZoneId.Set("a117e75f-d02e-4074-806d-889c61261394")
	ip, err := client.AssociateIpAddress(p)
	if err != nil {
		t.Errorf(err.Error())
	}

	if ip.IpAddress.String() != "1.1.1.1" {
		t.Errorf("ipaddress: actual %s, expected 1.1.1.1", ip.IpAddress.String())
	}

	if ip.client.EndPoint != endpoint {
		t.Errorf("endpoint: actual %v, expected %v", ip.client.EndPoint, endpoint)
	}
}
