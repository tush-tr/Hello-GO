package main
import "fmt"

func main(){
	fmt.Println("Removing value from a slice based on index")
	// How to remove a value from slices based on index
	var courses = []string{"reactjs","javascript","swift","elasticSearch","docker","Kubernetes"}
	index := 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)



}