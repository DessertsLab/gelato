package main

import (
	"sync"
	"testing"
)

func modi(x *string) {
	*x = "what ahi"
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	*x, *y = *y, *x

}

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	//	s := new(string)
	//	modi(s)
	//	fmt.Println(*s)
	//	x := 1.5
	//	square(&x)
	//	fmt.Print(x)
	//	x := 1
	//	y := 2
	//	swap(&x, &y)
	//	fmt.Print(x, y)
	//	var memoryAccess sync.Mutex
	//	var data int
	//	go func() {
	//		memoryAccess.Lock()
	//		data++
	//		memoryAccess.Unlock()
	//	}()
	//	//time.Sleep(1 * time.Second)
	//	memoryAccess.Lock()
	//	if data == 0 {
	//		fmt.Println("value is 0")
	//	} else {
	//		fmt.Printf("value is %v. \n", data)
	//	}
	//	memoryAccess.Unlock()

	//  死锁例子
	//	var wg sync.WaitGroup
	//	printSum := func(v1, v2 *value) {
	//		defer wg.Done()
	//		v1.mu.Lock()         // Here we attempt to enter the critical section for the incoming value
	//		defer v1.mu.Unlock() // Here we use the defer statement to exit the critical section before printSum returns.
	//
	//		time.Sleep(2 * time.Second) // Here we sleep for a period of time to simulate work (and trigger a deadlock).
	//		v2.mu.Lock()
	//		defer v2.mu.Unlock()
	//
	//		fmt.Println("sum=%v\n", v1.value+v2.value)
	//
	//	}
	//
	//	var a, b value
	//	wg.Add(2)
	//	go printSum(&a, &b)
	//	go printSum(&b, &a)
	//	wg.Wait()

	// 活锁例子
	//	cadence := sync.NewCond(&sync.Mutex{})
	//	go func() {
	//		for range time.Tick(1 * time.Millisecond) {
	//			cadence.Broadcast()
	//		}
	//	}()
	//
	//	takeStep := func() {
	//		cadence.L.Lock()
	//		cadence.Wait()
	//		cadence.L.Unlock()
	//	}
	//
	//	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
	//		fmt.Fprintf(out, " %v", dirName)
	//		atomic.AddInt32(dir, 1)
	//		takeStep()
	//		if atomic.LoadInt32(dir) == 1 {
	//			fmt.Fprint(out, ". Success!")
	//			return true
	//		}
	//		takeStep()
	//		atomic.AddInt32(dir, -1)
	//		return false
	//	}
	//
	//	var left, right int32
	//	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
	//	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }
	//
	//	walk := func(walking *sync.WaitGroup, name string) {
	//		var out bytes.Buffer
	//		defer func() { fmt.Println(out.String()) }()
	//		defer walking.Done()
	//		fmt.Fprintf(&out, "%v is trying to scoot:", name)
	//		for i := 0; i < 5; i++ {
	//			if tryLeft(&out) || tryRight(&out) {
	//				return
	//			}
	//		}
	//		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	//	}
	//
	//	var peopleInHallway sync.WaitGroup
	//	peopleInHallway.Add(2)
	//	go walk(&peopleInHallway, "Alice")
	//	go walk(&peopleInHallway, "Barbara")
	//	peopleInHallway.Wait()

	//饥饿例子
	//	var wg sync.WaitGroup
	//	var sharedLock sync.Mutex
	//	const runtime = 1 * time.Second
	//
	//	greedyWorker := func() {
	//		defer wg.Done()
	//
	//		var count int
	//		for begin := time.Now(); time.Since(begin) <= runtime; {
	//			sharedLock.Lock()
	//			time.Sleep(3 * time.Nanosecond)
	//			sharedLock.Unlock()
	//			count++
	//		}
	//
	//		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	//	}
	//
	//	politeWorker := func() {
	//		defer wg.Done()
	//
	//		var count int
	//		for begin := time.Now(); time.Since(begin) <= runtime; {
	//			sharedLock.Lock()
	//			time.Sleep(1 * time.Nanosecond)
	//			sharedLock.Unlock()
	//
	//			sharedLock.Lock()
	//			time.Sleep(1 * time.Nanosecond)
	//			sharedLock.Unlock()
	//
	//			sharedLock.Lock()
	//			time.Sleep(1 * time.Nanosecond)
	//			sharedLock.Unlock()
	//
	//			count++
	//		}
	//
	//		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	//	}
	//
	//	wg.Add(2)
	//	go greedyWorker()
	//	go politeWorker()
	//
	//	wg.Wait()

	//join point 例子
	//var wg sync.WaitGroup
	//sayHello := func() {
	//	defer wg.Done()
	//	fmt.Println("hello")
	//}
	//wg.Add(1)
	//go sayHello()
	//wg.Wait() // join point :block the main goroutine until the goroutine hosting the sayHello function terminates
	// //time.Sleep(1 * time.Second) // not join point only a race condition

	// goroutine 占用内存空间例子
	//memConsumed := func() uint64 {
	//	runtime.GC()
	//	var s runtime.MemStats
	//	runtime.ReadMemStats(&s)
	//	return s.Sys
	//}

	//var c <-chan interface{}
	//var wg sync.WaitGroup
	//noop := func() { wg.Done(); <-c } //this goroutine won’t exit until the process is finished

	//const numGoroutines = 1e4
	//wg.Add(numGoroutines)
	//before := memConsumed()
	//for i := numGoroutines; i > 0; i-- {
	//	go noop()
	//}
	//wg.Wait()
	//after := memConsumed()
	//fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)

}

// 上下文切换消耗
func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
