package main

import (
	"fmt"
)

func main() {
	i := make(chan int)
	go func() {
		for x := 0; x < 3; x++ {
			i <- x
		}
		close(i)
	}()

	for data := range i {
		fmt.Println(data)
	}
}
