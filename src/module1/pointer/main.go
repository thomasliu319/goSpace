package main

import "fmt"

func main() {

	para := ParameterStruct{Name: "aaa"}
	fmt.Println(para)

	changeParameter(&para, "bbb")
	fmt.Println(para)

	cannotChangeParameter(para, "ccc")
	fmt.Println(para)
}

type ParameterStruct struct {
	Name string
}

func changeParameter(para *ParameterStruct, value string) {
	para.Name = value
}

func cannotChangeParameter(para ParameterStruct, value string) {
	para.Name = value
}
