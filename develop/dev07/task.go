package dev07

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	orChannel := make(chan interface{})

	go func() {
		for i := 0; i < len(channels); i++ {
			// Restart loop
			if i == len(channels)-1 {
				i = 0
			}

			select {
			case <-channels[i]:
				if _, ok := <-channels[i]; !ok {
					close(orChannel)
					return
				}
			default:
				continue
			}
		}
	}()

	return orChannel
}

func OrChannel() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))

}
