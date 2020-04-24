package example_8

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func TestExample_8_7(wg *sync.WaitGroup) {
	//learnTick()

	//echoNum()

	test(wg)
}

func learnTick() {

	fmt.Println("Commencing countdown")

	about := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		about <- struct{}{}
	}()
	/*
		tick :=time.Tick(time.Second)

		for countdown:=10;countdown>0;countdown-- {
			fmt.Println(countdown)
			j := <-tick
			fmt.Printf("j is %T %s\n", j, j.Format("2006-01-02 15:04:05"))
		}*/

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("do something...")
	case <-about:
		fmt.Println("Launch aborted!")
		return
	}

}

func echoNum() {

	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Printf("echo num is %d\n", x)
		case ch <- i:
			fmt.Println("send i to ch")
		}

	}
}

func test(wg *sync.WaitGroup) {
	var x, y int
	go func() {
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()
	go func() {
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()
}
