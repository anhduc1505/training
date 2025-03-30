package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func countFirstFile(result chan int, txtLocation string, keyword string) {
	numOfOcc := 0

	fileContent, err := ioutil.ReadFile(txtLocation)
	if err != nil {
		fmt.Println(err)
		result <- 0
		return
	}

	numOfOcc = strings.Count(string(fileContent), keyword)

	result <- numOfOcc

	defer close(result)
}

func countSecondFile(result chan int, txtLocation string, keyword string) {
	numOfOcc := 0

	fileContent, err := ioutil.ReadFile(txtLocation)

	if err != nil {
		fmt.Println(err)
		result <- 0
		return
	}

	numOfOcc = strings.Count(string(fileContent), keyword)

	result <- numOfOcc
}

func main() {
	cntFirstChan := make(chan int)
	cntSecondChan := make(chan int)

	go countFirstFile(cntFirstChan, "./countWordFile/1.txt", "cần")
	go countSecondFile(cntSecondChan, "./countWordFile/2.txt", "sao")

	fmt.Println("Number of occurrences in first file: ", <-cntFirstChan)
	fmt.Println("Number of occurrences in second file: ", <-cntSecondChan)
}
