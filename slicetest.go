package main

import (
	"github.com/jianfengye/collection"
)

func main()  {
	a := []int{1,2}
	intColl := collection.NewIntCollection(a)
	intColl.Push(1)

	//intColl2 := intColl.NewEmpty()
	intColl.DD()
}
