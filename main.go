package main

import (
	"amigozone/controller"
	"amigozone/db"
	"amigozone/serv"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db.Open()
	defer db.Db.Close()
	serv.AuthInit()

	r := mux.NewRouter()

	r.HandleFunc("/", controller.MainPage)
	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/logout", controller.Logout)
	r.HandleFunc("/registration", controller.Register)

	r.Path("/post").Methods("POST").HandlerFunc(controller.CreatePost)
	r.Path("/post/change/{id}").Methods("GET").HandlerFunc(controller.EditPost)
	r.Path("/post/change/{id}").Methods("POST").HandlerFunc(controller.ChangePost)
	r.Path("/post/delete/{id}").Methods("POST").HandlerFunc(controller.DeletePost)

	r.Path("/user/{id}").Methods("GET").HandlerFunc(controller.GetPosts)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe("", r)
}
