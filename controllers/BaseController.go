package controllers

import (
	"cmsByBeego/models"
	"fmt"
	"github.com/astaxie/beego"
	"go/ast"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (c *BaseController)Prepare()  {
	//附值
	c.controllerName,c.actionName=c.GetControllerAndAction()
	beego.Informational(c.controllerName,c.actionName)
	//todo 保存用户数据
	user:=c.auth()

	fmt.Print("beego:perpare:"+c.controllerName+","+c.actionName)
	c.Data["Menu"]=models.MenuTreeStuct(user)
}

// 设置模板
// 第一个参数模板，第二个参数
func (c*BaseController)setTpl(template ...string)  {
	var tplName string
	layout:="common/layout.html"
	switch {
	case len(template)==1:
		tplName=template[0]
	case len(template)==2:
		tplName=template[0]
		layout=template[1]
	default:
		//不要"Controller"这个10个字母
		ctrlName:=strings.ToLower(c.controllerName[0:len(c.controllerName)-10])
		actionName:=strings.ToLower(c.actionName)
		tplName=ctrlName+"/"+actionName+".html"
	}

	_,found:=c.Data["Footer"]
	if !found{
		c.Data["Footer"]="menu/footerjs.html"
	}
	c.LayoutSections["layout"]=layout
  c.TplName=tplName


}

func (c *BaseController)jsonResult(code string,msg string,obj interface{})  {

	r:=&JSONS{code,msg,obj}
	c.Data["json"]=r
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController)listJsonResult(code string,msg string,count int64,obj interface{})  {
	r:=&JSONS{code,msg,count,obj}
	c.Data["json"]=r
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController)auth()models.UserModel  {
	user:=c.GetSession("xcmsuser")
	if user==nil{
		c.Redirect("/login",302)
		c.StopRun()
		return models.UserModel{}
	}else {
		return user.(models.UserModel)
	}
}