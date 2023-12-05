package main

import (
	"fmt"
	"reflect"
)

func main() {
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	println(foo[3])
	println(foo[4])
	bar := foo[1:4]
	bar[1] = 99
	println(bar[1])

	//v1 := data{}
	//v2 := data{}
	//fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2))
	//prints: v1 == v2: true
	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2))
	//prints: m1 == m2: true
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2))
	//prints: s1 == s2: true

}
