package main

import "fmt"

//test channel
var area, cube chan int

//buffer is the buffer channel that allow buffer
var buffer chan int

//roc,soc is the single direction channel
var roc <-chan int //receive-only channel
var soc chan<- int //send-only channel

func main() {
	//First of all, channel needs to make before use
	area, cube = make(chan int), make(chan int)
	//Buffer make
	buffer = make(chan int, 3)
	//Single direction
	roc = make(<-chan int)
	soc = make(chan<- int)

	//normal goroutine
	go areaF(area)

	//Anonymous goroutine
	go func() {
		cube <- 1000
	}()

	//Area channel send in the infomation
	area <- 10
	//Wait for the area goroutine to caculate the info
	fmt.Println("The area of go routine is ", <-area)
	//Wait for the Anonymous goroutine to send out the info
	fmt.Println("The cube of go routine is ", <-cube)

	//Testing the buffer channel
	buffer <- 1
	buffer <- 2
	buffer <- 3

	//If uncomment this line will throw the panic that says the all channel asleep.
	// buffer <- 4

	fmt.Println(<-buffer)
	fmt.Println(<-buffer)
	fmt.Println(<-buffer)

	//Range testing
	go func() {
		buffer <- 1
		buffer <- 2
		buffer <- 3
		// close(buffer)
	}()

	for v := range buffer {
		fmt.Println(v)
	}

}

//areaF will send out the number
func areaF(a chan int) {
	num := <-a
	a <- num * num
}
