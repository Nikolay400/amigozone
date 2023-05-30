package model

type Post struct {
	Id    int32
	Title string
	Text  string
	User  User
}
