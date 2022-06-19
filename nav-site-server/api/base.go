package api

import (
	"github.com/gin-gonic/gin"
	conf "nav-site-server/config"
	"net/http"
	"strings"
)

func response(c *gin.Context, r conf.JsonOutput) {
	c.JSON(http.StatusOK, r)
	return
}

func getAuthorization(c *gin.Context) string {
	return strings.ReplaceAll(c.GetHeader("authorization"), "Bearer ", "")
}
