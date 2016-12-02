package main

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"
)

import "os"

var wg sync.WaitGroup

func main() {
	basewd, _ := os.Getwd()
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go changeDirInGoRoutine(basewd, i)
	}
	wg.Wait()
}

func changeDirInGoRoutine(baseDir string, routineNum int) {
	defer wg.Done()
	wd, _ := os.Getwd()
	fmt.Printf("routine %v: starting in dir %s\n", routineNum, wd)
	os.Chdir(filepath.Join(baseDir, fmt.Sprintf("%s%v", "r", routineNum)))
	wd, _ = os.Getwd()
	fmt.Printf("routine %v, going to sleep in dir: %s\n", routineNum, wd)
	time.Sleep(1000 * time.Millisecond)
	wd, _ = os.Getwd()
	fmt.Printf("routine %v, woke up in dir %s\n", routineNum, wd)
}

// routine 2: starting in dir C:\golang\src\github.com\dpastoor\playground
// routine 3: starting in dir C:\golang\src\github.com\dpastoor\playground
// routine 1: starting in dir C:\golang\src\github.com\dpastoor\playground
// routine 2, going to sleep in dir: C:\golang\src\github.com\dpastoor\playground\r2
// routine 3, going to sleep in dir: C:\golang\src\github.com\dpastoor\playground\r3
// routine 1, going to sleep in dir: C:\golang\src\github.com\dpastoor\playground\r1
// routine 3, woke up in dir C:\golang\src\github.com\dpastoor\playground\r1
// routine 2, woke up in dir C:\golang\src\github.com\dpastoor\playground\r1
// routine 1, woke up in dir C:\golang\src\github.com\dpastoor\playground\r1
