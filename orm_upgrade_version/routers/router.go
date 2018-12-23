package routers

import (
	"2018/cn.zhou.beego/orm_upgrade_version/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
