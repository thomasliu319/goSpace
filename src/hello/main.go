package main
import (
	"fmt"
)


func main(){
	fmt.Println("Hello World!")

	words :=[] string{"foo","bar","baz"}

	for _, word:= range words{
		fmt.Println(word)
	}


	
	resource :=[] string{"I","am","stupid","and","weak"}

	fullSlice :=  make([]string, 5)

	for _, word:= range resource{
		if word == "stupid" {
			word = "smart"
		} else if word == "weak"{
			word = "strong"
		}

		fullSlice = append(fullSlice, word)

		fmt.Println(word)
	}
	fmt.Println(fullSlice)
	


	// res, err := http.Get("https://www.spreadsheetdb.io/")
    // defer res.Body.close()
	// if err != nil{
	// 	log.Fatal(err)
	// }



}

