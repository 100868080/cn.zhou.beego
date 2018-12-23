package routers

import (
	"2018/cn.zhou.beego/love_home/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
