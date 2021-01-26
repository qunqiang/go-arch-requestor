package requestor

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	. "net/http"
	url2 "net/url"
	"reflect"
	"time"
)

type Requestor struct {
	// URL scheme
	RequestURI string
	// or
	Schema      string
	Hostname    string
	Port        string
	Path        string
	QueryString string
	Anchor      string

	// Request Control
	Timeout     time.Duration

	// common data
	//Jar *RequestJar

	// Request Data section
	Method string
	RequestBody []byte
	Headers Header

	// Response Data section
	response 		*Response
	ResponseHeader Header
	ResponseBody   *bytes.Buffer
	StatusCode     int
}

// Just enough correctness for our redirect tests. Uses the URL.Host as the
// scope of all cookies.
//type RequestJar struct {
//	m      sync.Mutex
//	perURL map[string][]*Cookie
//}
//
//func (j *RequestJar) SetCookies(u *url2.URL, cookies []*Cookie) {
//	j.m.Lock()
//	defer j.m.Unlock()
//	if j.perURL == nil {
//		j.perURL = make(map[string][]*Cookie)
//	}
//	j.perURL[u.Host] = cookies
//}
//
//func (j *RequestJar) Cookies(u *url2.URL) []*Cookie {
//	j.m.Lock()
//	defer j.m.Unlock()
//	return j.perURL[u.Host]
//}


func (req *Requestor) HasTimeout() bool {
	return req.Timeout > 0
}

func (req *Requestor) IsSuccess() bool {
	ctx := context.Background()
	var cancel context.CancelFunc
	if req.HasTimeout() {
		ctx, cancel = context.WithTimeout(context.Background(), req.Timeout)
		defer cancel()
	}
	var url *url2.URL
	var err error
	if len(req.RequestURI) > 0 {
		url,err = url2.Parse(req.RequestURI)
		if err != nil {
			panic(err)
		}
	}

	request,err := NewRequestWithContext(ctx, req.Method, url.String(), bytes.NewBuffer(req.RequestBody))
	if err != nil {
		panic(err)
	}

	client := &Client{
		//Jar:           req.Jar,
		Timeout:       req.Timeout,
	}

	resp,err := client.Do(request)
	req.response = resp
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		// parse response
		req.StatusCode = resp.StatusCode
		req.ResponseHeader = resp.Header
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		req.ResponseBody = bytes.NewBuffer(body)
	}

	return req.StatusCode == StatusOK
}

func (req *Requestor) GetBody() []byte {
	return req.ResponseBody.Bytes()
}

func (req *Requestor) GetStatusCode() int {
	return req.StatusCode
}


func (req *Requestor) GetResponseHeader() Header {
	return req.ResponseHeader
}

func (req *Requestor) UnmarshalBody(obj interface{}) error {
	return json.Unmarshal(req.ResponseBody.Bytes(), obj)
}

func (req *Requestor) GetResponse() *Response {
	return req.response
}

func (req *Requestor) DumpResponse() {
	v := reflect.ValueOf(*req.GetResponse())
	t := v.Type()
	for i:=0; i < t.NumField(); i ++ {
		log.Println(t.Field(i).Name, "=", v.Field(i))
	}
}