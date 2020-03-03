package held

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// App 应用的上下文对象钩子
type App struct {
	Ctx    context.Context
	Gin    *gin.Engine
	Mgo    *mongo.Client
	Method *Method
}
