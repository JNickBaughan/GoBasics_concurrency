package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup

	c1 := make(chan string)
	c2 := make(chan string)

	setUpWorkGroupAndCount("hello", 50, c1, &waitGroup)
	setUpWorkGroupAndCount("world", 50, c2, &waitGroup)

	msg1 := <- c1;
	fmt.Println(msg1)

	msg2 := <- c2;
	fmt.Println(msg2)

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

func setUpWorkGroupAndCount(printMe string, times int, c chan string, waitGroup *sync.WaitGroup){
	waitGroup.Add(1)
	go count(printMe, times , waitGroup, c)
}

func addCountToWorkGroup(waitGroup *sync.WaitGroup){
	waitGroup.Add(1)
}

func count(printMe string, times int, waitGroup *sync.WaitGroup, c chan string){
	for i := 1; i <= times; i++ {
		fmt.Println("'" + printMe + "' has been printed " + strconv.FormatInt(int64(i), 10) + " times" )
		if(i == 2){
			c <- "this is a message from the channel"
			close(c)
		}
	}
	waitGroup.Done()
}