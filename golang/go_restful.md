# Go Restful
## 简介

Go Restful是一个用于构建RESTful API的Go语言框架。它提供了一组简单而强大的工具，帮助开发者轻松地创建和管理RESTful服务。

使用`go get`安装go restful框架
```shell
go get -u github.com/emicklei/go-restful
```
```go
package main

import (
   "github.com/emicklei/go-restful"
)

func main() {
   ws := new(restful.WebService)
   ws.Path("/hello").
      Consumes(restful.MIME_JSON).
      Produces(restful.MIME_JSON).
      Route(ws.GET("/").To(hello))

   restful.Add(ws)
   http.ListenAndServe(":8080", nil)
}

func hello(req *restful.Request, resp *restful.Response) {
   resp.WriteEntity(map[string]string{"result": "Hello, World!"})
}

```
### WebService中一些字段的含义
- Path(string): 指定WebService的基本路径。例如，ws.Path("/api/v1")将WebService的基本路径设置为/api/v1。
- Consumes(...string): 定义WebService可以接受的请求内容类型。例如，ws.Consumes(restful.MIME_JSON)表示该WebService可以处理JSON格式的请求。
- Produces(...string): 定义WebService可以产生的响应内容类型。例如，ws.Produces(restful.MIME_JSON)表示该WebService可以生成JSON格式的响应。
- Route(RouteBuilder): 定义API的路由规则。通过ws.Route方法，可以指定不同HTTP方法（GET、POST、PUT、DELETE等）的处理函数。
- Param(...Param): 定义路由参数，用于从URL中提取参数。例如，ws.Param(ws.PathParameter("id", "Identifier for the resource").DataType("string"))定义了一个名为"id"的路径参数。
- Produces(...string): 定义WebService可以产生的响应内容类型。例如，ws.Produces(restful.MIME_JSON)表示该WebService可以生成JSON格式的响应。
- ResponseHeader(string, string): 定义响应头。例如，ws.ResponseHeader("Cache-Control", "no-store")设置响应头"Cache-Control"为"no-store"。
- Consumes(...string): 定义WebService可以接受的请求内容类型。例如，ws.Consumes(restful.MIME_JSON)表示该WebService可以处理JSON格式的请求。