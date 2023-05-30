package controller

import (
	"amigozone/db"
	"amigozone/model"
	"amigozone/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	User := utils.GetCurrentUser(w, r)
	Posts := db.GetAllPosts()
	tmpl, err := template.New("main.html").Funcs(template.FuncMap{
		"arr": func(values ...interface{}) []interface{} {
			return values
		},
	}).ParseFiles(
		"templates/parts/common.html",
		"templates/main.html",
		"templates/parts/navbar.html",
		"templates/parts/postForm.html",
		"templates/parts/post.html",
	)
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(w, "common.html", map[string]any{
		"title": "Главная",
		"user":  User,
		"posts": Posts,
	})

	if err != nil {
		panic(err)
	}
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	User := utils.GetCurrentUser(w, r)
	userIdString := mux.Vars(r)["id"]
	userId, err := strconv.ParseInt(userIdString, 10, 32)
	if err != nil {
		panic(err)
	}
	Posts := db.GetPostsByUser(int32(userId))

	tmpl, err := template.New("main.html").Funcs(template.FuncMap{
		"arr": func(values ...interface{}) []interface{} {
			return values
		},
	}).ParseFiles(
		"templates/parts/common.html",
		"templates/userPosts.html",
		"templates/parts/navbar.html",
		"templates/parts/postForm.html",
		"templates/parts/post.html",
	)
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(w, "common.html", map[string]any{
		"title": "Главная",
		"user":  User,
		"posts": Posts,
	})
}

func EditPost(w http.ResponseWriter, r *http.Request) {
	User := utils.GetCurrentUser(w, r)
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		panic(err)
	}

	Post := db.GetPostById(int32(id))

	tmpl, err := template.ParseFiles(
		"templates/parts/common.html",
		"templates/changePost.html",
		"templates/parts/navbar.html",
		"templates/parts/postForm.html",
		"templates/parts/post.html",
	)
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(w, "common.html", map[string]any{
		"title":    "Главная",
		"user":     User,
		"post":     Post,
		"isChange": true,
	})
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	utils.GetCurrentUser(w, r)
	var post model.Post
	idStr := r.FormValue("userId")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		panic(err)
	}
	post.User.Id = int32(id)
	post.Title = r.FormValue("title")
	post.Text = r.FormValue("text")
	db.CreatePost(&post)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ChangePost(w http.ResponseWriter, r *http.Request) {
	utils.GetCurrentUser(w, r)
	var post model.Post
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		panic(err)
	}
	post.Id = int32(id)
	post.Title = r.FormValue("title")
	post.Text = r.FormValue("text")
	db.ChangePost(&post)
	http.Redirect(w, r, "/post/change/"+idStr, http.StatusSeeOther)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	utils.GetCurrentUser(w, r)
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		panic(err)
	}
	db.DeletePost(int32(id))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
