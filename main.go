package main

import (
	"fmt"
	"sync"
	"time"
)

import "os"

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go osExitInGoRoutine(i)
	}
	wg.Wait()
}

func osExitInGoRoutine(routineNum int) {
	defer wg.Done()
	fmt.Printf("routine %v: starting \n", routineNum)
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("routine %v: woke up and ready to do more stuff \n", routineNum)
	fmt.Printf("routine %v: about to exit", routineNum)
	os.Exit(1)
}
