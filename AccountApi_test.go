package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListAccounts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listaccountsresponse": {
        "account": [
            {
                "accounttype": 1,
                "cpuavailable": "Unlimited",
                "cpulimit": "Unlimited",
                "cputotal": 1,
                "domain": "ROOT",
                "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "groups": [],
                "id": "62f61a73-5f5a-11e5-bc86-0242ac11180a",
                "ipavailable": "Unlimited",
                "iplimit": "Unlimited",
                "iptotal": 1,
                "isdefault": true,
                "memoryavailable": "Unlimited",
                "memorylimit": "Unlimited",
                "memorytotal": 512,
                "name": "admin",
                "networkavailable": "Unlimited",
                "networklimit": "Unlimited",
                "networktotal": 1,
                "primarystorageavailable": "Unlimited",
                "primarystoragelimit": "Unlimited",
                "primarystoragetotal": 2,
                "projectavailable": "Unlimited",
                "projectlimit": "Unlimited",
                "projecttotal": 1,
                "receivedbytes": 100,
                "secondarystorageavailable": "Unlimited",
                "secondarystoragelimit": "Unlimited",
                "secondarystoragetotal": 0,
                "sentbytes": 100,
                "snapshotavailable": "Unlimited",
                "snapshotlimit": "Unlimited",
                "snapshottotal": 0,
                "state": "enabled",
                "templateavailable": "Unlimited",
                "templatelimit": "Unlimited",
                "templatetotal": 0,
                "user": [
                    {
                        "account": "admin",
                        "accountid": "62f61a73-5f5a-11e5-bc86-0242ac11180a",
                        "accounttype": 1,
                        "apikey": "BFCyYeY0HhhrvNWC7FBwHTlso52ow3XWZUzNuiv03jmJOQn-5UH5BTVRVheiBnaRfYFE9yOzjwiqkbahq9P-Rw",
                        "created": "2015-09-20T05:42:31+0000",
                        "domain": "ROOT",
                        "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                        "email": "admin@mailprovider.com",
                        "firstname": "Admin",
                        "id": "62f6394b-5f5a-11e5-bc86-0242ac11180a",
                        "iscallerchilddomain": false,
                        "isdefault": true,
                        "lastname": "User",
                        "secretkey": "Em1UPdKTYEDaUxPXKtpqy_1uYluHWU2ZHotPi-VnBi_-MOEUAANoFHg3ycDProluDzYAXmT3YRssk367ylr_4A",
                        "state": "enabled",
                        "username": "admin"
                    }
                ],
                "vmavailable": "Unlimited",
                "vmlimit": "Unlimited",
                "vmrunning": 1,
                "vmtotal": 1,
                "volumeavailable": "Unlimited",
                "volumelimit": "Unlimited",
                "volumetotal": 1,
                "vpcavailable": "Unlimited",
                "vpclimit": "Unlimited",
                "vpctotal": 0
            }
        ],
        "count": 1
    }
}
		`
		fmt.Fprintln(w, resp)
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewListAccountsParameter()
	accounts, _ := client.ListAccounts(p)

	if len(accounts) != 1 {
		t.Errorf("length: actual %d, expected 1", len(accounts))
	}

	if accounts[0].Name.String() != "admin" {
		t.Errorf("name: actual %s, expected admin", accounts[0].Name.String())
	}

	if len(accounts[0].User) != 1 {
		t.Errorf("length of user: actual %s, expected 1", len(accounts[0].User))
	}

	if accounts[0].User[0].UserName.String() != "admin" {
		t.Errorf("username: actual %s, expected admin", accounts[0].User[0].UserName.String())
	}
}
