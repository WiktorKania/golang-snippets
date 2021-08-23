package snippets

import (
	"fmt"
	"strconv"
)

type DaddyDaddy interface {
	Do(number int) string
}

type Thingy struct {
	Number int
}

type Integer int

func (i Integer) Do(number int) string {
	return strconv.Itoa(int(i) + number)
}

func (t Thingy) Do(number int) string {
	return strconv.Itoa(t.Number + number)
}

func init() {
	var notAJavaInt Integer = 3
	fmt.Println(notAJavaInt.Do(5))
	fmt.Println(Thingy{3}.Do(5))
}
