package main

import (
	"flag"
	"fmt"
)


func main(){

		// var name string //[1]
		// flag.StringVar(&name, "name", "everyone", "The greeting oject.") //[2]
		var name = flag.String("name", "everyone", "The greeting object.")
		flag.Parse()

		fmt.Printf("Hello , %v!\n", *name)
}