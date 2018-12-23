package main

import (
	_ "2018/cn.zhou.beego/helloBeego/routers"
	"github.com/astaxie/beego"
)

//项目的入口
func main() {
	beego.SetStaticPath("/","/home/zhou/workspace" +
		"/program/GO/work/src/2018/cn.zhou.beego/helloBeego/views")
	beego.SetStaticPath("/pictures","pictures")
	beego.Run()
}

