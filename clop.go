package main

import (
	"fmt"
	"github.com/guonaihong/clop"
	"os"
)

type Hello struct {
	File string `clop:"-f; --file" usage:"file"`
}

func main() {

	os.Args = append(os.Args,"-f")
	os.Args = append(os.Args,"test2")

	c := clop.New(os.Args[1:])
	h := Hello{}
	//clop.Bind(&h)
	c.Bind(&h)
	fmt.Printf("%+v\n", h)
}
// ./one -f test
// main.Hello{File:"test"}
// ./one --file test
// main.Hello{File:"test"}