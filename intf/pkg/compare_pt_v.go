package intf

import "fmt"

type counter interface {
	add()
	current() int
}

type vCounter struct {
	no int
}

func (v vCounter) add() {
	v.no++
}

func (v vCounter) current() int {
	return v.no
}

type pCounter struct {
	no int
}

func (p *pCounter) add() {
	p.no++
}

func (p *pCounter) current() int {
	return p.no
}

// TwoKindsOfImpls tests the two kinds of implementation of interface
func TwoKindsOfImpls() {
	var c1 counter = vCounter{10}
	fmt.Println("Testing value implementation...")
	fmt.Printf("Initialized value in counter is %d\n", c1.current())

	c1.add()
	// c1.add() is a value implementation, so c1.add() would not change c1.no
	fmt.Printf("Value after adding %d\n", c1.current())

	var c2 counter = &pCounter{100}
	fmt.Println("Testing pointer implementation...")
	fmt.Printf("Initialized value in counter is %d\n", c2.current())

	c2.add()
	// c2.add() is a pointer implementation, so c2.add() does change c2.no
	fmt.Printf("Value after adding %d\n", c2.current())

	var c3 counter = &vCounter{1000}
	fmt.Println("Testing value implementation with pointer...")
	fmt.Printf("Initialized value in counter is %d\n", c3.current())

	c3.add()
	// c3.add() is a value implementation, so c3.add() would not change c3.no
	fmt.Printf("Value after adding %d\n", c3.current())
}
