package main

import (
	"context"
	"fanIn/internal/tool"
	"fmt"
)

func main() {
	ctx := context.Background()
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		for i := range 5 {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := range 5 {
			ch2 <- i
		}
		close(ch2)
	}()

	ch := tool.FanIn(ctx, ch1, ch2)

	for {
		select {
			case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(v)
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}