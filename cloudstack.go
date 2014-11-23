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
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// convert APIParameter to map[string]interface{}
func convertParamToMap(p APIParameter) (m map[string]interface{}) {
	m = make(map[string]interface{})
	v := reflect.ValueOf(p).Elem()
	for i := 0; i < v.NumField(); i++ {
		m[strings.ToLower(v.Type().Field(i).Name)] = v.Field(i).Interface()
	}
	return m
}

// Get content from API response
func getContent(command string, response []byte) (content []byte, err error) {
	var v map[string]json.RawMessage
	var ok bool

	if err := json.Unmarshal(response, &v); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal: %s", string(response))
	}

	command = strings.ToLower(command)
	content, ok = v[command+"response"]
	if !ok {
		content, ok = v[command]
	}
	if !ok {
		return nil, fmt.Errorf("Failed to get content: %s", string(response))
	}
	return content, nil
}

// Get error message from API response
func getErrorText(b []byte) (errortext string) {
	var v map[string]json.RawMessage
	var err error

	if err = json.Unmarshal(b, &v); err != nil {
		fmt.Errorf("Failed to unmarshal: %s", string(b))
	}

	quotedBytes, ok := v["errortext"]
	if !ok {
		return ""
	}

	quotedStr := string(quotedBytes)
	errortext, err = strconv.Unquote(quotedStr)
	if err != nil {
		return quotedStr
	}
	return errortext
}

type Client struct {
	EndPoint        *url.URL
	APIKey          string
	SecretKey       string
	Username        string
	Password        string
	SessionKey      string
	PollingInterval time.Duration
	HTTPClient      *http.Client
}

func NewClient(endpoint *url.URL, apikey string, secretkey string,
	username string, password string) (*Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		log.Println("cookiejar.New failed:", err)
	}
	return &Client{
		EndPoint:        endpoint,
		APIKey:          apikey,
		SecretKey:       secretkey,
		Username:        username,
		Password:        password,
		SessionKey:      "",
		PollingInterval: config.PollingInterval,
		HTTPClient:      &http.Client{Jar: jar},
	}, nil
}

func (c *Client) Request(command string, params map[string]interface{}) (content []byte, err error) {

	if command != "login" && command != "logout" {
		if (c.APIKey == "" || c.SecretKey == "") &&
			(c.Username != "" && c.Password != "") {
			if err = c.LogIn(); err != nil {
				log.Println("LogIn failed:", err)
				return content, err
			}
			defer c.LogOut()
		}
	}

	queryURL := c.GenerateQueryURL(command, params)
	log.Println("QueryURL:", queryURL)
	log.Println("request cookie:", c.HTTPClient.Jar)

	httpReq, _ := http.NewRequest("GET", queryURL, nil)
	httpResp, err := c.HTTPClient.Do(httpReq)

	if err != nil {
		log.Println("HTTPClient.Do failed:", err)
		return content, err
	}
	defer httpResp.Body.Close()

	log.Println("statuscode:", httpResp.StatusCode)

	if httpResp.StatusCode != 200 {
		respBody, _ := ioutil.ReadAll(httpResp.Body)
		err = errors.New(
			fmt.Sprintf("StatusCode %d: %s", httpResp.StatusCode, respBody))
		return content, err
	}

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		log.Println("ioutil.ReadAll failed:", err)
		return content, err
	}

	log.Println("response:", string(respBody))
	log.Println("cookie:", c.HTTPClient.Jar)

	return getContent(command, respBody)
}

// Generate Query URL from command and paramters
func (c *Client) GenerateQueryURL(command string, params map[string]interface{}) string {

	queryURL := c.EndPoint

	values := url.Values{}
	values.Add("command", command)
	values.Add("response", "json")
	for k := range params {
		switch v := params[k].(type) {
		case []string:
			if len(v) > 0 {
				values.Add(k, strings.Join(v, ","))
			}
		case map[string]string:
			if len(v) > 0 {
				i := 0
				for key, value := range v {
					values.Add(fmt.Sprintf("%s[%d].key", k, i), key)
					values.Add(fmt.Sprintf("%s[%d].value", k, i), value)
					i += 1
				}
			}
		case NullBool, NullString, NullNumber, ID:
			if !v.(Nullable).IsNil() {
				values.Add(k, fmt.Sprint(v))
			}
		default:
			values.Add(k, fmt.Sprint(v))
		}
	}

	// values.Encode sort values by key
	if c.APIKey != "" && c.SecretKey != "" {
		values.Add("apikey", c.APIKey)
		queryStr := values.Encode()
		mac := hmac.New(sha1.New, []byte(c.SecretKey))
		mac.Write([]byte(
			strings.ToLower(strings.Replace(queryStr, "+", "%20", -1))))
		signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		signature = url.QueryEscape(signature)
		queryURL.RawQuery = queryStr + "&signature=" + signature
	} else {
		if command == "login" && c.Username != "" && c.Password != "" {
			values.Add("username", c.Username)
			values.Add("password", c.Password)
		} else if c.SessionKey != "" {
			values.Add("sessionkey", c.SessionKey)
		}
		queryURL.RawQuery = values.Encode()
	}
	return queryURL.String()
}

