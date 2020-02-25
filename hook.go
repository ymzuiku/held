package held

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Hook 应用的上下文对象钩子
type Hook struct {
	Ctx context.Context
	Gin *gin.Engine
	Mgo *mongo.Client
}
