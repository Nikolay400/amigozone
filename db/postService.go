package db

import (
	"amigozone/model"
)

func CreatePost(post *model.Post) {
	_, err := Db.Query("INSERT INTO post(title, text, user_id) VALUES($1,$2,$3)", post.Title, post.Text, post.User.Id)
	if err != nil {
		panic(err)
	}
}

func ChangePost(post *model.Post) {
	_, err := Db.Query("UPDATE post SET title=$1, text=$2 WHERE id=$3", post.Title, post.Text, post.Id)
	if err != nil {
		panic(err)
	}
}

func DeletePost(id int32) {
	_, err := Db.Query("DELETE FROM post WHERE id=$1", id)
	if err != nil {
		panic(err)
	}
}

func GetPostById(id int32) (post model.Post) {
	res, err := Db.Query("SELECT p.id, p.title, p.text, u.id, u.name, u.email FROM post p INNER JOIN usr u ON p.user_id=u.id WHERE p.id=$1", id)
	if err != nil {
		panic(err)
	}
	defer res.Close()
	if res.Next() {
		res.Scan(&post.Id, &post.Title, &post.Text, &post.User.Id, &post.User.Name, &post.User.Email)
	}
	return post
}

func GetAllPosts() (posts []model.Post) {
	res, err := Db.Query("SELECT p.id, p.title, p.text, u.id, u.name, u.email FROM post p INNER JOIN usr u ON p.user_id=u.id")
	if err != nil {
		panic(err)
	}
	defer res.Close()
	for res.Next() {
		var post model.Post
		res.Scan(&post.Id, &post.Title, &post.Text, &post.User.Id, &post.User.Name, &post.User.Email)
		posts = append(posts, post)
	}
	return posts
}

func GetPostsByUser(userId int32) (posts []model.Post) {
	res, err := Db.Query("SELECT p.id, p.title, p.text, u.id, u.name, u.email FROM post p INNER JOIN usr u ON p.user_id=u.id WHERE u.id=$1", userId)
	if err != nil {
		panic(err)
	}
	defer res.Close()
	for res.Next() {
		var post model.Post
		res.Scan(&post.Id, &post.Title, &post.Text, &post.User.Id, &post.User.Name, &post.User.Email)
		posts = append(posts, post)
	}
	return posts
}

func CreatePostTable() {
	_, err := Db.Query("CREATE TABLE post (" +
		"id bigserial not null," +
		"title varchar(255)," +
		"text varchar(2000)," +
		"user_id int8," +
		"primary key (id))")
	if err != nil {
		panic(err)
	}
}
