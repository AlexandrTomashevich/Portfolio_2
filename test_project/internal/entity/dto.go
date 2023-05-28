package entity

type NameDTO struct {
	Name string `db:"name"`
}

type FriendDTO struct {
	Id   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Age  string `db:"age" json:"age"`
}
