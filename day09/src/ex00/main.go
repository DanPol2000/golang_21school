package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepWrite(i int, c chan int, group *sync.WaitGroup) {
	time.Sleep(time.Duration(i) * time.Second)
	c <- i
	group.Done()
}

func sleepSort(array []int) chan int {
	fmt.Println("main() started")
    c := make(chan int, len(array))
	
	group := sync.WaitGroup{}
	for _, i := range array{
		group.Add(1)
		go sleepWrite(i, c, &group)
	}
	group.Wait()
	close(c)
	return c

}

func main() {
	array := []int{5, 6, 1, 9, 8, 2, 3, 7}

	c := sleepSort(array)
	for i := range c {
		fmt.Println(i)
	}
}