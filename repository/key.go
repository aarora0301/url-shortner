package repository

import (
	"errors"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/poc/url-shortner/cassandra"
	"github.com/poc/url-shortner/models"
	util2 "github.com/poc/url-shortner/util"
	_ "github.com/tokopedia/cm/util"
	"log"
)

/**
This should be a different microservice managing the lifecycle of keys only

*/

/***
Creates keys in database : it should be a standalone service
*/
func CreateKeys() (err error) {
	fmt.Println(" **** Creating new key ****\n")
	util2.GetKeys()
	var target [] string
	target = util2.KeySet
	var i int
	for i = 0; i < len(target); i++ {
		if err := cassandra.Session.Query("INSERT INTO available_key(key) VALUES(?)",
			target[i]).Exec(); err != nil {
			message := "Error while inserting key: " + target[i] + error.Error(err)
			log.Println(message)
			err = errors.New(message)
		}
	}
	return
}

/**
Get first available key from database
*/
func GetAvailableKey() ([]models.Key, error) {
	fmt.Println("Get key")
	var Key []models.Key
	m := map[string]interface{}{}
	var err error
	var available *gocql.Iter
	available = cassandra.Session.Query("select * from available_key limit 1").Iter()
	if available.NumRows() <= 0 {
		message := "No Record exist, Generate Keys first"
		log.Println(message)
		err = errors.New(message)
		return nil, err
	}
	for available.MapScan(m) {
		Key = append(Key, models.Key{
			Key: m["key"].(string),
		})
	}
	if len(Key) <= 0 {
		message := "Error in parsing database rows"
		log.Println(message)
		err = errors.New(message)
		return nil, err
	}
	moveKeyToUsed(Key[0].Key)
	return Key, err
}

/**
  Whenever key is fetched from database two actions will occur:
   1. Remove key from available pool of keys
   2. Add key in used pool of keys
*/
func moveKeyToUsed(key string) {
	log.Println("Move key from available pool to used pool")
	deleteAvailableKey(key)
	insertInUsedKey(key)

}

/**
Delete key from available keys
*/
func deleteAvailableKey(key string) (err error) {
	log.Println("Deleting key from available pool")
	if err := cassandra.Session.Query("DELETE FROM available_key  WHERE key = ?", key).Exec(); err != nil {
		message := "Error while deleting key: " + key + " " + error.Error(err)
		log.Println(message)
		err = errors.New(message)
	}
	return
}

/**
Delete key from used keys so that it can be reused later
*/
func deleteUsedKey(key string) (err error) {
	log.Println("Delete key from used pool ")
	if err := cassandra.Session.Query("DELETE FROM used_key  WHERE key = ?", key).Exec(); err != nil {
		message := "Error while deleting used key: " + key + " " + error.Error(err)
		log.Println(message)
		err = errors.New(message)
	}
	return
}

/**
Insert key in used keys whenever a key is accessed from available keys
*/
func insertInUsedKey(key string) (err error) {
	log.Println("Insert key in used_key")
	if err := cassandra.Session.Query("INSERT INTO used_key(key) VALUES(?)",
		key).Exec(); err != nil {
		message := "Error while inserting key in used keys: " + key + " " + error.Error(err)
		log.Println(message)
		err = errors.New(message)
	}
	return
}
