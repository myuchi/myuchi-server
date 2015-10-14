package main

import (
	"net/http"

	"github.com/myuchi/myuchi-server/response"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	resp := response.Status{"myuchi"}
	writeResponse(w, resp)
}
