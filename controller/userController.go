package controller

import (
	"amigozone/db"
	"amigozone/model"
	"amigozone/serv"
	"net/http"
	"text/template"
	"time"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles(
			"templates/parts/common.html",
			"templates/register.html",
			"templates/parts/userForm.html",
			"templates/parts/navbar.html")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, map[string]any{
			"title":   "Registration",
			"isLogin": false,
		})
	case "POST":
		var user model.User
		user.Name = r.FormValue("name")
		user.Password = r.FormValue("password")
		user.Email = r.FormValue("email")
		serv.MailSender(user.Email, "Hello", "Hello")
		db.CreateUser(&user)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles(
			"templates/parts/common.html",
			"templates/login.html",
			"templates/parts/userForm.html",
			"templates/parts/navbar.html",
		)
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, map[string]any{
			"title":   "Login",
			"isLogin": true,
		})
	case "POST":
		password := r.FormValue("password")
		email := r.FormValue("email")
		user := db.GetUserByEmailAndPassword(email, password)
		sessionId := uuid.New().String()
		db.CreateSession(user.Id, sessionId)
		expiration := time.Now().Add(5 * time.Minute)
		cookie := http.Cookie{Name: "session-id", Value: sessionId, Expires: expiration}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "session-id", Value: "", Expires: time.Now()}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
