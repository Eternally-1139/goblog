package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["test/controllers:AjaxController"] = append(beego.GlobalControllerRouter["test/controllers:AjaxController"],
		beego.ControllerComments{
			Method: "UserRegister",
			Router: `/user/register`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["test/controllers:AjaxController"] = append(beego.GlobalControllerRouter["test/controllers:AjaxController"],
		beego.ControllerComments{
			Method: "UserLogin",
			Router: `/user/login`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["test/controllers:MainController"] = append(beego.GlobalControllerRouter["test/controllers:MainController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["test/controllers:MainController"] = append(beego.GlobalControllerRouter["test/controllers:MainController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/main`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["test/controllers:MainController"] = append(beego.GlobalControllerRouter["test/controllers:MainController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["test/controllers:MainController"] = append(beego.GlobalControllerRouter["test/controllers:MainController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["test/controllers:MainController"] = append(beego.GlobalControllerRouter["test/controllers:MainController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

	beego.GlobalControllerRouter["test/controllers:MainController"] = append(beego.GlobalControllerRouter["test/controllers:MainController"],
		beego.ControllerComments{
			Method: "School",
			Router: `/school`,
			AllowHTTPMethods: []string{"*"},
			Params: nil})

}
