package server

import (
	"embed"
	"errors"
	"github.com/gin-gonic/gin"
	"io/fs"
	conf "nav-site-server/config"
	"nav-site-server/extend/util"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

//go:embed static/index.html
var HtmlIndex []byte

//go:embed static/favicon.ico
var Favicon []byte

//go:embed static
var Static embed.FS

type HtmlResource struct {
	fs   embed.FS
	path string
}

func NewHtmlResource() *HtmlResource {
	return &HtmlResource{
		fs:   Static,
		path: "static",
	}
}

func (h *HtmlResource) Open(name string) (fs.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	fullName := filepath.Join(h.path, filepath.FromSlash(path.Clean("/static/"+name)))
	file, err := h.fs.Open(fullName)
	return file, err
}

func InitHtmlResource(engine *gin.Engine) *gin.Engine {
	engine.StaticFS("/static", http.FS(NewHtmlResource()))
	return engine
}

func InitEmbedResource(engine *gin.Engine) {
	js, _ := fs.Sub(Static, "static/js")
	css, _ := fs.Sub(Static, "static/css")
	engine.StaticFS("/js", http.FS(js))
	engine.StaticFS("/css", http.FS(css))

	if len(conf.App.Static.Root) > 0 {
		engine.Static("/data", conf.App.Static.Root+"/data")
	} else {
		engine.Static("/data", "./data")
	}

	htmlHandler := NewHtmlHandler()
	engine.GET("/", htmlHandler.Index)

	exists, _ := util.FileExists("./conf/favicon.ico")
	if exists {
		engine.StaticFile("/favicon.ico", "./conf/favicon.ico")
	} else {
		engine.GET("/favicon.ico", htmlHandler.Favicon)
	}
}

func InitLocalResource(engine *gin.Engine) {
	// 以下代码为读取本地文件, 目前使用embed方式打包, 如果代码量大请使用此种方式
	engine.LoadHTMLGlob("static/*.html")
	engine.Static("/static", "./static")
	engine.Static("/css", "./static/css")
	engine.Static("/js", "./static/js")
	engine.StaticFile("/favicon.ico", "favicon.ico")
	engine.StaticFile("/background.jpg", "./static/background.jpg")
	//
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}
