package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello guys ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating from 1 to 5 = ")
	input, _ := reader.ReadString('\n')
	n, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n + 1)
	}

}
