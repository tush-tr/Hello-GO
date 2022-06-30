package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev/"

func main(){
	a := "Hello"
	fmt.Println(a)

	response, err := http.Get(url)
	databytes, err := ioutil.ReadAll(response.Body)
	if(err!=nil){
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)

	// let's check type of response here
	fmt.Printf("type of response: %T",response)
	defer response.Body.Close(); //callers responsibility to close the connection

	


}