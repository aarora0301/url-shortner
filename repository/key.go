package repository

import (
	"fmt"
	"github.com/poc/url-shortner/cassandra"
	"github.com/poc/url-shortner/models"
	util2 "github.com/poc/url-shortner/util"
	_ "github.com/tokopedia/cm/util"
)

func CreateKeys() {
	fmt.Println(" **** Creating new key ****\n")
	util2.GetKeys()
	var target [] string
	target = util2.KeySet
	var i int
	for i = 0; i < len(target); i++ {
		if err := cassandra.Session.Query("INSERT INTO available_key(key) VALUES(?)",
			target[i]).Exec(); err != nil {
			fmt.Println("Error while inserting key")
			fmt.Println(err)
		}
	}
}

func GetAvailableKey() []models.Key {
	fmt.Println("Get key")
	var Key []models.Key
	m := map[string]interface{}{}
	available := cassandra.Session.Query("select * from available_key limit 1").Iter()
	for available.MapScan(m) {
		Key = append(Key, models.Key{
			Key: m["key"].(string),
		})
	}
	moveKeyToUsed(Key[0].Key)
	return Key
}

func moveKeyToUsed(key string) {
	fmt.Println("Move key")
	deleteAvailableKey(key)
	insertInUsedKey(key)

}
func deleteAvailableKey(key string) {
	fmt.Println("Deleting key")
	if err := cassandra.Session.Query("DELETE FROM available_key  WHERE key = ?", key).Exec(); err != nil {
		fmt.Println("Error while deleting Emp")
		fmt.Println(err)
	}
}

func deleteUsedKey(key string) {
	fmt.Println("Deleting key")
	if err := cassandra.Session.Query("DELETE FROM used_key  WHERE key = ?", key).Exec(); err != nil {
		fmt.Println("Error while deleting Emp")
		fmt.Println(err)
	}
}

func insertInUsedKey(key string) {
	fmt.Println("Insert key in used_key")
	if err := cassandra.Session.Query("INSERT INTO used_key(key) VALUES(?)",
		key).Exec(); err != nil {
		fmt.Println("Error while deleting Emp")
		fmt.Println(err)
	}
}
