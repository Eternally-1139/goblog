package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	//"test/controllers/base"
	"test/models"
	"test/service"
)

type AjaxController struct {
	beego.Controller
}

// @router /user/register [*]
func (this *AjaxController) UserRegister() {

	username := this.GetString("username")
	passwords := this.GetString("password")
	password := service.Strtomd5(passwords);
	mobile := this.GetString("mobile")
	result := make(map[string]interface{})
	var user models.User
	user.Name = username
	user.Password = password
	user.Mobile = mobile
	if _, err := user.Insert(); err != nil {
		result["status"] = 10002
		result["message"] = "custom insert error"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result["status"] = 10000
	result["message"] = "success"
	result["user"] = user
	this.Data["json"] = result
	this.ServeJSON()

}

// @router /user/login [*]
func (this *AjaxController) UserLogin() {

	username := this.GetString("username")
	password := this.GetString("password")
	result := make(map[string]interface{})
	user := models.User{Name: username}
	if err := user.FindByUserName(); err != nil {
		result["status"] = 10002
		result["message"] = "用户名或密码错误"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	if user.Password != service.Strtomd5(password) {
		result["status"] = 10002
		result["message"] = "用户名或密码错误"
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	if err := user.Read(); err != nil {
		beego.Error("user read error", err)
	}
	result["status"] = 10000
	result["message"] = "success"
	result["user"] = user
	this.Data["json"] = result
	this.SetSession("userinfo", user)
	beego.Debug("User Login:", user.Name)
	this.ServeJSON()

}

