package models

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}
