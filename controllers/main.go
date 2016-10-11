package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"test/models"
)

type MainController struct {
	beego.Controller
}

// @router / [*]
func (this *MainController) Get() {
	userinfo := this.GetSession("userinfo")
	if userinfo != nil {
		this.Ctx.Redirect(302, "/main")
	}
	this.TplName = "login.tpl"

}
// @router /main [*]
func (this *MainController) Index(){
	var users []*models.User
	orm.NewOrm().QueryTable("user").All(&users)
	this.Data["users"] = users
	userinfo := this.GetSession("userinfo")
	if userinfo !=nil {
		this.Data["userinfo"] = userinfo
		this.TplName = "index.tpl"
	}else{
		this.Ctx.Redirect(302,"/login")
	}

}
// @router /logout [*]
func (this *MainController) Logout() {
	this.DelSession("userinfo")
	this.Ctx.Redirect(302, "/")
}

// @router /register [*]
func (c *MainController) Register() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "register.tpl"
}

// @router /login [*]
func (c *MainController) Login() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login.tpl"
}

// @router /school [*]
func (this *MainController) School() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "school.html"
}
