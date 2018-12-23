package models

import "github.com/astaxie/beego/orm"
import _"github.com/go-sql-driver/mysql"


type User struct {
	Id int		//id默认是自增长主键,否则用 `orm:pk,auto`来设置自增长主键
	Name string `orm:size(100)`  //设置表的属性,等同于varchar(100)
	Sex_beego string `orm:size(100)`
}


func init() {
	//注册数据库（具体的哪一个数据库）
	orm.RegisterDataBase("default",
		"mysql","root:zhou123@tcp(localhost:3306)/" +
		"zhou?charset=utf8")

	/*
	  注册数据库驱动
	  默认已经注册的三个数据库之一，可以不用注册
	 */
	 orm.RegisterDriver("mysql",orm.DRMySQL)


	orm.RegisterModel(new(User))  //注册表

	/*
		参数
		1:将要操作的数据库(数据库别名)
		2:　true：drop table 后再建表　(发送drop语句)
			false： 跳过此操作 (不发送drop语句)
		3:true：打印执行过程 (打印sql语句)
			false： 跳过此操作
	 */
	 orm.RunSyncdb("default",false,true)  //创建表

}


