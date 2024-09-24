package main

import "fmt"

const Hlo float32 = 4.567 
/*
if name of variable or constant starts from capital letter than it means it 
is available publicly'
*/

func main() {
	fmt.Println("variables")
	var username string = "Karan"
	fmt.Println(username)
	fmt.Printf("%s and it is of the type %T \n", username, username)
	var idHere bool = true
	fmt.Println(idHere)
	var small_value uint8 = 255
	fmt.Println(small_value)
	var smallfloat float64 = 1.25428874568564864
	fmt.Println(smallfloat)
	var hi int
	fmt.Println(hi)

	//implicit way
	var website = "hlo web user"
	fmt.Printf("%T \n", website)
	//coolest decalaration => we cant use walrus operator in global scope
	user := 8
	fmt.Println(user)
	fmt.Printf("%T,%f", Hlo, Hlo)
}
