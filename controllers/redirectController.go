package controllers

import (
	"strings" 	
	"net/http"
	"url_shortener/models"
	u "url_shortener/utils"

	"github.com/gorilla/mux"
)

//RedirectShortURL gets the shortened url and redirects to the original url
var RedirectShortURL = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["shortenedURL"]
	
	if shortURL == "" {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	if strings.Contains(shortURL, "."){
		u.Respond(w, u.Message(false, "Invalid URL"))
		return
	}

	redirectURL, err := models.GetURLByShortenedURL(shortURL)

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid URL"))
		return
	}

	err = redirectURL.UpdateViews(redirectURL.Views + 1)
	if err != nil {
		u.Respond(w, u.Message(false, "Internal error"))
		return
	}

	//For testing purposes
	// fmt.Println(redirectURL.URL)
	// fmt.Println(redirectURL.Views)

	http.Redirect(w, r, redirectURL.URL, http.StatusTemporaryRedirect)
}
