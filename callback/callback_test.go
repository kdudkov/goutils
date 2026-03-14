package callback

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"math/rand/v2"
)

func TestRemove(t *testing.T) {
	cb := New[string]()

	for range 30 {
		cb.Subscribe(func(msg string) bool {
			if rand.IntN(1000) == 1 {
				fmt.Printf("remove\n")
				return false
			}

			return true
		})
	}

	n := 10

	ctx, cancel := context.WithCancel(context.Background())

	wg := new(sync.WaitGroup)

	for range n {

		wg.Go(func() {
			for ctx.Err() == nil {
				cb.AddMessage("aaa")

				time.Sleep(time.Millisecond * time.Duration(rand.IntN(100)))
			}

		})
	}

	time.Sleep(time.Second * 5)
	cancel()

	wg.Wait()

}
