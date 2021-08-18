package snippets

import (
	"fmt"
	"strings"
	"time"
)

type harimatis struct {
	JakPatrzeTakToMnieWidac string
	aleJakTakToJuzNie       string
}

func (h *harimatis) Pokaż() {
	for i := 0; i < 3; i += 1 {
		if i == 1 {
			h.JakPatrzeTakToMnieWidac = h.JakPatrzeTakToMnieWidac[:strings.LastIndex(h.JakPatrzeTakToMnieWidac, " ")]
			h.aleJakTakToJuzNie = h.aleJakTakToJuzNie[:strings.LastIndex(h.aleJakTakToJuzNie, " ")]
		}
		for _, c := range h.JakPatrzeTakToMnieWidac {
			fmt.Printf("%c", c)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println()
		for _, c := range h.aleJakTakToJuzNie {
			fmt.Printf("%c", c)
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println()
		if i == 1 {
			fmt.Println("Troche widać")
		}
	}
}

func New() *harimatis {
	fmt.Println("- Ave Cezarze, chwała i sława")
	return &harimatis{"Widać mnie", "Nie widać mnie"}
}
