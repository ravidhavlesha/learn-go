package main

import (
	"fmt"
	"sync"
)

var (
	x int = 0
	y int = 0
	z int = 0
)

// Putting it in another var group to show SOC
var (
	balance int = 1000
)

func increment(wg *sync.WaitGroup) {
	x++
	wg.Done()
}

func incrementWithMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	// Lock will lock this critical section to be executed by only one go routine
	m.Lock()
	y++
	m.Unlock()
	wg.Done()
}

func incrementWithChannel(wg *sync.WaitGroup, ch chan bool) {
	// Since the buffered channel has capacity of 1 all other goroutines are blocked until the value is read from the channel
	ch <- true
	z++
	<-ch
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	var m sync.Mutex

	ch := make(chan bool, 2)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(&wg)

		wg.Add(1)
		go incrementWithMutex(&wg, &m)

		wg.Add(1)
		go incrementWithChannel(&wg, ch)
	}

	wg.Wait()

	fmt.Println("Final value of x", x)
	fmt.Println("Final value of y", y)
	fmt.Println("Final value of z", z)

	var wg1 sync.WaitGroup
	wg1.Add(1)
	credit := 500
	go func() {
		m.Lock()
		fmt.Printf("Depositing %d to account with balance: %d\n", credit, balance)
		balance += credit
		m.Unlock()
		wg1.Done()
	}()

	wg1.Add(1)
	debit := 700
	go func() {
		m.Lock()
		fmt.Printf("Withdrawing %d to account with balance: %d\n", debit, balance)
		balance -= debit
		m.Unlock()
		wg1.Done()
	}()

	wg1.Wait()

	fmt.Printf("Final balance in account: %d\n", balance)
}
