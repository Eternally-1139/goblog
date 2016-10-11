
package base

import (
"github.com/astaxie/beego"
//"github.com/astaxie/beego/orm"
"test/models"
)

type BaseRouter struct {
	beego.Controller
	User          models.User
}

type CtrPrepare interface {
	CtrPrepare()
}

func (this *BaseRouter) ReturnJson(status int, message string, args ...interface{}) {

	result := make(map[string]interface{})
	result["status"] = status
	result["message"] = message

	key := ""

	for _, arg := range args {
		switch arg.(type) {
		case string:
			key = arg.(string)
		default:
			result[key] = arg
		}
	}

	this.Data["json"] = result
	this.ServeJSON()

}
