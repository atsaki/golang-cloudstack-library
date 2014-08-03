package cloudstack

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"testing"
)

const (
	apikey    = ""
	secretkey = ""
	username  = "admin"
	password  = "password"
)

var (
	endpoint = url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "client/api",
	}
)

func TestListZones(t *testing.T) {
	client, err := NewClient(endpoint, apikey, secretkey, username, password)
	if err != nil {
		t.FailNow()
	}
	params := ListZonesParameter{}
	params.SetId("1")
	r, err := client.ListZones(params)
	if err != nil {
		t.FailNow()
	}
	b, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		t.FailNow()
	}
	fmt.Println(string(b))
}

func TestListVirtualMachines(t *testing.T) {
	client, err := NewClient(endpoint, apikey, secretkey, username, password)
	if err != nil {
		t.FailNow()
	}
	params := ListVirtualMachinesParameter{}
	r, err := client.ListVirtualMachines(params)
	if err != nil {
		t.FailNow()
	}
	b, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		t.FailNow()
	}
	fmt.Println(string(b))
}

func TestLogIn(t *testing.T) {
	client, err := NewClient(endpoint, apikey, secretkey, username, password)
	if err != nil {
		log.Fatal(err)
	}
	client.LogIn()
	t.Errorf("error")
}
