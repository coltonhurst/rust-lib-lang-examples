package cinterface

import "fmt"
import "reflect"
import "strconv"

/*
#cgo LDFLAGS: -L ../../../rustlib/target/debug -lrustlib
extern int add(int left, int right);
*/
import "C"

// type Library interface {
// 	RunCommand(var1 int, var2 int) (string, error)
// }

// type LibraryImpl struct{}

// func NewLibrary() Library {
// 	return &LibraryImpl{}
// }

func RunCommand(var1 int, var2 int) (string, error) {
	var1C := C.int(var1)
	var2C := C.int(var2)
	result := int(C.add(var1C, var2C))
	fmt.Println(result)
	fmt.Printf("'result' is of type: %s\n", reflect.TypeOf(result))
	resultString := strconv.Itoa(result)
	return resultString, nil
}
