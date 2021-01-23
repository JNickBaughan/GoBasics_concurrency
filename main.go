package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup

	setUpWorkGroupAndCount("hello", 50, &waitGroup)
	setUpWorkGroupAndCount("world", 50, &waitGroup)

	waitGroup.Wait()

	// two go routines won't do anything. The main go routine will finish before they can be run
	// go count("hello", 50)
	// go count("world",50)
	// sleeping the program can fix it
	//time.Sleep(time.Second * 5)
	// so can this - it will wait for user input to stop the program
	//fmt.Scanln()
}

type addCountToWorkGroupFunc func(waitGroup *sync.WaitGroup)

type countFunc func(printMe string, times int, waitGroup *sync.WaitGroup)

func setUpWorkGroupAndCount(printMe string, times int, waitGroup *sync.WaitGroup){
	waitGroup.Add(1)
	go count(printMe, times , waitGroup )
}

func addCountToWorkGroup(waitGroup *sync.WaitGroup){
	waitGroup.Add(1)
}

func count(printMe string, times int, waitGroup *sync.WaitGroup){
	for i := 1; i <= times; i++ {
		fmt.Println("'" + printMe + "' has been printed " + strconv.FormatInt(int64(i), 10) + " times" )
	}
	waitGroup.Done()
}