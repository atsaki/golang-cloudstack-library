package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListHosts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listhostsresponse": {
        "count": 4,
        "host": [
            {
                "created": "2016-07-14T00:22:54+0000",
                "events": "AgentDisconnected; ManagementServerDown; HostDown; Remove; PingTimeout; WaitedTooLong; AgentConnected; Ping",
                "id": "d430cbe4-cea0-49e1-bcf9-16375c7cb3eb",
                "ipaddress": "172.16.15.38",
                "islocalstorageactive": false,
                "lastpinged": "1970-01-17T14:20:38+0000",
                "name": "s-1-QA",
                "podid": "64ca09fe-9269-462d-8b4e-e38ef543986e",
                "podname": "POD0",
                "resourcestate": "Enabled",
                "state": "Disconnected",
                "type": "SecondaryStorageVM",
                "version": "4.5.2",
                "zoneid": "f902785c-9dda-45c0-afc3-f1a29d202b2e",
                "zonename": "Sandbox-simulator"
            },
            {
                "capabilities": "hvm",
                "clusterid": "11370fb1-a9b6-4c3a-b8d0-c7b8ae378c79",
                "clustername": "C1",
                "clustertype": "CloudManaged",
                "cpuallocated": "0%",
                "cpunumber": 4,
                "cpuspeed": 8000,
                "cpuused": "0%",
                "cpuwithoverprovisioning": "32000.0",
                "created": "2015-09-20T06:27:43+0000",
                "disconnected": "2016-07-14T00:22:13+0000",
                "events": "ManagementServerDown; AgentDisconnected; HostDown; Remove; PingTimeout; ShutdownRequested; AgentConnected; Ping; StartAgentRebalance",
                "hahost": false,
                "hosttags": "TEST02,TEST03",
                "hypervisor": "Simulator",
                "id": "4d12f1f6-595c-4c5d-a7eb-30693d289ecb",
                "ipaddress": "172.16.15.7",
                "islocalstorageactive": false,
                "lastpinged": "1970-01-17T14:28:12+0000",
                "managementserverid": 2485377892354,
                "memoryallocated": 0,
                "memorytotal": 8589934592,
                "memoryused": 0,
                "name": "SimulatedAgent.ac91f9eb-aab7-4e69-9695-af725b2e8b29",
                "networkkbsread": 32768,
                "networkkbswrite": 16384,
                "podid": "64ca09fe-9269-462d-8b4e-e38ef543986e",
                "podname": "POD0",
                "resourcestate": "Enabled",
                "state": "Up",
                "type": "Routing",
                "version": "4.5.2",
                "zoneid": "f902785c-9dda-45c0-afc3-f1a29d202b2e",
                "zonename": "Sandbox-simulator"
            },
            {
                "capabilities": "hvm",
                "clusterid": "b980b10b-92f4-45f5-98c8-68844d4d8287",
                "clustername": "C0",
                "clustertype": "CloudManaged",
                "cpuallocated": "1.56%",
                "cpunumber": 4,
                "cpuspeed": 8000,
                "cpuused": "0%",
                "cpuwithoverprovisioning": "32000.0",
                "created": "2015-09-20T06:26:43+0000",
                "disconnected": "2016-07-14T00:22:13+0000",
                "events": "ManagementServerDown; AgentDisconnected; HostDown; Remove; PingTimeout; ShutdownRequested; AgentConnected; Ping; StartAgentRebalance",
                "hahost": false,
                "hosttags": "TEST01",
                "hypervisor": "Simulator",
                "id": "a2a92fbc-5525-43b9-9032-6c42b0a9f537",
                "ipaddress": "172.16.15.0",
                "islocalstorageactive": false,
                "lastpinged": "1970-01-17T14:28:12+0000",
                "managementserverid": 2485377892354,
                "memoryallocated": 536870912,
                "memorytotal": 8589934592,
                "memoryused": 0,
                "name": "SimulatedAgent.ec6e61cc-96a7-4e6c-a552-61fdb10da58c",
                "networkkbsread": 32768,
                "networkkbswrite": 16384,
                "podid": "64ca09fe-9269-462d-8b4e-e38ef543986e",
                "podname": "POD0",
                "resourcestate": "Enabled",
                "state": "Up",
                "type": "Routing",
                "version": "4.5.2",
                "zoneid": "f902785c-9dda-45c0-afc3-f1a29d202b2e",
                "zonename": "Sandbox-simulator"
            },
            {
                "capabilities": "hvm",
                "clusterid": "b980b10b-92f4-45f5-98c8-68844d4d8287",
                "clustername": "C0",
                "clustertype": "CloudManaged",
                "cpuallocated": "1.56%",
                "cpunumber": 4,
                "cpuspeed": 8000,
                "cpuused": "0%",
                "cpuwithoverprovisioning": "32000.0",
                "created": "2015-09-20T06:26:42+0000",
                "disconnected": "2016-07-14T00:22:13+0000",
                "events": "ManagementServerDown; AgentDisconnected; HostDown; Remove; PingTimeout; ShutdownRequested; AgentConnected; Ping; StartAgentRebalance",
                "hahost": false,
                "hypervisor": "Simulator",
                "id": "474c30c5-4c15-4ef1-a42b-e383e13b3183",
                "ipaddress": "172.16.15.10",
                "islocalstorageactive": false,
                "lastpinged": "1970-01-17T14:28:12+0000",
                "managementserverid": 2485377892354,
                "memoryallocated": 1073741824,
                "memorytotal": 8589934592,
                "memoryused": 0,
                "name": "SimulatedAgent.516d3d06-2857-4205-91d3-41c54f0af8fe",
                "networkkbsread": 32768,
                "networkkbswrite": 16384,
                "podid": "64ca09fe-9269-462d-8b4e-e38ef543986e",
                "podname": "POD0",
                "resourcestate": "Enabled",
                "state": "Up",
                "type": "Routing",
                "version": "4.5.2",
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
	p := NewListHostsParameter()
	hosts, _ := client.ListHosts(p)

	if len(hosts) != 4 {
		t.Errorf("length: actual %d, expected 4", len(hosts))
	}

	if hosts[0].Name.String() != "s-1-QA" {
		t.Errorf("name: actual %s, expected s-1-QA", hosts[0].Name.String())
	}
	if hosts[1].Name.String() != "SimulatedAgent.ac91f9eb-aab7-4e69-9695-af725b2e8b29" {
		t.Errorf("name: actual %s, expected SimulatedAgent.ac91f9eb-aab7-4e69-9695-af725b2e8b29", hosts[1].Name.String())
	}

	if hosts[2].State.String() != "Up" {
		t.Errorf("state: actual %s, expected Up", hosts[2].State.String())
	}
}

func TestListHostTags(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listhosttagsresponse": {
        "count": 3,
        "hosttag": [
            {
                "hostid": 2,
                "name": "TEST01"
            },
            {
                "hostid": 3,
                "name": "TEST02"
            },
            {
                "hostid": 3,
                "name": "TEST03"
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
	p := NewListHostTagsParameter()
	tags, _ := client.ListHostTags(p)

	if len(tags) != 3 {
		t.Errorf("length: actual %d, expected 3", len(tags))
	}

	if tags[0].HostId.String() != "2" {
		t.Errorf("hostid: actual %d, expected 2", tags[0].HostId)
	}
	if tags[1].Name.String() != "TEST02" {
		t.Errorf("name: actual %s, expected TEST02", tags[1].Name.String())
	}
}

func TestListDedicatedHosts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listdedicatedhostsresponse": {
        "count": 1,
        "dedicatedhost": [
            {
                "affinitygroupid": "d33549e4-d5b9-4089-9141-04eaf37fe4f0",
                "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "hostid": "474c30c5-4c15-4ef1-a42b-e383e13b3183",
                "hostname": "SimulatedAgent.516d3d06-2857-4205-91d3-41c54f0af8fe",
                "id": "56e6b6b7-dd2f-4a45-9169-5cb3e35ff554"
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
	p := NewListDedicatedHostsParameter()
	hosts, _ := client.ListDedicatedHosts(p)

	if len(hosts) != 1 {
		t.Errorf("length: actual %d, expected 1", len(hosts))
	}

	if hosts[0].Id.String() != "56e6b6b7-dd2f-4a45-9169-5cb3e35ff554" {
		t.Errorf("id: actual %d, expected 56e6b6b7-dd2f-4a45-9169-5cb3e35ff554", hosts[0].Id)
	}
}
