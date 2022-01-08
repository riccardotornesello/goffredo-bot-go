package models

type Sound struct {
	Id     int
	Name   string `orm:"size(100)"`
	UserId string `orm:"size(100)"`
	Ready  bool
}
