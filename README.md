golang-cloudstack-library
=========================

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
```

This returns JSON response as follows.

```json
{
    "count": 1,
    "zone": [
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
}
```

Projects Using This Library
===========================

* lockgate
  * https://github.com/atsaki/lockgate


