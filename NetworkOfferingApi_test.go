package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListNetworkOfferings(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		responses := map[string]string{
			"listNetworkOfferings": `
{
    "listnetworkofferingsresponse": {
        "count": 1, 
        "networkoffering": [
            {
                "availability": "Required", 
                "conservemode": true, 
                "displaytext": "Offering for Isolated networks with Source Nat service enabled", 
                "egressdefaultpolicy": false, 
                "forvpc": false, 
                "guestiptype": "Isolated", 
                "id": "00000000-0000-0000-0000-000000000000", 
                "isdefault": true, 
                "ispersistent": true, 
                "name": "DefaultIsolatedNetworkOffering", 
                "networkrate": 2000, 
                "service": [
                    {
                        "name": "Firewall", 
                        "provider": [
                            {
                                "name": "VirtualRouter"
                            }
                        ]
                    }, 
                    {
                        "capability": [
                            {
                                "canchooseservicecapability": false, 
                                "name": "SupportedLBIsolation", 
                                "value": "dedicated"
                            }, 
                            {
                                "canchooseservicecapability": false, 
                                "name": "ElasticLb", 
                                "value": "false"
                            }, 
                            {
                                "canchooseservicecapability": false, 
                                "name": "InlineMode", 
                                "value": "false"
                            }
                        ], 
                        "name": "Lb", 
                        "provider": [
                            {
                                "name": "VirtualRouter"
                            }
                        ]
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
                        "name": "SourceNat", 
                        "provider": [
                            {
                                "name": "VirtualRouter"
                            }
                        ]
                    }, 
                    {
                        "name": "Vpn", 
                        "provider": [
                            {
                                "name": "VirtualRouter"
                            }
                        ]
                    }, 
                    {
                        "name": "UserData", 
                        "provider": [
                            {
                                "name": "VirtualRouter"
                            }
                        ]
                    }, 
                    {
                        "capability": [
                            {
                                "canchooseservicecapability": false, 
                                "name": "ElasticIp", 
                                "value": "false"
                            }, 
                            {
                                "canchooseservicecapability": false, 
                                "name": "AssociatePublicIP", 
                                "value": "true"
                            }
                        ], 
                        "name": "StaticNat", 
                        "provider": [
                            {
                                "name": "VirtualRouter"
                            }
                        ]
                    }, 
                    {
                        "name": "Dhcp", 
                        "provider": [
                            {
                                "name": "VirtualRouter"
                            }
                        ]
                    }, 
                    {
                        "name": "PortForwarding", 
                        "provider": [
                            {
                                "name": "VirtualRouter"
                            }
                        ]
                    }, 
                    {
                        "name": "Dns", 
                        "provider": [
                            {
                                "name": "VirtualRouter"
                            }
                        ]
                    }
                ], 
                "serviceofferingid": "00000000-0000-0000-0000-000000000000", 
                "specifyipranges": false, 
                "specifyvlan": false, 
                "state": "Enabled", 
                "traffictype": "Guest"
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
	p := NewListNetworkOfferingsParameter()
	nos, err := client.ListNetworkOfferings(p)
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(nos) != 1 {
		t.Errorf("length actual %d, expected 1", len(nos))
	}

	if nos[0].Name.String() != "DefaultIsolatedNetworkOffering" {
		t.Errorf("name actual %s, expected DefaultIsolatedNetworkOffering",
			nos[0].Name.String())
	}

}
