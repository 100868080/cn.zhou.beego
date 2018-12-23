package main

import (
	"2018/cn.zhou.beego/orm/models"
	_ "2018/cn.zhou.beego/orm/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func judgeError(err error){
	if err!=nil {
		beego.Info(err)
		return
	}
}


//使用orm对mysql数据库 插入数据
func insert(){
	o:=orm.NewOrm() //获得orm句柄 (sql语句发送器)
	o.Using("default") //对目标数据库发送sql语句

	u:=new(models.User)
	u.Id=120
	u.Name="copper"
	u.Sex_beego="v1.0"

    i,err:=o.Insert(u)
	if err !=nil{
		beego.Info(err)
	}
	beego.Info("insert into database:",i)
}

//使用orm对mysql数据库　更新数据
func update(){
	o:=orm.NewOrm()
	o.Using("default")

	u:=new(models.User)
	u.Id=120
	u.Sex_beego="王八蛋"

	i,err:=o.Update(u)
	judgeError(err)
	beego.Info(i)

}

//使用orm对mysql数据库 查询数据
//注意：查询所返回的结果在给入的结构体里面，并没有另外返回一个变量
func query()  {
	o:=orm.NewOrm()
	u:=models.User{Id:120}

	 err:=o.Read(&u)
	 judgeError(err)
	 beego.Error(u)
	 beego.Info(u)

}

//使用orm对mysql数据库　删除数据
func delete(){
	o:=orm.NewOrm()
	u:=models.User{Id:120}
	i,err:=o.Delete(&u)
	judgeError(err)
	beego.Info(i)
}

func main() {
	 //insert()
	//update()
	 query()
	//delete()
	beego.Run()
}

