package util

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)

var KeySet []string

func GetKeys()  {
	finished := make(chan [] string)
	latest := time.Now().Unix()
	go generateKey(finished, latest)
	fmt.Println("processing")
	<-finished
	return
}

func generateKey(finished chan [] string , latest int64) {
	var i int64
	fmt.Println("worker started")
	for i = 0000000000000000; i < 130000; i++ {
		str := strconv.FormatInt(latest+i, 31)
		encoded := base64.StdEncoding.EncodeToString([]byte(str))
		KeySet = append(KeySet, encoded)
		fmt.Println(encoded)
	}
	fmt.Println("worker processed")
	finished <- KeySet
}
