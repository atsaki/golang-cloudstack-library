package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListNetworks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		responses := map[string]string{
			"listNetworks": `
{
    "listnetworksresponse": {
        "count": 1, 
        "network": [
            {
                "account": "account1", 
                "acltype": "Account", 
                "broadcastdomaintype": "Vlan", 
                "canusefordeploy": true, 
                "cidr": "10.3.0.0/22", 
                "displaynetwork": true, 
                "displaytext": "network1", 
                "dns1": "8.8.8.8", 
                "dns2": "8.8.8.8", 
                "domain": "domain1", 
                "domainid": "00000000-0000-0000-0000-000000000000", 
                "gateway": "10.3.0.1", 
                "id": "00000000-0000-0000-0000-000000000000", 
                "ispersistent": true, 
                "issystem": false, 
                "name": "network1", 
                "netmask": "255.255.248.0", 
                "networkcidr": "10.3.0.0/21", 
                "networkdomain": "examplecloud.internal", 
                "networkofferingavailability": "Required", 
                "networkofferingconservemode": true, 
                "networkofferingdisplaytext": "Offering for Isolated networks with Source Nat service enabled", 
                "networkofferingid": "00000000-0000-0000-0000-000000000000", 
                "networkofferingname": "DefaultIsolatedNetworkOffering", 
                "physicalnetworkid": "00000000-0000-0000-0000-000000000000", 
                "related": "00000000-0000-0000-0000-000000000000", 
                "reservediprange": "10.3.3.255-10.3.7.254", 
                "restartrequired": false, 
                "service": [
                    {
                        "capability": [
                            {
                                "canchooseservicecapability": false, 
                                "name": "SupportedTrafficDirection", 
                                "value": "ingress, egress"
                            }, 
                            {
                                "canchooseservicecapability": false, 
                                "name": "SupportedEgressProtocols", 
                                "value": "tcp,udp,icmp, all"
                            }, 
                            {
                                "canchooseservicecapability": false, 
                                "name": "SupportedProtocols", 
                                "value": "tcp,udp,icmp"
                            }, 
                            {
                                "canchooseservicecapability": false, 
                                "name": "TrafficStatistics", 
                                "value": "per public ip"
                            }, 
                            {
                                "canchooseservicecapability": false, 
                                "name": "MultipleIps", 
                                "value": "true"
                            }
                        ], 
                        "name": "Firewall"
                    }, 
                    {
                        "capability": [
                            {
                                "canchooseservicecapability": false, 
                                "name": "SupportedProtocols", 
                                "value": "tcp, udp"
                            }, 
                            {
                                "canchooseservicecapability": false, 
                                "name": "LbSchemes", 
                                "value": "Public"
                            }, 
                            {
                                "canchooseservicecapability": false, 
                                "name": "SupportedStickinessMethods", 
                                "value": "[{\"methodname\":\"LbCookie\",\"paramlist\":[{\"paramname\":\"cookie-name\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"mode\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"nocache\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"indirect\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"postonly\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"domain\",\"required\":false,\"isflag\":false,\"description\":\" \"}],\"description\":\"This is loadbalancer cookie based stickiness method.\"},{\"methodname\":\"AppCookie\",\"paramlist\":[{\"paramname\":\"cookie-name\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"length\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"holdtime\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"request-learn\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"prefix\",\"required\":false,\"isflag\":true,\"description\":\" \"},{\"paramname\":\"mode\",\"required\":false,\"isflag\":false,\"description\":\" \"}],\"description\":\"This is App session based sticky method. Define session stickiness on an existing application cookie. It can be used only for a specific http traffic\"},{\"methodname\":\"SourceBased\",\"paramlist\":[{\"paramname\":\"tablesize\",\"required\":false,\"isflag\":false,\"description\":\" \"},{\"paramname\":\"expire\",\"required\":false,\"isflag\":false,\"description\":\" \"}],\"description\":\"This is source based Stickiness method, it can be used for any type of protocol.\"}]"
                            }, 
                            {
                                "canchooseservicecapability": false, 
                                "name": "SupportedLbAlgorithms", 
                                "value": "roundrobin,leastconn,source"
                            }, 
                            {
                                "canchooseservicecapability": false, 
                                "name": "SupportedLBIsolation", 
                                "value": "dedicated"
                            }
                        ], 
                        "name": "Lb"
                    }, 
                    {
                        "capability": [
                            {
                                "canchooseservicecapability": false, 
                                "name": "SupportedSourceNatTypes", 
                                "value": "peraccount"
                            }, 
                            {
                                "canchooseservicecapability": false, 
                                "name": "RedundantRouter", 
                                "value": "true"
                            }
                        ], 
                        "name": "SourceNat"
                    }, 
                    {
                        "capability": [
                            {
                                "canchooseservicecapability": false, 
                                "name": "SupportedVpnTypes", 
                                "value": "pptp,l2tp,ipsec"
                            }, 
                            {
                                "canchooseservicecapability": false, 
                                "name": "VpnTypes", 
                                "value": "removeaccessvpn"
                            }
                        ], 
                        "name": "Vpn"
                    }, 
                    {
                        "name": "UserData"
                    }, 
                    {
                        "name": "StaticNat"
                    }, 
                    {
                        "capability": [
                            {
                                "canchooseservicecapability": false, 
                                "name": "DhcpAccrossMultipleSubnets", 
                                "value": "true"
                            }
                        ], 
                        "name": "Dhcp"
                    }, 
                    {
                        "name": "PortForwarding"
                    }, 
                    {
                        "capability": [
                            {
                                "canchooseservicecapability": false, 
                                "name": "AllowDnsSuffixModification", 
                                "value": "true"
                            }
                        ], 
                        "name": "Dns"
                    }
                ], 
                "specifyipranges": false, 
                "state": "Implemented", 
                "tags": [], 
                "traffictype": "Guest", 
                "type": "Isolated", 
                "zoneid": "00000000-0000-0000-0000-000000000000", 
                "zonename": "zone1"
            }
        ]
    }
}
			`,
		}

		fmt.Fprintln(w, responses[r.FormValue("command")])
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewListNetworksParameter()
	nws, err := client.ListNetworks(p)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(nws) != 1 {
		t.Errorf("length actual %d, expected 1", len(nws))
	}

	if nws[0].ZoneName.String() != "zone1" {
		t.Errorf("zonename actual %s, expected zone1", nws[0].ZoneName.String())
	}

}
