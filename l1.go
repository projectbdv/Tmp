package main

import (
	"fmt"
	"time"
	"runtime"
	_ "net/http/pprof"
	"github.com/pkg/profile"
)
func m1 () {
	var buf0 []int
		for {
		buf0 = append(buf0, 1, 2, 3, 4, 5)
		//time.After(10 * time.Second)
		time.Sleep(2*time.Second)
		fmt.Println(buf0)
		runtime.Gosched()
	}
}
 func main ()  {
	 for i:=0;i <100 ;i++ {
		 go m1()
	 }
	 defer profile.Start().Stop()
	 var buf int
	fmt.Scanln(&buf)
 }

//test of test tett