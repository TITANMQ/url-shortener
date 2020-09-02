package controllers

import (
	"fmt"
	"url_shortener/models"
	"net/http"
	u "url_shortener/utils"
	"github.com/gorilla/mux"
)

//RedirectShortURL gets the shortened url and redirects to the original url
var RedirectShortURL = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["shortenedURL"]

	if shortURL == ""{
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	redirectURL, err := models.GetURLByShortenedURL(shortURL)

	if err != nil{
		u.Respond(w, u.Message(false, "Invalid URL"))
		return
	}
	fmt.Println(redirectURL.Url)
	http.Redirect(w, r, redirectURL.Url, http.StatusPermanentRedirect)
}
