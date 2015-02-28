package main

import "fmt"

func main() {

  nums := []int{1,2,3}

  sum := 0
  for _,num := range nums {
    sum += num
  }

  fmt.Println("Sum is ",sum)

  for index,num := range nums {
    if num == 3 {
      fmt.Println("Index: ",index)
    }
  }

  kvs := map[string]string{"a": "apple", "b": "banana"}
  for key,val := range kvs {
    fmt.Printf("%s->%s\n",key,val)
  }

  for i, c := range "go" {
        fmt.Println(i, c)
  }

}
