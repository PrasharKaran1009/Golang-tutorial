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
	fmt.Println("Let's start pointers, enter your age = ")
	input, _ := reader.ReadString('\n')
	n, err := strconv.ParseInt(strings.TrimSpace(input), 10,64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("so your age is ", n)
	}
	//Let's start with pointer
	// var ptr *int64
	//fmt.Println("the value of pointer is ",ptr)
	var ptr = &n
	fmt.Println("the value of pointer is ",ptr)
}
