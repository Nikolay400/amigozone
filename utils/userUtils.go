package utils

import (
	"amigozone/db"
	"amigozone/model"
	"net/http"
)

func GetCurrentUser(w http.ResponseWriter, r *http.Request) (user model.User) {
	sessionId, err := r.Cookie("session-id")
	if err == nil {
		user = db.GetUserBySession(sessionId.Value)
	}
	if user.Id == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	return user
}