type LogInResponse struct {
	Username       NullString `json:"username"`
	UserId         NullString `json:"userid"`
	Password       NullString `json:"password"`
	DomainId       NullString `json:"domainid"`
	Timeout        NullString `json:"timeout"`
	Account        NullString `json:"account"`
	FirstName      NullString `json:"firstname"`
	LastName       NullString `json:"lastname"`
	Type           NullString `json:"type"`
	TimeZone       NullString `json:"timezone"`
	TimeZoneOffset NullString `json:"timezoneoffset"`
	SessionKey     NullString `json:"sessionkey"`
}

type LogOutResponse struct {
	Description NullString `json:"description"`
}

func (c *Client) LogIn() error {

	var params map[string]interface{}
	b, err := c.Request("login", params)
	if err != nil {
		log.Println(err)
		return err
	}

	r := LogInResponse{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal: %s", string(b))
	}

	if r.SessionKey.IsNil() {
		return errors.New("login API didn't return valid sessionkey.")
	}

	c.SessionKey = r.SessionKey.String()
	return nil
}

func (c *Client) LogOut() error {

	var params map[string]interface{}
	b, err := c.Request("logout", params)

	r := LogOutResponse{}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal: %s", string(b))
	}

	if r.Description.IsNil() || r.Description.String() != "success" {
		return fmt.Errorf("logout API is failed. description: %v", r.Description)
	}
	return nil
}

// Get API command from Parameter (ex. ListZonesParameter -> listZones)
func getCommandNameFromParam(param APIParameter) string {
	paramName := reflect.TypeOf(param).Elem().Name()
	paramName = strings.ToLower(paramName[0:1]) + paramName[1:]
	re := regexp.MustCompile("(.*\\.)?([^.]+)Parameter$")
	return re.ReplaceAllString(paramName, "$2")
}

// request executes SyncRequest and unmarshals response
func (c *Client) request(param APIParameter, obj interface{}) error {
	var v map[string]json.RawMessage
	var content json.RawMessage
	var ok bool

	cmdName := getCommandNameFromParam(param)
	cmd := config.Commands[cmdName]

	resp, err := c.SyncRequest(cmdName, convertParamToMap(param))
	if err != nil {
		return err
	}

	if cmd.Object == "result" {
		content = resp
	} else {
		err = json.Unmarshal(resp, &v)
		if err != nil {
			return fmt.Errorf(
				"Failed to unmarshal SyncRequest result (%s): %s", err, string(resp))
		}

		content, ok = v[cmd.Object]
		if !ok {
			if len(v) == 0 {
				return nil
			}
			errortext, _ := v["errortext"]
			if ok {
				return fmt.Errorf(string(errortext))
			} else {
				return fmt.Errorf(
					"Unexpected SyncRequest response format: response doesn't contain '%s' or 'errortext' %s",
					cmd.Object, string(resp))
			}
		}
	}

	if err = json.Unmarshal(content, obj); err != nil {
		return fmt.Errorf(
			"Failed to unmarshal content (%s): %s", err, string(content))
	}

	return nil
}

// SyncRequest executes command and wait for job complettion
func (c *Client) SyncRequest(command string, params map[string]interface{}) (b []byte, err error) {

	cmd, ok := config.Commands[command]
	if !ok {
		return nil, fmt.Errorf("%v : No such command", command)
	}

	b, err = c.Request(command, params)
	if !cmd.IsAsync || err != nil {
		return b, err
	}

	job := new(AsyncJobResult)
	if err = json.Unmarshal(b, job); err != nil {
		return b, fmt.Errorf("Failed to unmarshal: %s", string(b))
	}

	if job.JobId.IsNil() {
		return b, errors.New(getErrorText(b))
	}

	job, err = c.Wait(job.JobId.String())
	if job == nil || err != nil {
		return nil, err
	}
	return job.JobResult, nil
}

// Wait waits for job completion
func (c *Client) Wait(jobid string) (job *AsyncJobResult, err error) {

	p := new(QueryAsyncJobResultParameter)
	p.JobId.Set(jobid)

	for {
		job, err = c.QueryAsyncJobResult(p)
		if err != nil {
			return job, err
		}

		if job.JobStatus.IsNil() || job.JobStatus.String() != "0" {
			break
		}
		time.Sleep(c.PollingInterval * time.Second)
	}

	return job, err
}
