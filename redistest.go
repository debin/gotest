package main

import (
	"fmt"
	"gotest/util"
)

func main() {

	//client := lib.RedisClient.Get()
	//client.Do("SET", "test777", "001", "EX", "500")
	//
	//reply, err := client.Do("GET", "test777")
	//s, _ := redis.String(reply,err)

	s1, _ := util.RedisUtil.Set("test777", "998", 1000)
	s, _ := util.RedisUtil.Get("test777")
	fmt.Println(s1, s)

}
