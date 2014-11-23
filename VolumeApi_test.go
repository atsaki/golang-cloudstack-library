package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestDeleteVolume(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		responses := map[string]string{
			"deleteVolume": `{"deletevolumeresponse":{"success":"true"}}`,
		}

		fmt.Fprintln(w, responses[r.FormValue("command")])
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewDeleteVolumeParameter("0728020b-cdee-4f63-bb96-de831c384a73")
	result, err := client.DeleteVolume(p)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !result.Success.Bool() {
		t.Errorf("success: actual true, expected true")
	}

}
