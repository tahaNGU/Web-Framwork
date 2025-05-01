package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthStruct struct {
}

type Person struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile"`
}

func NewHealthStruct() *HealthStruct {
	return &HealthStruct{}
}

func (h *HealthStruct) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *HealthStruct) GetHead(c *gin.Context) {
	person := Person{}
	c.BindHeader(&person)
	c.JSON(200, gin.H{
		"message": person,
	})
}

func (h *HealthStruct) GetBody(c *gin.Context) {
	person := Person{}
	error := c.ShouldBindJSON(&person)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validationError": error.Error(),
		})
		return
	}
	c.JSON(200, person)
}

func (h *HealthStruct) GetQuery(c *gin.Context) {
	c.JSON(200, gin.H{
		"firstName": c.Query("firstName"),
		"lastName":  c.Query("lastName"),
	})
}
