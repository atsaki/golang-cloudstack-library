package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/atsaki/golang-cloudstack-library"
)

func main() {

	log.SetOutput(ioutil.Discard)

	endpoint := url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "client/api",
	}
	apikey := ""
	secretkey := ""
	username := "admin"
	password := "password"

	client, _ := cloudstack.NewClient(endpoint, apikey, secretkey, username, password)

	params := cloudstack.ListZonesParameter{}
	params.SetName("zone1")

	r, _ := client.ListZones(params)
	b, _ := json.MarshalIndent(r, "", "    ")

	fmt.Println(string(b))
}
