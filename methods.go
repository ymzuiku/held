package held

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

// Method 一个兼容HTTP和WebSocket的注册器
type Method struct {
	Events map[string]interface{}
}

// InitHTTPAndSocket 注册Http及Socket
//
// 在gin项目初始化之后，注册Method，这样会以 baseURL 注册一个路由，并且分发方法至 HTTP 和 WebSocket
//
// 使用方法：
//		held.MethodRegisterHTTPAndSocket(app, "/methods")
func (method *Method) InitHTTPAndSocket(app *gin.Engine, baseURL string) {
	method.Events = map[string]interface{}{}
	app.POST(baseURL, func(ctx *gin.Context) {
		body, err := GinBody(ctx)
		println(body)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "json format error"})
			return
		}

		if reflect.TypeOf(body["url"]).String() != "string" {
			ctx.JSON(500, gin.H{"error": "url is not string"})
			return
		}
		url := body["url"].(string)
		method := method.Events[url].(func(map[string]interface{}) (interface{}, error))
		data := body["data"].(map[string]interface{})

		if method == nil {
			ctx.JSON(500, gin.H{"error": "url is empty"})
			return
		}

		res, err := method(data)

		if err != nil {
			ctx.JSON(200, gin.H{"url": url, "error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"url": url, "data": res})
	})
}

// Add 注册一个method
func (method *Method) Add(url string, fn func(map[string]interface{}) (interface{}, error)) {
	method.Events[url] = fn
}
