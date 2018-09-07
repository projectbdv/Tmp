package main

import (
	//"fmt"
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	//file, err := os.Open("mas.txt")
	//if err != nil {
//return
//log.Fatal(err)
	//}
//defer file.Close()


//stat, err := file.Stat()
//if err != nil {
//return
//	}

	//bs := make([]byte,125)
//	_, err = file.Read(bs)
//	if err != nil {
//		return
//	}
data, err := ioutil.ReadFile("mas.txt")
if err !=nil {
	log.Fatal(err)
}
////file, err := os.Open("mas.txt")

	//if err != nil {
	//	log.Fatal(err)
	//}

//data, err := ioutil.ReadAll(file)
//if err != nil {
//	log.Fatal(err)
//}
//	fmt.Println(string(data))

fmt.Printf(string(data))
}