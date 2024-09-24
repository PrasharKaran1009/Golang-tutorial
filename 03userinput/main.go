package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello Karan")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating = ")
	input, _ := reader.ReadString('\n')
	fmt.Println("the rating is ", input)
	fmt.Printf("the type of rating is %T", input)
}
