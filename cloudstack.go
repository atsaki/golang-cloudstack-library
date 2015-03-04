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
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func convertJsonToMap(b []byte) (map[string]json.RawMessage, error) {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("Failed to convert json to map: %v", err)
	}
	return m, nil
}

// Get error message from API response
func getErrorText(m map[string]json.RawMessage) string {

	quoted, ok := m["errortext"]
	if !ok {
		return ""
	}

	errortext, err := strconv.Unquote(string(quoted))
	if err != nil {
		return string(quoted)
	}
	return errortext
}

func getResponseContent(cmd *Command, resp []byte) (respBody []byte, err error) {

	var ok bool

	respMap, err := convertJsonToMap(resp)
	if err != nil {
		return nil, err
	}

	respBody, ok = respMap[strings.ToLower(cmd.Name)+"response"]
	if !ok {
		// some API's response are not ended with "response"
		respBody, ok = respMap[strings.ToLower(cmd.Name)]
	}
	if !ok {
		if errortext := getErrorText(respMap); errortext != "" {
			return nil, errors.New(errortext)
		}
		return nil, fmt.Errorf(
			"Unexpected format: response doesn't contain %s or %s or errortext",
			cmd.Name+"response", cmd.Name)
	}
	return respBody, nil
}

// Get content from API response
func getObjectJson(cmd *Command, resp []byte, isJobResult bool) (objJson []byte, err error) {

	var respBody []byte

	if isJobResult {
		respBody = resp
	} else {
		respBody, err = getResponseContent(cmd, resp)
		if err != nil {
			return nil, err
		}
		if cmd.IsAsync {
			// return jobid
			return respBody, nil
		}
	}

	switch cmd.ObjectType {
	case "result":
		return respBody, nil
	default:
		respBodyMap, err := convertJsonToMap(respBody)
		if err != nil {
			return nil, err
		}

		objJson, ok := respBodyMap[cmd.ObjectType]
		if !ok {
			if errortext := getErrorText(respBodyMap); errortext != "" {
				return nil, errors.New(errortext)
			}
			// respBodyMap can be empty. For example, list api returns nothing.
			if len(respBodyMap) == 0 {
				return nil, nil
			}
			// Type can be null. For example, destroyvirtualmachine is executed with expunge=true
			if objJson, ok := respBodyMap["null"]; ok {
				return objJson, nil
			}
			return nil, fmt.Errorf(
				"Unexpected format: response doesn't contain %s or errortext",
				cmd.ObjectType)
		}
		return objJson, nil
	}
}

func (client *Client) convertResponseJsonToObject(cmd *Command, resp []byte, isJobResult bool) (interface{}, error) {

	objJson, err := getObjectJson(cmd, resp, isJobResult)
	if err != nil {
		return nil, err
	}

	p := cmd.Pointer()
	if objJson == nil {
		return reflect.ValueOf(p).Elem().Interface(), nil
	}
	if err := json.Unmarshal(objJson, p); err != nil {
		return nil, fmt.Errorf("Failed to convert json to object: %v", err)
	}
	obj := reflect.ValueOf(p).Elem().Interface()

	// Set Object's Client field
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			v.Index(i).Interface().(Resource).setClient(client)
		}
	} else if v.Kind() == reflect.Ptr {
		obj.(Resource).setClient(client)
	}

	return obj, nil
}

// convert APIParameter to map[string]interface{}
func convertParamToMap(p APIParameter) (m map[string]interface{}) {
	m = make(map[string]interface{})
	v := reflect.ValueOf(p).Elem()
	for i := 0; i < v.NumField(); i++ {
		if unicode.IsUpper(rune(v.Type().Field(i).Name[0])) {
			m[strings.ToLower(v.Type().Field(i).Name)] = v.Field(i).Interface()
		}
	}
	return m
}

