package main

import (
	"github.com/astaxie/beego"
	//"order/controllers/errorpage"
	"test/models"
	_ "test/routers"
	"os"
)

func init() {
	models.Connect()
	initArgs()
}

func main() {
	//beego.ErrorController(&errorpage.ErrorController{})
	beego.Run()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncbd(false)
			os.Exit(0)
		}
		if v == "-syncdb-force" {
			models.Syncbd(true)
			os.Exit(0)
		}
	}
}
