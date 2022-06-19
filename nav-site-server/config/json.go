package conf

import "github.com/gin-gonic/gin"

type JsonOutput struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Debug   string      `json:"debug"`
}

type JsonOutputList struct {
	Total    int `json:"total"`
	Rows     interface{}
	PageSize int
	PageNo   int
}

func HandleNotFound(c *gin.Context) {
	handleErr := JsonOutput{}
	handleErr.Message = "not found : " + c.Request.Method + " " + c.Request.URL.String()
	c.JSON(handleErr.Code, handleErr)
	return
}
