package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
	"os"
    "os/signal"
    "syscall"
	"context"
)

func main() {
	c := make(chan string, 20)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		fmt.Println("Program killed !")
		cancel()
	}()
	for i := 0; i < 20; i++{
		c <- "http://localhost:8080"
	}
	close(c)
	res := crawlWeb(c, ctx)
	for i := range res{
		fmt.Println(*i)
	}
}

func crawlWeb(array <-chan string, ctx context.Context) <-chan *string{
	c := make(chan *string, len(array))
	guard := make(chan struct{}, 8)
	group := sync.WaitGroup{}
	for range array {
		group.Add(1)
		go func() {
			defer group.Done()
			select {
			case <-ctx.Done():
				// break
				return
			default:
			guard <- struct{}{}
				resp, err := http.Get("http://localhost:8080")
				if err != nil {
					log.Fatalln(err)
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}
				res := string(body)
				c <- &res
				time.Sleep(1 * time.Second)
				fmt.Println("CheckGuard")
				<-guard
				// group.Done()
			}
		}()
	}
	fmt.Println(runtime.NumGoroutine())
	group.Wait()
	close(c)
	return c
}
