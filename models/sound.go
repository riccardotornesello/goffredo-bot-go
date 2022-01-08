package models

type Sound struct {
	Id      int
	Name    string `orm:"size(100)"`
	SoundId string `orm:"size(100)"`
	Ready   bool
}
