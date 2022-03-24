package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	for i := 0; i < 3; i++ {
		go func(a, b int) {
			c <- a + b

			fmt.Println("released ", a+b)
		}(i, i+10)
	}

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println(<-c)
	}

}
