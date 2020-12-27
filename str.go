package main

import "fmt"

// #include <stdio.h>
import "C"







func main() {
	C.puts(C.CString("Hello, World\n"))


	a := []byte("你好")
	fmt.Printf("%s",a)


	b := (string)(a)
	fmt.Printf("\n%s",b)


	c := ([]byte)("你好")
	fmt.Printf("\n%s",c)



}
