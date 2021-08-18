package snippets

import (
	"fmt"
	"time"
)

func Interesting() {
	a := []string{"never", "gonna", "give", "you", "up"}
	for _, v := range a {
		fmt.Println(v)
	}
}

func YouSpinMeRound() {
	ever := true
	for ever {
		fmt.Println("Like a record, baby")
		ever = false
	}
}

func YouSpinMeRoundRound() {
	fmt.Print("desert you")
	for {
		fmt.Print("u")
		time.Sleep(1 * time.Second)
		// im making a note here
		// huge success
		// it's hard to overstate my

	}
}
