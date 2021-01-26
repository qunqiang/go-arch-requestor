# requestor

#### 介绍
config as request

#### 软件架构
软件架构说明


#### 安装教程
```shell
go get github.com/qunqiangrequestor
```
#### 使用说明


```go
package demo

import (
	req "github.com/qunqiangrequestor"
	"log"
	"net/http"
)

func demo() {
	requestor := &req.Requestor{
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
		requestor.DumpResponse()
		log.Println("请求失败", requestor.GetStatusCode())
	}
}
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request
