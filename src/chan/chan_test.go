package _chan

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

//超时问题
func doBadthing(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func doGoodthing(done chan bool) {
	time.Sleep(time.Second)
	select {
	case done <- true:
	default:
		return
	}
}
func timeout(f func(chan bool)) error {
	done := make(chan bool)
	go f(done)
	select {
	case <-done:
		print("done")
		return nil
	case <-time.After(time.Millisecond):
		print("1")
		return fmt.Errorf("timeout")

	}
}

func timeoutWithBuffer(f func(chan bool)) error {
	done := make(chan bool, 1)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")

	}
}
func test(t *testing.T, f func(chan bool)) {
	t.Helper()
	for i := 0; i < 100; i++ {
		timeout(f)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}
func TestBadTimeout(t *testing.T) {
	test(t, doBadthing)
}

func TestBufferTimeout(t *testing.T) {
	for i := 0; i < 100; i++ {
		timeoutWithBuffer(doBadthing)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func TestGoodTimeout(t *testing.T) {
	test(t, doGoodthing)
}

func test1() {
}

func testnum() {
	i := 0
	i++
	print(i)
}
