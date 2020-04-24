package example_8

import (
	"fmt"
)

func TestExample_8_4() {

	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}

func TestExample_8_4_2()  {
	naturals := make(chan int)
	squares := make(chan int)
	
	go counter(naturals)
	go squarer(squares, naturals)
	go printer(squares)
	printer(squares)
}

func counter(out chan<- int)  {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int)  {
	for x:= range in{
		out<- x*x
	}
	close(out)
}

func printer(in <-chan int)  {
	for x:= range in{
		fmt.Println(x)
	}
}


func TestExample_8_4_4(){
	fmt.Println(mirroredQuery())
	fmt.Printf("type is %T\n", mirroredQuery())
}


func mirroredQuery() string {
	responses := make(chan string)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses // return the quickest response
}

func request(hostname string) (response string) {
	return hostname
}