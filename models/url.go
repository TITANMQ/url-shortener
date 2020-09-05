package models

import (
	"fmt"
	"math/rand"
	"time"
	u "url_shortener/utils"
)

//URL holds URL data
type URL struct {
	ID       uint   `json:"id" gorm:"column:id;PRIMARY_KEY;AUTO_INCREMENT"`
	URL      string `json:"url" gorm:"column:url"`
	ShortURL string `json:"short_url" gorm:"column:short_url"`
	Views 	 uint 	`json:"views" gorm:"column:views"`
}

//CreateURL creates a url entry to the database
func (url *URL) CreateURL() map[string]interface{} {
	url.ShortURL = GenerateShortenedURL()
	url.Views = 0
	err := GetDB().Create(&url).Error

	if err != nil {
		return u.Message(false, "Database Error")
	}
	return u.Message(true, "URL created successfully")
}

//ShortenedURLLength length of the shortend URL
var ShortenedURLLength = 6
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

//GenerateShortenedURL generates a shortened url
func GenerateShortenedURL() string {
	length := 6
	url := make([]rune, length)
	rand.Seed(time.Now().UnixNano())
	for i := range url {
		url[i] = letters[rand.Intn(len(letters))]
	}
	fmt.Println(string(url))
	return string(url)
}

//GetURLByID gets a url from the database by ID
func GetURLByID(id uint) (*URL, error) {
	url := &URL{}
	err := GetDB().Where("id=?", id).First(url).Error
	if err != nil {
		return url, err
	}
	return url, nil
}

// GetURLByShortenedURL gets a URL by shortened URL
func GetURLByShortenedURL(shortenedURL string) (*URL, error) {
	url := &URL{}
	err := GetDB().Where("short_url=?", shortenedURL).First(url).Error
	if err != nil {
		return url, err
	}
	return url, nil
}

//UpdateViews updates views for a URL
func (url *URL)UpdateViews(views uint)error{
	url.Views = views
	err := GetDB().Save(&url).Error
	return err
}
