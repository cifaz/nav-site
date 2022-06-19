package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HtmlHandler struct{}

func NewHtmlHandler() *HtmlHandler {
	return &HtmlHandler{}
}

func (h *HtmlHandler) RedirectIndex(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
	return
}

func (h *HtmlHandler) Index(c *gin.Context) {
	c.Header("content-type", "text/html;charset=utf-8")
	c.String(http.StatusOK, string(HtmlIndex))
	//c.SetCookie("test-domain", "abc", 60000, "/", "test.com", false, true)
}

func (h *HtmlHandler) Favicon(c *gin.Context) {
	file, _ := Static.ReadFile("static/favicon.ico")
	c.Data(
		http.StatusOK,
		"image/x-icon",
		file,
	)

}
