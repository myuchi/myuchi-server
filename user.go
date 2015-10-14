package main

import (
	"net/http"
	"path/filepath"

	"github.com/myuchi/myuchi-server/response"
)

var rootUser = &dataUser{
	"root",
	"root",
	"",
	true,
}

type dataUser struct {
	ID       string `json:"-"`
	Name     string `json:"name"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func userPath(id string) string {
	return filepath.Join(database, "users", id+".json")
}

func loadUser(id string) (*dataUser, error) {
	var user *dataUser
	if err := loadData(userPath(id), user); err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}

func (u *dataUser) saveUser() error {
	if u == nil {
		return nil
	}
	return saveData(userPath(u.ID), u)
}

func idRule(id string) bool {
	if id == "" {
		return false
	}
	if len(id) > 16 {
		return false
	}
	for _, c := range id {
		if !(('a' <= c && c <= 'z') || ('0' <= c && c <= '9') || ('_' == c)) {
			return false
		}
	}
	return true
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		userPostHandler(w, r)
		return
	case "GET":
		userGetHandler(w, r)
		return
	}
}

func userPostHandler(w http.ResponseWriter, r *http.Request) {
	_, root := authentication(r)
	if root != rootUser {
		return
	}
	r.ParseForm()
	id := r.FormValue("id")
	name := r.FormValue("name")
	password, passerr := authGenerate(r.FormValue("password"))
	if !idRule(id) || name == "" || passerr != nil {
		return
	}
	user := &dataUser{ID: id, Name: name, Password: password, IsAdmin: false}
	user.saveUser()
	return
}

func userGetHandler(w http.ResponseWriter, r *http.Request) {
	_, user := authentication(r)
	if user != nil {
		return
	}
	resp := response.User{
		ID:      user.ID,
		Name:    user.Name,
		IsAdmin: user.IsAdmin,
	}
	writeResponse(w, &resp)
	return
}
