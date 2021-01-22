package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		
		count("hello", 50)
		waitGroup.Done()
	}()

	go func() {

		count("world", 50)
		waitGroup.Done()
	}()

	waitGroup.Wait()

	// two go routines won't do anything. The main go routine will finish before they can be run
	// go count("hello", 50)
	// go count("world",50)
	// sleeping the program can fix it
	//time.Sleep(time.Second * 5)
	// so can this - it will wait for user input to stop the program
	//fmt.Scanln()
}

func count(printMe string, times int){
	for i := 1; i <= times; i++ {
		fmt.Println("'" + printMe + "' has been printed " + strconv.FormatInt(int64(i), 10) + " times" )
	}
}