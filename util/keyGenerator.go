package util

import (
	"encoding/base64"
	"log"
	"strconv"
	"time"
)

var KeySet []string

/**
Generates Unique keys
*/
func GetKeys() {
	finished := make(chan [] string)
	latest := time.Now().Unix()
	/**
	use go routine to generate keys
	*/
	go generateKey(finished, latest)
	log.Println("processing")
	<-finished
	return
}

/**
Generates 130K unique keys in a second
*/
func generateKey(finished chan [] string, latest int64) {
	var i int64
	log.Println("worker started")
	for i = 0000000000000000; i < 130000; i++ {
		str := strconv.FormatInt(latest+i, 31)
		encoded := base64.StdEncoding.EncodeToString([]byte(str))
		KeySet = append(KeySet, encoded)
	}
	log.Println("worker processed")
	finished <- KeySet
}
