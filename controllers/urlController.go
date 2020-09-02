package controllers

import (
	"encoding/json"
	"net/http"
	"url_shortener/models"
	u "url_shortener/utils"
)

//CreateURLHandler handles create url requests
var CreateURLHandler = func(w http.ResponseWriter, r *http.Request) {

	url := &models.URL{}
	err := json.NewDecoder(r.Body).Decode(url)

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
	}

	resp := url.CreateURL()
	resp["data"] = url
	u.Respond(w, resp)
}
