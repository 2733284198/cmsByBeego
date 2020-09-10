package controllers

import (
	"cmsByBeego/models"
	"github.com/astaxie/beego"
	"strings"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController)Index()  {
	if c.Ctx.Request.Method="POST"{
		userkey:=strings.TrimSpace(c.GetString("username"))
		password:=strings.TrimSpace(c.GetString("password"))

		if len(userkey)>0&&len(password)>0{
			password :=utils.Md5([]byte(password))
			user:=models.GetUserByName(userkey)
			if password ==user.Password{
				c.SetSession("xcmsuser",user)
				c.Redirect("/",302)
				c.StopRun()
			}
		}
	}
	c.TplName="login/index.html"
}