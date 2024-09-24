package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time travel ")
	presentTime := time.Now()
	// fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05"))
	creatDate := time.Date(2024, time.December, 20, 17, 0, 0, 0, time.Local)
	fmt.Println(creatDate.Format("02-01-2006 Monday"))
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()

}
