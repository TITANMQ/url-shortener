package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
	"url_shortener/models"
	u "url_shortener/utils"

	"github.com/gorilla/mux"
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

//URLStatsHandler handles URL stats requests
var URLStatsHandler = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["shortenedURL"]

	if shortURL == "" {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	if strings.Contains(shortURL, ".") {
		u.Respond(w, u.Message(false, "Invalid URL"))
		return
	}

	redirectURL, err := models.GetURLByShortenedURL(shortURL)

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid URL"))
		return
	}
	
	fmt.Fprintf(w, "<h1>Shortened URL stats</h1>" ) 
	fmt.Fprintf(w, "<p>Redirect URL: %s</p>", redirectURL.URL)
	fmt.Fprintf(w, "<p>Shortened URL: %s/%s</p>", r.Host , redirectURL.ShortURL)
	fmt.Fprintf(w, "<p>Total views: %d</p>", redirectURL.Views)
}
