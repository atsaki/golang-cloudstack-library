package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListPortForwardingRulesEmptyResult(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		responses := map[string]string{
			"listPortForwardingRules": `
			{ "listportforwardingrulesresponse" : { } }
`,
		}

		fmt.Fprintln(w, responses[r.FormValue("command")])
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewListPortForwardingRulesParameter()
	pfs, err := client.ListPortForwardingRules(p)
	if err != nil {
		t.Errorf(err.Error())
	}
	if len(pfs) != 0 {
		t.Errorf("length actual %d, expected 0", len(pfs))
	}
}
