package main

import (
	"fmt"
	"time"
)
func m1 () {
	var buf0 []int
		for {
		buf0 = append(buf0, 1, 2, 3, 4, 5)
		//time.After(10 * time.Second)
		time.Sleep(5*time.Second)
		fmt.Println(buf0)
	}
}
 func main ()  {
	go m1()
	 var buf int
	fmt.Scanln(&buf)
}

