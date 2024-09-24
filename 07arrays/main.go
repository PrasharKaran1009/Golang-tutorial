package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to arrays, enter your age = ")
	input, _ := reader.ReadString('\n')
	n, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your age is", n)
	}

	// Initialize fruitList correctly
	var fruitList [5]string
	fruitList[0] = "Apple"
	fruitList[1] = "Maple"
	fruitList[2] = "Banana"
	fruitList[3] = "Grapes" // You can add more fruits if needed
	fruitList[4] = "Peach"

	fmt.Println(fruitList)
}
