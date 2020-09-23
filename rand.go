package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func GetRandomString2(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}


func main(){
	string2 := GetRandomString2(5)

	//var _ = string2
	fmt.Println(string2)
}


