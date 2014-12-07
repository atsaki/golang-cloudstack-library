package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestCreateLoadBalancerRule(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		responses := map[string]string{
			"createLoadBalancerRule": `
{
    "createloadbalancerruleresponse": {
        "id": "dec16dce-5185-4334-8b28-bd052ac7cb95", 
        "jobid": "bd9108e4-54a6-4266-969b-7f69f6f2668c"
    }
}
`,
			"queryAsyncJobResult": `
{
    "queryasyncjobresultresponse": {
        "accountid": "00000000-0000-0000-0000-000000000000", 
        "cmd": "org.apache.cloudstack.api.command.user.loadbalancer.CreateLoadBalancerRuleCmd", 
        "created": "2014-12-06T22:44:31+0900", 
        "jobid": "bd9108e4-54a6-4266-969b-7f69f6f2668c", 
        "jobprocstatus": 0, 
        "jobresult": {
            "loadbalancer": {
                "account": "00000000000", 
                "algorithm": "roundrobin", 
                "cidrlist": "", 
                "domain": "00000000000", 
                "domainid": "00000000-0000-0000-0000-000000000000", 
                "id": "dec16dce-5185-4334-8b28-bd052ac7cb95", 
                "name": "lb1", 
                "networkid": "00000000-0000-0000-0000-000000000000", 
                "privateport": "80", 
                "protocol": "tcp", 
                "publicip": "1.1.1.1", 
                "publicipid": "00000000-0000-0000-0000-000000000000", 
                "publicport": "80", 
                "state": "Add", 
                "tags": [], 
                "zoneid": "00000000-0000-0000-0000-000000000000"
            }
        }, 
        "jobresultcode": 0, 
        "jobresulttype": "object", 
        "jobstatus": 1, 
        "userid": "00000000-0000-0000-0000-000000000000"
    }
}
`,
		}

		fmt.Fprintln(w, responses[r.FormValue("command")])
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewCreateLoadBalancerRuleParameter("roundrobin", "lb", 22, 22)
	p.PublicIpId.Set("00000000-0000-0000-0000-000000000000")
	lb, err := client.CreateLoadBalancerRule(p)
	if err != nil {
		t.Errorf(err.Error())
	}
	if lb.Algorithm.String() != "roundrobin" {
		t.Errorf("algorithm actual %s, expected roundrobin", lb.Algorithm.String())
	}
}

func TestCreateLoadBalancerRuleCantDefineIpOwner(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		responses := map[string]string{
			"createLoadBalancerRule": `
{
    "createloadbalancerruleresponse": {
        "cserrorcode": 4350, 
        "errorcode": 431, 
        "errortext": "Can't define IP owner. Either specify account/domainId or publicIpId", 
        "uuidList": []
    }
}
`,
		}

		fmt.Fprintln(w, responses[r.FormValue("command")])
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewCreateLoadBalancerRuleParameter("roundrobin", "lb", 22, 22)
	_, err := client.CreateLoadBalancerRule(p)
	expected := "Can't define IP owner. Either specify account/domainId or publicIpId"
	if err.Error() != expected {
		t.Errorf("error message: actual: \"%s\", expected \"%s\"", err.Error(), expected)
	}
}
