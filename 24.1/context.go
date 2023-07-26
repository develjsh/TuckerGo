package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)


var wg sync.WaitGroup
func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go PrintEverySeconds(ctx)
	time.Sleep(5*time.Second)
	cancel()

	wg.Wait()
}

func PrintEverySeconds(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <- ctx.Done():
			wg.Done()
			return
		case <- tick:
			fmt.Println("tick")
		}
	}
}