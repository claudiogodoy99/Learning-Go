package main

import (
	"fmt"
	"time"
)

type X struct {
	a int
}

func (r *X) modify(n int) {
	r.a = n
}

func main() {

	ch := make(chan X, 5)

	// go forLoopReciever(ch)
	go forLoopSender(ch)

	for {
		select {
		case v := <-ch:
			{
				time.Sleep(5 * time.Second)
				fmt.Println("MAIN GO ROUTINE SAYS: %d", v)
			}
		default:
			time.Sleep(5 * time.Second)
			fmt.Println("nada recebido")
		}

	}
}
func forLoopSender(ch chan X) {
	n := 0
	for {
		n++
		ch <- X{
			a: n,
		}

		if n%10 == 0 {
			for i := 0; i < 5; i++ {
				fmt.Println("Consuming 5: ", (<-ch).a)
			}
		}

		fmt.Println("sent: ", time.Now())
	}
}

func forLoopReciever(ch chan X) {

	for value := range ch {
		time.Sleep(5 * time.Second)
		fmt.Println("Some routine says %d", value.a)
	}
}
