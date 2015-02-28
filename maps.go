package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k3"] = 8
	m["k2"] = 9

	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println(v1)
	fmt.Println("Len", len(m))

	delete(m, "k2")
	fmt.Println(m)

	_, prs := m["k2"]
	fmt.Println(prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

}
