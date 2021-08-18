package snippets

import (
	"fmt"
	"sync"
	"time"
)

func Away() {
	fmt.Println("Goodbye :(")
}
func aha() {
	go Away()
}

func BetterLoveStoryThanTwilight() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("I'm slow")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("I waited!")
	fmt.Println("Learn about context bitch")
}

func ToSleep() {
	fmt.Println("Going to sleep")
	time.Sleep(1 * time.Second)
	fmt.Println("Can't sleep :(")
}
