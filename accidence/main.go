package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

var a=100

//每个　package 里面度可以有零个多个 init函数
//并且 init　函数在程序之前被调用，或者导入包之前调用一次
func init() {
fmt.Println("init 1")
a=500
}

func init() {
	fmt.Println("init 2")

}

func init() {
	fmt.Println("init 3")
a=999
}

func init() {
	a=654321
}


type  myint int
func (a myint) add(b int){

 fmt.Println(a+b)

}

func (a myint) cadd(b int){
 fmt.Println(a*b)

}

func main() {
	 beego.Info("hello,beego") //控制台输出hello,beego

	 //fmt.Println("a=",a)
	 f:=myint(231)
	 f .cadd(12)
	 f.add(100)

}
