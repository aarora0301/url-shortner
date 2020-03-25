package repository

import (
	"errors"
	"github.com/gocql/gocql"
	"github.com/poc/url-shortner/cassandra"
	"github.com/poc/url-shortner/models"
	"log"
)

/***
Create Mapping of URL and hash in database
*/
func CreatUrl(url models.Url) (err error) {
	log.Println(" **** Creating new url ****\n")
	if err := cassandra.Session.Query("INSERT INTO url(hash, creation_date, expiration_date, original_url) VALUES(?, ?, ?, ?)",
		url.Hash, url.CreationDate, url.ExpirationDate, url.OriginalUrl).Exec(); err != nil {
		message := "Error while inserting  url " + error.Error(err)
		log.Println(message)
		err = errors.New(message)
	}
	return
}

/**
Fetch URL for hash
*/
func GetOriginalURL(key string) (url []models.Url, err error) {
	log.Println(" **** Fetching url ****\n")
	m := map[string]interface{}{}
	var urlQuery *gocql.Iter
	urlQuery = cassandra.Session.Query("Select * from url where hash= ?", key).Iter()
	if urlQuery.NumRows() <= 0 {
		err = errors.New("No mapping exists for key: " + key)
		return
	}
	for urlQuery.MapScan(m) {
		url = append(url, models.Url{
			OriginalUrl: m["original_url"].(string),
		})
	}
	return
}
