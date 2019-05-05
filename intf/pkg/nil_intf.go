package intf

import "fmt"

type myerr struct {
	msg string
}

func (e myerr) Error() string {
	return e.msg
}

func retNil() interface{} {
	return nil
}

func retNilStruct() interface{} {
	var ret *myerr
	return ret
}

// NilIntf tests all the interface related samples.
func NilIntf() {
	i1 := retNil()
	fmt.Printf("Type: %T, %v\n", i1, i1)
	fmt.Printf("Is returned value nil? %v\n", i1 == nil)

	i2 := retNilStruct()
	fmt.Printf("Type: %T, %v\n", i2, i2)
	// i2 == nil -> false, because the type of i2 is *myerr.
	// in Golang, if an interface has a certain type, i2 is not nil
	fmt.Printf("Is returned value nil? %v\n", i2 == nil)
}
