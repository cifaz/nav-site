package api

import (
	"encoding/json"
	"io/ioutil"
	conf "nav-site-server/config"
	"nav-site-server/extend/util"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestLoginParams struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

func Login(c *gin.Context) {
	app := conf.App
	output := conf.JsonOutput{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		output.Debug = err.Error()
		output.Code = conf.ParamsInvalid
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.ParamsInvalid]
		response(c, output)
		return
	}
	data := RequestLoginParams{}
	if err := json.Unmarshal(body, &data); err != nil {
		output.Debug = err.Error()
		output.Code = conf.AuthLoginAccountInvalid
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.AuthLoginAccountInvalid]
		response(c, output)
		return
	}

	if checkAccount(&data) == false {
		output.Debug = "account or password invalid"
		output.Code = conf.AuthLoginAccountInvalid
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.AuthLoginAccountInvalid]
		response(c, output)
		return
	}
	j := util.JWT{}
	token, err := j.Make(data.Name, app.Account.Secret, app.Account.CookieExpireSeconds)
	if err != nil {
		output.Debug = err.Error()
		output.Code = conf.Error
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.Error]
		response(c, output)
		return
	}
	res := make(map[string]interface{})
	res["token"] = token
	res["expire"] = app.Account.CookieExpireSeconds
	output.Debug = ""
	output.Code = conf.Success
	output.Data = res
	output.Message = conf.ErrorMsg[conf.Success]
	response(c, output)
	return
}

func checkAccount(r *RequestLoginParams) bool {
	members := conf.App.Account.Members
	for _, user := range members {
		if user.Password == r.Pwd && user.Name == r.Name {
			return true
		}
	}
	return false
}

func checkAuth(name, methodType string) bool {
	members := conf.App.Account.Members
	admin := conf.App.Account.Admin
	for _, user := range members {
		if user.Name == admin {
			return true
		}
		if user.Name == name {
			auths := strings.Split(user.Rule, ",")
			for _, auth := range auths {
				return auth == methodType
			}
		}
	}
	return false
}
