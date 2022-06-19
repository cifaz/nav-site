package server

import (
	"github.com/gin-gonic/gin"
	"nav-site-server/api"
)

func Router(gin *gin.Engine) {
	website(gin)
	upload(gin)
	auth(gin)
}

func website(gin *gin.Engine) {
	//gin.GET("/", api.WebSiteList)
	website := gin.Group("/api/website")
	{
		website.GET("/list", api.WebSiteList)
		website.POST("/add", api.WebSiteAdd)
		website.POST("/update", api.WebSiteUpdate)
		website.POST("/delete", api.WebSiteDelete)
		website.GET("/groups", api.WebsiteGroups)
		website.PUT("/order/list", api.WebSiteOrder)
		website.PUT("/order/group", api.WebSiteGroupOrder)
	}

}

func upload(r *gin.Engine) {
	upload := r.Group("/api/upload")
	{
		upload.POST("/image", api.Image)
	}
}

func auth(r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", api.Login)
	}
}
