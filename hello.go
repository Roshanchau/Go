package main

import (
	"fmt"

	"rsc.io/quote"

	"bytes"

	"sync"
)

type Person struct {
	age int
}

type syncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func (x Person) Age() int {
	return x.age
}

func main() {
	fmt.Println(quote.Go())

	x := 7
	y := 2
	fmt.Println(x<<8 + y<<16)

	if x > 10 {
		fmt.Println("y =>", y)
	} else if x > 5 {
		fmt.Printf("x is %d \n", x)
	}

	p := Person{age: 30}
	fmt.Println("Person's age:", p.Age())

	sum := 0
	for i := range 10 {
		sum += i
	}
	fmt.Println("Sum of first 10 integers:", sum)

	numbers := []int{1, 20, 25}
	for index, value := range numbers {
		fmt.Printf("index: %d , value: %d \n", index, value)
	}

	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	fmt.Println("Main body")

	s := new(syncedBuffer)

	s.lock.Lock()
	s.buffer.WriteString("Hello, World! \n")
	s.lock.Unlock()
	fmt.Print(s.buffer.String())

	var slice *[]int = new([]int) // Allocates memory for the slice's "Table of Contents" (header) and zeros it.
	// The zeroed header means its internal pointer is nil, length is 0, capacity is 0.

	fmt.Println(slice) // Output: 0x... (a memory address, because it's a pointer)
	fmt.Println(*slice)

	var v []int = make([]int, 10, 100)

	fmt.Println(v)
	fmt.Println(len(v))
	fmt.Println(cap(v))

	v[0] = 1
	v[9] = 10
	fmt.Println(v)
}
