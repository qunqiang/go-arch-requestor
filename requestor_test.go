package requestor

import (
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestRequestor_IsSuccess(t *testing.T) {
	requestor := &Requestor{
		RequestURI: "http://httpbin.org/anything",
		Headers: http.Header {
			"content-type" : {"application/json"},
			"x-trace-id": {"123jadfn3829afl3"},
		},
		Method: "POST",
		RequestBody: []byte("{\"abc\":\"hello world\"}"),
	}

	type ResponseStruct struct {
		Origin string `json:"origin"`
		Headers map[string]string `json:"headers"`
		Method string `json:"method"`
		Body	string `json:"body"`
		Data    string `json:"data"`
		Json 	interface{} `json:"json"`
		File    interface{} `json:"file"`
		Form    interface{} `json:"form"`
	}
	resp := ResponseStruct{}
	if requestor.IsSuccess() {
		requestor.DumpResponse()
		log.Println(requestor.GetStatusCode())
		log.Println(requestor.GetResponseHeader())
		log.Println(requestor.GetBody())
		err := requestor.UnmarshalBody(&resp)
		if err != nil {
			panic(err)
		}
		v := reflect.ValueOf(resp)
		t := v.Type()

		for i:=0; i < t.NumField(); i ++ {
			log.Println(t.Field(i).Name, "=", v.Field(i))
		}
	} else {
		log.Println("请求失败", requestor.GetStatusCode())
	}
}