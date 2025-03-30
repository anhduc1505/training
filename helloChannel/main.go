package main

import (
	"log"
	"sync"
)

func printNumber(wg *sync.WaitGroup, numChan chan int) {
	var result int
	for i := 0; i < 10; i++ {
		result += i
	}
	numChan <- result
	wg.Done()
}

func printChar(wg *sync.WaitGroup, strChan chan string) {
	var result string
	for i := 'A'; i < 'A'+26; i++ {
		result += string(i)
	}
	strChan <- result
	wg.Done()
}

func main() {

	chanPrintNumber := make(chan int)
	chanPrintChar := make(chan string)

	var wg sync.WaitGroup

	wg.Add(2)

	go printNumber(&wg, chanPrintNumber)
	go printChar(&wg, chanPrintChar)

	log.Println("result printNum: ", <-chanPrintNumber)
	log.Println("result printChar: ", <-chanPrintChar)

	wg.Wait()

	log.Println("done")

}
