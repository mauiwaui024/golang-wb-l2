package main

import (
	"fmt"
	"reflect"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		cases := make([]reflect.SelectCase, len(channels))
		for i, ch := range channels {
			cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
		}
		for {
			chosen, value, ok := reflect.Select(cases)
			if !ok {
				// Channel is closed, remove it from cases
				cases[chosen].Chan = reflect.ValueOf(nil)
				cases[chosen].Dir = reflect.SelectDefault
				continue
			}
			// Send value to result channel
			c <- value.Interface()
		}
	}()
	return c
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("Done after %v", time.Since(start))
}