type LogInResponse struct {
	Username       NullString `json:"username"`
	UserId         ID         `json:"userid"`
	Password       NullString `json:"password"`
	DomainId       ID         `json:"domainid"`
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

type AsyncApiResponse struct {
	Id    ID `json:"id"`
	JobId ID `json:"jobid"`
}

type AsyncJobResult struct {
	AccountId       NullString      `json:"accountid"`
	Cmd             NullString      `json:"cmd"`
	Created         NullString      `json:"created"`
	JobId           ID              `json:"jobid"`
	JobInstanceId   ID              `json:"jobinstanceid"`
	JobInstanceType NullString      `json:"jobinstancetype"`
	JobProcsSatus   NullNumber      `json:"jobprocstatus"`
	JobResult       json.RawMessage `json:"jobresult"`
	JobResultCode   NullNumber      `json:"jobresultcode"`
	JobResultType   NullString      `json:"jobresulttype"`
	JobStatus       NullNumber      `json:"jobstatus"`
	UserId          ID              `json:"userid"`
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

// Generate Query URL from command and paramters
func (c *Client) GenerateQueryURL(command string, params map[string]interface{}) string {

	queryURL := c.EndPoint

	values := url.Values{}
	values.Add("command", command)
	values.Add("response", "json")
	for k := range params {

		if k == "userdata" {
			_, ok := params[k]
			if ok {
				s := fmt.Sprint(params["userdata"])
				// There seems to be a issue if base64 encoded string ended with
				// padding character "="
				// add space to original string to make result not ended with "="
				for {
					if len(s)%3 == 0 {
						break
					}
					s += " "
				}
				values.Add(k, base64.StdEncoding.EncodeToString([]byte(s)))
			}
			continue
		}

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

	if c.APIKey != "" && c.SecretKey != "" {
		values.Add("apikey", c.APIKey)

		// queryStr must not be URL encoded.  You can't use values.Encode.
		keys := make([]string, 0, len(values))
		for k := range values {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		params := make([]string, 0, len(keys))
		for _, k := range keys {
			params = append(params, fmt.Sprintf("%s=%s", k, values[k][0]))
		}
		queryStr := strings.Join(params, "&")

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

func (c *Client) AsyncRequestJson(command string, params map[string]interface{}) (resp []byte, err error) {
	cmd := getCommand(command)

	queryURL := c.GenerateQueryURL(cmd.Name, params)
	log.Println("QueryURL:", queryURL)
	log.Println("request cookie:", c.HTTPClient.Jar)

	httpReq, _ := http.NewRequest("GET", queryURL, nil)
	httpResp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		log.Println("HTTPClient.Do failed:", err)
		return nil, err
	}
	defer httpResp.Body.Close()

	resp, err = ioutil.ReadAll(httpResp.Body)

	log.Println("statuscode:", httpResp.StatusCode)
	log.Println("response:", string(resp))
	log.Println("cookie:", c.HTTPClient.Jar)

	if httpResp.StatusCode != 200 {
		return resp, fmt.Errorf("StatusCode %d: %s", httpResp.StatusCode, resp)
	}

	return resp, err
}

func (c *Client) AsyncRequest(command string, params map[string]interface{}) (interface{}, error) {
	cmd := getCommand(command)
	resp, err := c.AsyncRequestJson(cmd.Name, params)
	if err != nil {
		return nil, err
	}
	if cmd.IsAsync {
		var r *AsyncApiResponse
		respJson, err := getObjectJson(cmd, resp, false)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(respJson, &r); err != nil {
			return nil, err
		}
		if r.JobId.IsNil() {
			m, _ := convertJsonToMap(respJson)
			errortext := getErrorText(m)
			if errortext != "" {
				return nil, errors.New(errortext)
			}
			return nil, fmt.Errorf("No jobid returned. %s", string(respJson))
		}
		return r, nil
	}
	return c.convertResponseJsonToObject(cmd, resp, false)
}

func (c *Client) RequestJson(command string, params map[string]interface{}) ([]byte, error) {
	cmd := getCommand(command)
	if cmd.IsAsync {
		r, err := c.AsyncRequest(cmd.Name, params)
		if err != nil {
			return nil, err
		}
		job, err := c.Wait(r.(*AsyncApiResponse).JobId.String())
		if err != nil {
			return nil, err
		}
		return job.JobResult, nil
	}
	return c.AsyncRequestJson(command, params)
}

func (c *Client) Request(command string, params map[string]interface{}) (interface{}, error) {
	cmd := getCommand(command)
	resp, err := c.RequestJson(cmd.Name, params)
	if err != nil {
		return nil, err
	}
	// If command is Async, pass JobResult.
	return c.convertResponseJsonToObject(cmd, resp, cmd.IsAsync)
}

func (c *Client) LogIn() (err error) {

	r, err := c.AsyncRequestJson("login", nil)
	if err != nil {
		return err
	}
	lr := LogInResponse{}
	if err = json.Unmarshal(r, &lr); err != nil {
		return err
	}

	if lr.SessionKey.IsNil() {
		return errors.New("login API didn't return valid sessionkey.")
	}

	c.SessionKey = lr.SessionKey.String()
	return nil
}

func (c *Client) LogOut() error {

	r, err := c.AsyncRequestJson("logout", nil)
	lr := LogOutResponse{}
	if err = json.Unmarshal(r, &lr); err != nil {
		return err
	}

	if lr.Description.String() != "success" {
		return fmt.Errorf("logout API is failed. description: %v", lr.Description)
	}
	return nil
}

func (c *Client) QueryAsyncJobResult(jobid string) (job *AsyncJobResult, err error) {
	resp, err := c.AsyncRequestJson(
		"queryAsyncJobResult", map[string]interface{}{"jobid": jobid})
	if err != nil {
		return nil, err
	}

	m, err := convertJsonToMap(resp)
	if err != nil {
		return nil, err
	}

	jobJson, ok := m["queryasyncjobresultresponse"]
	if !ok {
		if errortext := getErrorText(m); errortext != "" {
			return nil, errors.New(errortext)
		}
		return nil, fmt.Errorf(
			"Unexpected format: response doesn't contain queryasyncjobresultresponse or errortext")
	}

	if err = json.Unmarshal(jobJson, &job); err != nil {
		return nil, fmt.Errorf("Failed to convert json to AsyncJobResult: %v", err)
	}

	return job, nil
}

func (c *Client) Wait(jobid string) (job *AsyncJobResult, err error) {

	for {
		job, err = c.QueryAsyncJobResult(jobid)
		if err != nil {
			return nil, err
		}

		if job.JobStatus.String() != "0" {
			break
		}
		time.Sleep(c.PollingInterval * time.Second)
	}

	if job.JobStatus.String() != "1" {
		m, err := convertJsonToMap(job.JobResult)
		if err != nil {
			return nil, err
		}
		errortext := getErrorText(m)
		if errortext != "" {
			return nil, errors.New(errortext)
		}
		return nil, fmt.Errorf(
			"job doesn't finished properly: %s", job.JobStatus.String())
	}

	return job, nil
}
