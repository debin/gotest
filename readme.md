package main

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world~go")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run(":777")
}
