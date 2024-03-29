package main

import "fmt"

type Country struct {
	Name string
}

type City struct {
	Name string
}

type Printable interface {
	PrintStr()
}

func (c Country) PrintStr() {
	fmt.Println(c.Name)
}

func (c City) PrintStr() {
	fmt.Println(c.Name)
}

func main() {
	c1 := Country{"China"}
	c2 := City{"Beijing"}
	c1.PrintStr()
	c2.PrintStr()
}
