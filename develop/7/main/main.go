package main

import (
	"fmt"
	"sync"
	"time"
)

func or(cs ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	out := make(chan interface{})          //обьединенный канал
	output := func(c <-chan interface{}) { //копируем в общий канал
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(1) //количество элементов в waitgroup (один)
	for _, c := range cs {
		go output(c)
	}

	go func() { //отрабатывает после wg.Add, закрываем канал после завершения всех
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
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
		sig(2*time.Second),
		sig(1*time.Second),
		sig(3*time.Second),
	)

	fmt.Printf("done after %v\n", time.Since(start))

}
