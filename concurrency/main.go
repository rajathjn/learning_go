// In go, Don't communicate by sharing memory; share memory by communicating.
// go keyword for concurrency, channels for communication, waitgroups for synchornization and select for co-ordination
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func worker[T any](i T) {
	fmt.Printf("Worker %v of type %T has start to work\n", i, i)
	sleep_time := rand.Float64() * 2000
	time.Sleep(time.Duration(int64(sleep_time)) * time.Millisecond)
	fmt.Printf("Worker %v of type %T has completed work\n", i, i)
	wg.Done()
}

func toss(ch chan string) {
	defer close(ch)
	options := [2]string{"Heads", "Tails"}
	for {
		choice := options[rand.IntN(len(options))]
		ch <- choice
		if choice == "Heads" {
			break
		}
	}
}

func roll(ch chan int) {
	defer close(ch)
	dice := [6]uint8{1,2,3,4,5,6}
	for {
		choice := dice[rand.IntN(len(dice))]
		ch <- int(choice)
		if choice == 6 {
			break
		}
	}
}

func main() {
	total_workers := 5
	for i := range total_workers {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()

	// For getting data
	t_ch := make(chan string)
	r_ch := make(chan int)
	go toss(t_ch)
	go roll(r_ch)

	for {
		select {
			case val, ok := <-t_ch:
				if !ok {
					t_ch = nil
					continue
				}
				fmt.Println("Toss result:", val)
			case val, ok := <-r_ch:
				if !ok {
					r_ch = nil
					continue
				}
				fmt.Println("Dice roll:", val)
			case <-time.After(time.Duration(5) * time.Millisecond):
				fmt.Println("Exiting due to time expiry")
				return
		}
		
		if t_ch == nil && r_ch == nil {
			break
		}
	}
}