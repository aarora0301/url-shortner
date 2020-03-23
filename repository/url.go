package repository

import (
	"fmt"
	"github.com/poc/url-shortner/cassandra"
	"github.com/poc/url-shortner/models"
)

func CreatUrl(url models.Url) {
	fmt.Println(" **** Creating new url ****\n")
	if err := cassandra.Session.Query("INSERT INTO url(hash, creation_date, expiration_date, original_url) VALUES(?, ?, ?, ?)",
		url.Hash, url.CreationDate, url.ExpirationDate, url.OriginalUrl).Exec(); err != nil {
		fmt.Println("Error while inserting  url")
		fmt.Println(err)

	}
}

func GetOriginalURL(key string) []models.Url {
	fmt.Println(" **** Fetching url ****\n")
	m := map[string]interface{}{}
	var url []models.Url
	urlQuery := cassandra.Session.Query("Select * from url where hash= ?", key).Iter()
	for urlQuery.MapScan(m) {
		url = append(url, models.Url{
			OriginalUrl: m["original_url"].(string),
		})
	}
	fmt.Println("Error while inserting  url")
	return url
}
