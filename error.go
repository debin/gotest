package main

import (
	"errors"
	"fmt"
	"golang.org/x/xerrors"
)

func main() {
	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err3: [%w]", err2)

	fmt.Println(err3)
	fmt.Printf("%+v\n", err3)
	fmt.Println(errors.Unwrap(err3))
	fmt.Println(errors.Unwrap(errors.Unwrap(err3)))

	err111 := xerrors.New("original_error")
	err222 := xerrors.Errorf("err222:%w",err111)
	fmt.Println(err222)
	fmt.Printf("%+v",err222)

}
