package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestListUsers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := `
{
    "listusersresponse": {
        "count": 1,
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
}
		`
		fmt.Fprintln(w, resp)
	}))
	defer server.Close()

	endpoint, _ := url.Parse(server.URL)
	client, _ := NewClient(endpoint, "APIKEY", "SECRETKEY", "", "")
	p := NewListUsersParameter()
	users, _ := client.ListUsers(p)

	if len(users) != 1 {
		t.Errorf("length: actual %d, expected 1", len(users))
	}

	if users[0].UserName.String() != "admin" {
		t.Errorf("username: actual %s, expected admin", users[0].UserName.String())
	}

	if users[0].Domain.String() != "ROOT" {
		t.Errorf("domain: actual %s, expected ROOT", users[0].Domain.String())
	}
}
