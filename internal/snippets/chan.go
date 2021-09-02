package snippets

import (
	"fmt"
	"time"
)

func createWaffle(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "Waffle created!"
}

func createPancake(ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- "Pancake created!"
}

func init() {
	waffleChannel := make(chan string)
	pancakeChannel := make(chan string)
	go createWaffle(waffleChannel)
	go createPancake(pancakeChannel)
	select {
	case result := <-waffleChannel:
		fmt.Println("A message from the waffle kingdom:", result)
	case result := <-pancakeChannel:
		fmt.Println("A message from the pancake kingdom:", result)
	}
}
