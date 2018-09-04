package main

import (
	"fmt"
	"time"
)

func main() {
	var buf0 []int
	for {
		buf0 = append(buf0, 1, 2, 3, 4, 5)
		time.After(10*time.Second)
		fmt.Println(buf0)
	}
}

//func aa ()  {
//	for  {
//
//	}
//}
