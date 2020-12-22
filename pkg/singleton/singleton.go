package main

import (
	"fmt"
	"github.com/SandeepPissay/playground/pkg/singleton/pattern1"
	"github.com/SandeepPissay/playground/pkg/singleton/pattern2"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Test Pattern 1 - Hiding the singleton behind an interface
	fmt.Println("Testing pattern 1...")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
			foo := pattern1.GetFoo()
			foo.SetValue(strconv.Itoa(i), strconv.Itoa(i))
		}(&wg, i)
	}
	wg.Wait()
	pattern1.GetFoo().Print()

	// Test Pattern 2 - Expose only the public methods that work on a package local singleton
	fmt.Println("Testing pattern 2...")
	for j := 10; j < 20; j++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, j int) {
			defer wg.Done()
			pattern2.SetValue(strconv.Itoa(j), strconv.Itoa(j))
		}(&wg, j)
	}
	wg.Wait()
	pattern2.Print()
}