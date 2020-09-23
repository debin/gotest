package main

import "fmt"

func main() {


	if true {
		defer fmt.Printf("1")
	} else {
		defer fmt.Printf("2")
	}
	fmt.Printf("3")


	var ttt map[string]string
	for s  := range ttt {
		fmt.Println(s )
	}



	type aaa struct {
		a string
		b int
	}

	//var tt1 = aaa{"111",123456}
	var tt1 = aaa{a: "111",b: 123456}
	var tt2 = new(aaa)
	var tt3 [10]int
	tt3[0] = 4
	//tt.a = "fwef"
	//tt.b = 1234

	//fmt.Printf("%s,%T",tt,tt)
	fmt.Println(tt1,tt2)
	fmt.Println(tt3)

	var touch = []int{1,2,3}
	touch,touch[0] = nil,3
	fmt.Printf("touch:%##v\n",touch)


	var tt aaa
	fmt.Printf("tt:%+v\n",tt)


}
