package controllers

import (
	"github.com/astaxie/beego"
)

//处理"/"的请求,求并给与响应
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	//2018/cn.zhou.beego/helloBego/views/
 }

//处理"/zhou"的请求,求并给与响应
type  ZhouController struct {
	beego.Controller
}

func (c *ZhouController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.Ctx.WriteString("helo,beegon")
	c.Controller.Data["ok"]="are you ok??"
}
