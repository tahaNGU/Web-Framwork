package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"tahaNGU/api/router"
	"tahaNGU/api/validation"
)

func InitApi() {
	Server := gin.New()
	Server.Use(gin.Recovery(), gin.Logger())
	validator, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		validator.RegisterValidation("mobile", validation.IranianMobileNumberValidation, true)
	}
	v1 := Server.Group("/api/v1")
	router.Test(v1)
	Server.Run(":5001")
}
