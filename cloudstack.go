package cloudstack

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

func takeContentFromAPIResponse(
	command string, response []byte) ([]byte, error) {
	var v map[string]json.RawMessage
	if err := json.Unmarshal(response, &v); err != nil {
		log.Println("json.Unmarshal failed:", err)
		return nil, err
	}
	content, ok := v[strings.ToLower(command)+"response"]
	if !ok {
		return nil, errors.New("Unexpected Response format")
	}
	return content, nil
}

type Client struct {
	EndPoint   url.URL
	APIKey     string
	SecretKey  string
	Username   string
	Password   string
	SessionKey string
	HTTPClient *http.Client
}

func NewClient(endpoint url.URL, apikey string, secretkey string,
	username string, password string) (*Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		log.Println("cookiejar.New failed:", err)
	}
	return &Client{
		EndPoint:   endpoint,
		APIKey:     apikey,
		SecretKey:  secretkey,
		Username:   username,
		Password:   password,
		SessionKey: "",
		HTTPClient: &http.Client{Jar: jar},
	}, nil
}

func (c *Client) RequestNoWait(command string, params map[string]string) ([]byte, error) {

	if command != "login" && command != "logout" {
		if (c.APIKey == "" || c.SecretKey == "") &&
			(c.Username != "" && c.Password != "") {
			if err := c.LogIn(); err != nil {
				log.Println("LogIn failed:", err)
				return nil, err
			}
			defer c.LogOut()
		}
	}

	queryURL := c.GenerateQueryURL(command, params)
	log.Println("queryURL:", queryURL)
	log.Println("request cookie:", c.HTTPClient.Jar)

	req, _ := http.NewRequest("GET", queryURL, nil)
	resp, err := c.HTTPClient.Do(req)
	log.Println("http response:", resp)

	if err != nil {
		log.Println("HTTPClient.Do failed:", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		response, _ := ioutil.ReadAll(resp.Body)
		err = errors.New(
			fmt.Sprintf("StatusCode %d: %s", resp.StatusCode, response))
		return nil, err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	log.Println("response:", string(response))
	log.Println("response cookie:", c.HTTPClient.Jar)
	if err != nil {
		log.Println("ioutil.ReadAll failed:", err)
		return response, err
	}

	content, err := takeContentFromAPIResponse(command, response)
	if err != nil {
		log.Println("takeContentFromAPIResponse failed:", err)
		return nil, err
	}
	log.Println("content:", string(content))

	return content, nil
}

func (c *Client) GenerateQueryURL(command string, params map[string]string) string {

	var queryURL string
	values := url.Values{}
	values.Add("command", command)
	values.Add("response", "json")
	for k, v := range params {
		values.Add(k, v)
	}

	if c.APIKey != "" && c.SecretKey != "" {
		values.Add("apikey", c.APIKey)
		queryStr := values.Encode()

		mac := hmac.New(sha1.New, []byte(c.SecretKey))
		mac.Write([]byte(
			strings.ToLower(strings.Replace(queryStr, "+", "%20", -1))))
		signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		signature = url.QueryEscape(signature)
		path := c.EndPoint.Path
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		queryURL = fmt.Sprintf(
			"%s://%s%s?%s&signature=%s",
			c.EndPoint.Scheme, c.EndPoint.Host, path,
			queryStr, signature)
	} else if command == "login" && c.Username != "" && c.Password != "" {
		values.Add("username", c.Username)
		values.Add("password", c.Password)
		queryStr := values.Encode()
		path := c.EndPoint.Path
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		queryURL = fmt.Sprintf(
			"%s://%s%s?%s",
			c.EndPoint.Scheme, c.EndPoint.Host, path, queryStr)
	} else if c.SessionKey != "" {
		values.Add("sessionkey", c.SessionKey)
		queryStr := values.Encode()
		path := c.EndPoint.Path
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		queryURL = fmt.Sprintf(
			"%s://%s%s?%s",
			c.EndPoint.Scheme, c.EndPoint.Host, path, queryStr)
	} else {
		log.Println("Failed to authenticate: You must provide apikey/secretkey or username/password")
	}

	return queryURL
}

type LogInResponse struct {
	Username       NullString `json:"username"`
	Userid         ID         `json:"userid"`
	Password       NullString `json:"password"`
	Domainid       ID         `json:"domainid"`
	Timeout        NullString `json:"timeout"`
	Account        NullString `json:"account"`
	Firstname      NullString `json:"firstname"`
	Lastname       NullString `json:"lastname"`
	Type           NullString `json:"type"`
	Timezone       NullString `json:"timezone"`
	Timezoneoffset NullString `json:"timezoneoffset"`
	Sessionkey     NullString `json:"sessionkey"`
}

type LogOutResponse struct {
	Description NullString `json:"description"`
}

func (c *Client) LogIn() error {

	params := map[string]string{}
	b, err := c.RequestNoWait("login", params)
	if err != nil {
		log.Println(err)
		return err
	}

	lr := LogInResponse{}
	err = json.Unmarshal(b, &lr)
	if err != nil {
		log.Println("json.Unmarshal failed:", err)
		return err
	}

	if !lr.Sessionkey.Valid {
		return errors.New("login API didn't return valid sessionkey.")
	}

	c.SessionKey = lr.Sessionkey.String
	return nil
}

func (c *Client) LogOut() error {

	params := map[string]string{}
	b, err := c.RequestNoWait("logout", params)

	lr := LogOutResponse{}
	err = json.Unmarshal(b, &lr)
	if err != nil {
		log.Println("json.Unmarshal failed:", err)
		return err
	}

	if !lr.Description.Valid || lr.Description.String != "success" {
		return errors.New(
			fmt.Sprintf("logout API is failed. description: %v", lr.Description))
	}
	return nil
}

func (c *Client) Request(command string, params map[string]string) ([]byte, error) {

	isAsync, ok := isAsyncMap[command]
	if !ok {
		err := errors.New(fmt.Sprintf("%v : No such command", command))
		log.Println(err)
		return nil, err
	}

	b, err := c.RequestNoWait(command, params)
	if err != nil {
		log.Println("RequestNoWait failed:", err)
		log.Println("RequestNoWait returned:", string(b))
		return b, err
	}

	if isAsync {
		qr := new(QueryAsyncJobResultResponse)

		err = json.Unmarshal(b, qr)
		if err != nil {
			log.Println("Unmarshal failed", err)
			return b, err
		}

		qr, err = c.Wait(qr.Jobid.String)
		if err != nil {
			log.Println("Wait failed", err)
			return qr.Jobresult, err
		}

		return qr.Jobresult, nil
	} else {
		return b, nil
	}
}

func (c *Client) Wait(jobid string) (*QueryAsyncJobResultResponse, error) {
	params := map[string]string{
		"jobid": jobid,
	}
	qr := new(QueryAsyncJobResultResponse)

	for {
		b, err := c.RequestNoWait("queryAsyncJobResult", params)
		if err != nil {
			log.Println("Request failed", err)
			return nil, err
		}
		err = json.Unmarshal(b, qr)
		if err != nil {
			log.Println("Unmarshal failed", err)
			return nil, err
		}
		if !(qr.Jobstatus.Valid && qr.Jobstatus.Int64 == 0) {
			break
		}
		time.Sleep(waitInterval * time.Second)
	}

	if qr.Jobstatus.Valid && qr.Jobstatus.Int64 == 1 {
		return qr, nil
	} else {
		err := errors.New(string(qr.Jobresult))
		log.Println(err)
		log.Println(qr)
		return qr, err
	}
}
