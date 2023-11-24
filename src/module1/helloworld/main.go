package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)

	fullString := fmt.Sprintf("Hello %s from Go \n", *name)
	fmt.Printf(fullString)

	err, result := DuplicateString("aaa")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

func DuplicateString(input string) (error, string) {
	if input == "aaa" {
		return fmt.Errorf("aaa is not allowed"), ""
	}
	return nil, input + input
}
