package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListProjects(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listprojectsresponse": {
        "count": 1,
        "project": [
            {
                "account": "admin",
                "cpuavailable": "40",
                "cpulimit": "40",
                "cputotal": 0,
                "displaytext": "Project01",
                "domain": "ROOT",
                "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "id": "caea82cb-75a9-4bdb-936c-62b5283956d8",
                "ipavailable": "20",
                "iplimit": "20",
                "iptotal": 0,
                "memoryavailable": "40960",
                "memorylimit": "40960",
                "memorytotal": 0,
                "name": "Project01",
                "networkavailable": "20",
                "networklimit": "20",
                "networktotal": 0,
                "primarystorageavailable": "200",
                "primarystoragelimit": "200",
                "primarystoragetotal": 0,
                "secondarystorageavailable": "400",
                "secondarystoragelimit": "400",
                "secondarystoragetotal": 0,
                "snapshotavailable": "20",
                "snapshotlimit": "20",
                "snapshottotal": 0,
                "state": "Active",
                "tags": [],
                "templateavailable": "20",
                "templatelimit": "20",
                "templatetotal": 0,
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
        ]
    }
}
		`
		fmt.Fprintln(w, resp)
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewListProjectsParameter()
	projects, _ := client.ListProjects(p)

	if len(projects) != 1 {
		t.Errorf("length: actual %d, expected 1", len(projects))
	}

	if projects[0].Name.String() != "Project01" {
		t.Errorf("name: actual %s, expected Project01", projects[0].Name.String())
	}

	if projects[0].Account.String() != "admin" {
		t.Errorf("account: actual %s, expected admin", projects[0].Account.String())
	}
}

func TestListProjectAccounts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listprojectaccountsresponse": {
        "count": 1,
        "projectaccount": [
            {
                "account": "admin",
                "accountid": "62f61a73-5f5a-11e5-bc86-0242ac11180a",
                "accounttype": 1,
                "domain": "ROOT",
                "domainid": "62f5fe2e-5f5a-11e5-bc86-0242ac11180a",
                "project": "Project01",
                "projectid": "caea82cb-75a9-4bdb-936c-62b5283956d8",
                "role": "Admin",
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
                ]
            }
        ]
    }
}


		`
		fmt.Fprintln(w, resp)
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	projectid := "caea82cb-75a9-4bdb-936c-62b5283956d8"
	p := NewListProjectAccountsParameter(projectid)
	accounts, _ := client.ListProjectAccounts(p)

	if len(accounts) != 1 {
		t.Errorf("length: actual %d, expected 1", len(accounts))
	}

	if accounts[0].Role.String() != "Admin" {
		t.Errorf("role: actual %s, expected Admin", accounts[0].Role.String())
	}

	if len(accounts[0].User) != 1 {
		t.Errorf("length of user: actual %s, expected 1", len(accounts[0].User))
	}

	if accounts[0].User[0].UserName.String() != "admin" {
		t.Errorf("username: actual %s, expected admin", accounts[0].User[0].UserName.String())
	}
}
