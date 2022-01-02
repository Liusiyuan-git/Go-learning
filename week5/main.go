package main

import (
	"time"
	"week5/Sliding"
)

func main() {
	r := Sliding.NewSliding(Sliding.WithBucketDuration(2 * time.Millisecond))
	for i := 0; i < 4; i++ {
		go func() {
			for {
				time.Sleep(1 * time.Millisecond)
				r.Add(1)
			}
		}()
	}

	go func() {
		for {
			time.Sleep(1 * time.Millisecond)
			r.Sum()
		}
	}()

	time.Sleep(100 * time.Millisecond)
}
