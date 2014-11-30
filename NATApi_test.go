package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestDisableStaticNatNotAssociatedIpAddress(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		responses := map[string]string{
			"disableStaticNat": `
{
    "disablestaticnatresponse": {
        "jobid": "6c9e2fe9-48e5-4e41-a0fb-6ef24dcd6bea"
    }
}
`,
			"queryAsyncJobResult": `
{
    "queryasyncjobresultresponse": {
        "accountid": "00000000-0000-0000-0000-000000000000", 
        "cmd": "org.apache.cloudstack.api.command.user.nat.DisableStaticNatCmd", 
        "created": "2014-11-30T17:29:41+0900", 
        "jobid": "6c9e2fe9-48e5-4e41-a0fb-6ef24dcd6bea", 
        "jobprocstatus": 0, 
        "jobresult": {
            "errorcode": 530, 
            "errortext": "Specified IP address id is not associated with any vm Id"
        }, 
        "jobresultcode": 530, 
        "jobresulttype": "object", 
        "jobstatus": 2, 
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
	p := NewDisableStaticNatParameter("f66517fa-3ec0-43c7-8843-5e68e5a70751")
	_, err := client.DisableStaticNat(p)
	if err == nil || err.Error() != "Specified IP address id is not associated with any vm Id" {
		t.Errorf("Failed to raise Error \"Specified IP address id is not associated with any vm Id\"")
	}
}
