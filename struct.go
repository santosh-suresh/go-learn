package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
  fmt.Println(person{"Box",20})
  fmt.Println(person{age: 40, name:"Santosh"})
  fmt.Println(person{name:"Hello"})

  fmt.Println(&person{"Bob",30})

  s := person{name:"Siddarth", age:10}
  fmt.Println(s.age)

  sp := &s
  fmt.Println(sp.age)

  sp.age = 51
  fmt.Println(s.age)
}
