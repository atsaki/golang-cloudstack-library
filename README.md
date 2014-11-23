golang-cloudstack-library
=========================

[![Build Status](https://travis-ci.org/atsaki/golang-cloudstack-library.svg?branch=master)](https://travis-ci.org/atsaki/golang-cloudstack-library)

Go language library for CloudStack API

Installation
============

```
go get github.com/atsaki/golang-cloudstack-library
```

Usage
=====

```go
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
```

This returns JSON response as follows.

```
Count: 1
[
    {
        "allocationstate": "Enabled",
        "capacity": null,
        "description": null,
        "dhcpprovider": "VirtualRouter",
        "displaytext": null,
        "dns1": "8.8.8.8",
        "dns2": null,
        "domain": null,
        "domainid": null,
        "domainname": null,
        "guestcidraddress": "10.1.1.0/24",
        "id": "7fdea428-8228-4dfd-850b-d60eed5004a4",
        "internaldns1": "8.8.4.4",
        "internaldns2": null,
        "ip6dns1": null,
        "ip6dns2": null,
        "localstorageenabled": false,
        "name": "zone1",
        "networktype": "Advanced",
        "resourcedetails": null,
        "securitygroupsenabled": false,
        "tags": [],
        "vlan": null,
        "zonetoken": "c263e076-5887-355b-afe9-cc988a1ccac7"
    }
]
```

Projects Using This Library
===========================

* terraform-cloudstack-provider
  * https://github.com/atsaki/terraform-provider-cloudstack
* lockgate
  * https://github.com/atsaki/lockgate


