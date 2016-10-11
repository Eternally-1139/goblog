package routers

import (
	"test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    //beego.Router("/user/register", &controllers.ajaxController{},"Register")
	beego.Include(&controllers.AjaxController{},
		&controllers.MainController{})

}
