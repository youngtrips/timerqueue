package main

import (
	"fmt"

	"github.com/youngtrips/timerqueue"
)

func main() {
	q := timerqueue.New(20)
	q.Run()

	ch := make(chan int64)
	t1 := timerqueue.NowMS()
	q.Schedule(5000, ch)
	id := <-ch
	t2 := timerqueue.NowMS()
	fmt.Printf("expired: id=%d, start=%d, expired=%d, diff=%d\n", id, t1, t2, t2-t1)
}
