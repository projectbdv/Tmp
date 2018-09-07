package main

import (
	"os"
	"log"
	//"fmt"
	"io/ioutil"
	"fmt"
)


func main() {
	// Open file for reading
	file, err := os.Open("mas.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	}