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
                "accountdetails": {
                    "key0": "value0",
                    "key1": "value1"
                },
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

	if accounts[0].AccountDetails["key0"] != "value0" {
		t.Errorf("accountdetails[\"key0\"]: actual %s, expected value0", accounts[0].AccountDetails["key0"])
	}
}

func TestCreateAccounts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "createaccountresponse": {
        "account": {
            "accountdetails": {
                "key0": "value0",
                "key1": "value1"
            },
            "accounttype": 2,
            "cpuavailable": "40",
            "cpulimit": "40",
            "cputotal": 0,
            "domain": "ROOT",
            "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
            "groups": [],
            "id": "c9cb9df8-dcd5-44c5-a40d-2e7d266669a6",
            "ipavailable": "20",
            "iplimit": "20",
            "iptotal": 0,
            "isdefault": false,
            "memoryavailable": "40960",
            "memorylimit": "40960",
            "memorytotal": 0,
            "name": "user01",
            "networkavailable": "20",
            "networklimit": "20",
            "networktotal": 0,
            "primarystorageavailable": "200",
            "primarystoragelimit": "200",
            "primarystoragetotal": 0,
            "projectavailable": "Unlimited",
            "projectlimit": "Unlimited",
            "projecttotal": 0,
            "secondarystorageavailable": "400",
            "secondarystoragelimit": "400",
            "secondarystoragetotal": 0,
            "snapshotavailable": "20",
            "snapshotlimit": "20",
            "snapshottotal": 0,
            "state": "enabled",
            "templateavailable": "20",
            "templatelimit": "20",
            "templatetotal": 0,
            "user": [
                {
                    "account": "user01",
                    "accountid": "c9cb9df8-dcd5-44c5-a40d-2e7d266669a6",
                    "accounttype": 2,
                    "created": "2016-07-26T02:17:17+0000",
                    "domain": "ROOT",
                    "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                    "email": "user01@example.com",
                    "firstname": "user01",
                    "id": "1d69ac45-a8f7-469c-ab24-d2b0382fa8be",
                    "iscallerchilddomain": false,
                    "isdefault": false,
                    "lastname": "user01",
                    "state": "enabled",
                    "username": "user01"
                }
            ],
            "vmavailable": "20",
            "vmlimit": "20",
            "vmtotal": 0,
            "volumeavailable": "20",
            "volumelimit": "20",
            "volumetotal": 0,
            "vpcavailable": "20",
            "vpclimit": "20",
            "vpctotal": 0
        }
    }
}
		`
		fmt.Fprintln(w, resp)
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	accounttype := 2
	email := "user01@example.com"
	firstname := "user01"
	lastname := "user01"
	password := "password"
	username := "user01"
	p := NewCreateAccountParameter(accounttype, email, firstname, lastname, password, username)
	p.AccountDetails = map[string]string{"key0": "value0", "key1": "value1"}
	account, _ := client.CreateAccount(p)

	if account.Name.String() != "user01" {
		t.Errorf("name: actual %s, expected user01", account.Name.String())
	}

	if len(account.User) != 1 {
		t.Errorf("length of user: actual %s, expected 1", len(account.User))
	}

	if account.User[0].UserName.String() != "user01" {
		t.Errorf("username: actual %s, expected admin", account.User[0].UserName.String())
	}

	if account.AccountDetails["key0"] != "value0" {
		t.Errorf("accountdetails[\"key0\"]: actual %s, expected value0", account.AccountDetails["key0"])
	}

	if account.AccountDetails["key1"] != "value1" {
		t.Errorf("accountdetails[\"key1\"]: actual %s, expected value1", account.AccountDetails["key1"])
	}
}
