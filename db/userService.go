package db

import (
	"amigozone/model"
)

func GetUserById(id int32) model.User {
	query, err := Db.Query("SELECT name FROM usr WHERE id=$1", id)
	if err != nil {
		panic(err)
	}
	defer query.Close()
	var user model.User
	if query.Next() {
		err = query.Scan(&user.Name)
		if err != nil {
			panic(err)
		}
	}
	return user
}

func GetUserByEmailAndPassword(email string, password string) model.User {
	query, err := Db.Query("SELECT * FROM usr WHERE email=$1 AND password=$2", email, password)
	if err != nil {
		panic(err)
	}
	defer query.Close()
	var user model.User
	if query.Next() {
		err = query.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			panic(err)
		}
	}
	return user
}

func CreateUser(user *model.User) {
	_, err := Db.Query("INSERT INTO usr(name, password, email) VALUES ($1, $2, $3)", user.Name, user.Password, user.Email)
	if err != nil {
		panic(err)
	}
}

func CreateSession(userId int32, sessionId string) {
	res, err := Db.Query("SELECT user_id FROM session WHERE user_id=$1", userId)
	if err != nil {
		panic(err)
	}
	defer res.Close()
	if res.Next() {
		_, err := Db.Query("UPDATE session SET session=$1 WHERE user_id=$2", sessionId, userId)
		if err != nil {
			panic(err)
		}
	} else {
		_, err := Db.Query("INSERT INTO session VALUES ($1, $2)", userId, sessionId)
		if err != nil {
			panic(err)
		}
	}
}

func GetUserBySession(sessionId string) (user model.User) {
	res, err := Db.Query("SELECT u.id, u.name, u.email From session s INNER JOIN usr u ON s.user_id=u.id  WHERE s.session=$1", sessionId)
	if err != nil {
		panic(err)
	}
	defer res.Close()
	if res.Next() {
		res.Scan(&user.Id, &user.Name, &user.Email)
	}
	return user
}

func CreateUserTable() {
	_, err := Db.Query("CREATE TABLE usr (" +
		"id bigserial not null," +
		"name varchar(255)," +
		"password varchar(255)," +
		"email varchar(255)," +
		"primary key (id))")
	if err != nil {
		panic(err)
	}
}

func CreateSessionTable() {
	_, err := Db.Query("CREATE TABLE session (" +
		"user_id integer," +
		"session varchar(255)," +
		"primary key (user_id, session))")
	if err != nil {
		panic(err)
	}
}
