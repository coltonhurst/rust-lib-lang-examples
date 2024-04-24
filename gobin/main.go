package main

import (
	"fmt"
)

/*
#cgo LDFLAGS: -L ../rustlib/target/debug -lrustlib
extern int add(int left, int right);
*/
import "C"

func main() {
	result := C.add(5, 3)
	fmt.Println("The sum is:", result)
}
