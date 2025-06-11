package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type HotdogType struct {
	Id          int            `orm:"auto;pk"`
	Name        string         `orm:"unique"`
	Description string         `orm:"type(text)"`
	Price       float64        `orm:"digits(10);decimals(2);index"`
	Inventory   []*HotdogStock `orm:"reverse(many)"`
}

type Stand struct {
	Id          int            `orm:"auto;pk"`
	Name        string         `orm:"unique;index"`
	Address     string         `orm:"type(text)"`
	Description string         `orm:"type(text)"`
	Inventory   []*HotdogStock `orm:"reverse(many)"`
}

type HotdogStock struct {
	Id       int         `orm:"auto;pk"`
	Quantity int         `orm:"default(0)"`
	Hotdog   *HotdogType `orm:"rel(fk);on_delete(cascade);index"`
	Stand    *Stand      `orm:"rel(fk);on_delete(cascade);index"`
}

func (u *Stand) TableName() string {
	return "hotdog_stand"
}

// Unused is a model that is not registered with Beego ORM.
type Unused struct {
	Id int `orm:"auto;pk"`
}

func init() {
	orm.RegisterModel(new(HotdogType))
}

func init() {
	orm.RegisterModel(new(Stand))
	x()
}

func x() {
	orm.RegisterModel(new(HotdogStock))
}
