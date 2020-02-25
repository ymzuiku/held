package held

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// GinBody 获取 gin 的body
func GinBody(ctx *gin.Context) (body map[string]interface{}, err error) {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		err = errors.New("body is not json")
		return nil, err
	}
	json.Unmarshal(data, &body)

	return body, nil
}
