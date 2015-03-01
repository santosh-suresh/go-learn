package main

import (
    "fmt"
    "math"
)


type Shape interface {
   area() float64
}

type Circle struct {
  x,y,r float64
}

type Rectangle struct {
  w,h float64
}

type MultiShape struct {
  shapes []Shape
}

func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
  return r.w * r.h
}

func (m *MultiShape) area() float64 {
  var area float64
  for _,s := range m.shapes {
    area += s.area()
  }
  return area
}

func main() {
  c := Circle{x:0,y:0,r:25}
  r := Rectangle{w:30,h:10}
  c1 := Circle{x:0,y:0,r:2}
  c2 := Circle{x:0,y:0,r:1.6}
  c3 := Circle{x:0,y:0,r:8}
  m1 := MultiShape{shapes: []Shape{&c1,&c2,&c3}}
  s := []Shape{&c,&r, &m1}
  for _,ss := range s {
    fmt.Println(ss.area())
  }
}
