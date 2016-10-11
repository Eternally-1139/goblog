package models

import (
"github.com/astaxie/beego/orm"
)

type User struct {
	Id            int64
	Mobile        string           `orm:"null"`
	Name          string           `orm:"null"`
	Password      string           `orm:"null"`
}

func (this *User) FindByMobile() error {
	return orm.NewOrm().QueryTable("user").Filter("mobile", this.Mobile).RelatedSel().One(this)
}
func (this *User) FindByUserName() error {
	return orm.NewOrm().QueryTable("user").Filter("name", this.Name).RelatedSel().One(this)
}

func (this *User) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (this *User) Read() error {
	return orm.NewOrm().QueryTable("user").Filter("id", this.Id).RelatedSel().One(this)
}

func (this *User) Update() (int64, error) {
	return orm.NewOrm().Update(this)
}

