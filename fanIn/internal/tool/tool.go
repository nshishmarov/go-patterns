package tool

import (
	"context"
	"sync"
)

func FanIn(ctx context.Context, chans ...chan int) chan int {
	out := make(chan int, len(chans))

	go func() {
		wg := &sync.WaitGroup{}
		for _, c := range chans {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case v, ok := <-c:
						if !ok {
							return
						}
						select {
						case out <- v:
						case <-ctx.Done():
							return
						}
					case <-ctx.Done():
						return
					}
				}
			}()
		}

		wg.Wait()
		close(out)
	}()

	return out
}
