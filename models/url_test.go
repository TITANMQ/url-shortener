package models

import "testing"

func TestGenerateShortenedURL(t *testing.T) {
	shortenedURL := GenerateShortenedURL()

	if shortenedURL == ""{
		t.Errorf(`GenerateShortenedURL() returned "%v" an empty string`, shortenedURL)
	}

	if len(shortenedURL) != ShortenedURLLength{
		t.Errorf(`GenerateShortenedURL() returned "%v" and expected a string of length %v`, shortenedURL,ShortenedURLLength )
	}
}
