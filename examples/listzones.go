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

	endpoint, _ := url.Parse("http://localhost:8080/client/api")
	apikey := ""
	secretkey := ""
	username := "admin"
	password := "password"

	client, _ := cloudstack.NewClient(endpoint, apikey, secretkey, username, password)

	params := cloudstack.NewListZonesParameter()
	params.Name.Set("zone1")

	zones, _ := client.ListZones(params)
	b, _ := json.MarshalIndent(zones, "", "    ")

	fmt.Println("Count:", len(zones))
	fmt.Println(string(b))
}
