package router

import (
	"github.com/gin-gonic/gin"
	"tahaNGU/api/handler"
)

func Test(c *gin.RouterGroup) {
	Health := handler.NewHealthStruct()
	c.GET("ping", Health.Ping)
	c.GET("head", Health.GetHead)
	c.POST("body", Health.GetBody)
	c.GET("query", Health.GetQuery)
}
