package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/atsaki/golang-cloudstack-library"
)

func TestKeyAuthentication(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "")
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := cloudstack.NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	client.Request("listZones", map[string]interface{}{})
}

func TestPasswordAuthentication(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "")
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := cloudstack.NewClient(endpoint, "", "", "admin", "password")
	client.Request("listZones", map[string]interface{}{})
}
