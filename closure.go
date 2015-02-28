package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	incr := intSeq()

	fmt.Println(incr())
	fmt.Println(incr())
	fmt.Println(incr())
	fmt.Println(incr())

	inc1 := intSeq()
	fmt.Println(inc1())
}
