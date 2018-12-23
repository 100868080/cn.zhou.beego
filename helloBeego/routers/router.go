package routers

import (
	"2018/cn.zhou.beego/helloBeego/controllers"
	"github.com/astaxie/beego"
)

//路由－－访问路径，根据不同的访问路径交给(调用)不同的Controller处理并给与响应
//路由接受请求，Controller处理请求并给与响应
func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/zhou",&controllers.ZhouController{})
}
