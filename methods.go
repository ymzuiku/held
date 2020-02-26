package held

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

// MethodFn method函数类型
type MethodFn func(map[string]interface{}) (map[string]interface{}, error)

// Methods 所有方法集合
var Methods map[string]MethodFn

// MethodRegisterHTTPAndSocket 注册Http及Socket
func MethodRegisterHTTPAndSocket(app *gin.Engine) {
	app.POST("/methods", func(ctx *gin.Context) {
		body, err := GinBody(ctx)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "json format error"})
			return
		}

		if reflect.TypeOf(body["url"]).String() != "string" {
			ctx.JSON(500, gin.H{"error": "url is not string"})
			return
		}
		url := body["url"].(string)
		method := Methods[url]
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

// AddMethod 注册一个method
func AddMethod(method string, fn MethodFn) {
	Methods[method] = fn
}
