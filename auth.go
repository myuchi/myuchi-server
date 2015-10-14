package main

import (
	"encoding/base64"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func authGenerate(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hash), nil
}

func authCompare(password, hash string) error {
	hashed, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(hashed, []byte(password))
}

func authentication(r *http.Request) (bool, *dataUser) {
	id, password, ok := r.BasicAuth()
	if !ok {
		return false, nil
	}
	if id == "root" {
		if password == config.RootPassword {
			return true, rootUser
		}
		return true, nil
	}
	user, err := loadUser(id)
	if err != nil {
		return true, nil
	}
	if err := authCompare(password, user.Password); err != nil {
		return true, nil
	}
	return true, user
}
