package main

import (
	"fmt"
	"github.com/jianfengye/collection"
)

func main()  {
	a := []int{1,2}
	intColl := collection.NewIntCollection(a)
	intColl.Push(1)

	//intColl2 := intColl.NewEmpty()
	intColl.DD()

	ints := make([]int, 3, 10)
	for i,v := range ints {
		fmt.Println(i,v)
	}
	//ints[3] = 3
	//fmt.Print(ints)
}
