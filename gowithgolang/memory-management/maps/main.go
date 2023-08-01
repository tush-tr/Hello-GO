package main
import "fmt"

func main(){
	fmt.Println("Maps in Golang")
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["SOL"] = "Solidity"
	languages["PY"] = "Python"
	fmt.Println(languages)
	fmt.Println(languages["JS"])
	delete(languages, "RB")
	fmt.Println(languages)
}