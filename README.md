# requestor

#### 介绍
config as request, no more decorations, just with net/http.
Why is this component called go-arch-requestor, because it's a mirror from gitee.com/go-arch/requestor.

#### 软件架构
软件架构说明


#### 安装教程
```shell
go get github.com/qunqiang/requestor
```
#### 使用说明

```go
package demo

import (
	req "github.com/qunqiang/requestor"
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
		Method: "GET",
		Body: []byte("hello world"),
	}

	type ResponseStruct struct {
		Origin string `json:"origin"`
		Headers map[string]string `json:"headers"`
		Method string `json:"method"`
		Body	string `json:"json"`
		Data    string `json:"data"`
		File    interface{} `json:"file"`
		Form    interface{} `json:"form"`
	}
	resp := &ResponseStruct{}

	if requestor.IsSuccess() {
		log.Println(requestor.GetStatusCode())
		log.Println(requestor.GetResponseHeader())
		log.Println(requestor.GetBody())
		err := requestor.UnmarshalBody(resp)
		if err != nil {
			panic(err)
		}
		for k,v := range resp {
			
        }
		log.Println(resp.Method, resp.Body, resp.Headers["User-Agent"])
	} else {
		log.Println("请求失败", requestor.GetStatusCode())
	}
}
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
