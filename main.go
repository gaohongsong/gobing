package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2*time.Second)
		fmt.Println("1号完成")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2*time.Second)
		fmt.Println("2号完成")
	}()

	wg.Wait()

	fmt.Println("好了，大家都干完了，放工")
}
