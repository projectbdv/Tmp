package main

import (
	"fmt"
	"time"
)
func main () {
	var buf0 []int
		for {
		buf0 = append(buf0, 1, 2, 3, 4, 5)
		//time.After(10 * time.Second)
		time.Sleep(5*time.Second)
		fmt.Println(buf0)
	}
}
/* func main ()  {

	go m( )
	var input  string
	fmt.Scanln(&input)
}
*/
