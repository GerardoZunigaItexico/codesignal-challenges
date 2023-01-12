package _3_is_infinite_process

import (
	"context"
	"time"
)

func solution01(a int, b int) bool {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	r := make(chan bool, 1)

	go isInfiniteLoop(ctxTimeout, a, b, r)

	select {
	case <-ctxTimeout.Done():
		return true
	case rb := <-r:
		return rb
	}
}

func isInfiniteLoop(ctx context.Context, a, b int, r chan<- bool) {
	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -MaxInt - 1
	for a != b && (a < MaxInt || b > MinInt) {
		a++
		b--
		if a == b {
			r <- false
		}
	}
	r <- true
}

func solution02(a int, b int) bool {
	if b < a {
		return true
	}
	if isPair := (b-a)%2 == 0; isPair {
		return false
	}
	return true
}
func solution03(a int, b int) bool {
	return !(b >= a && (b-a)%2 == 0)
}
