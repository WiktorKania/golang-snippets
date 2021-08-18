package snippets

import "fmt"

type Person struct {
	Name string
}

func gimmeDat() *Person {
	localHomie := Person{"Adrian"}
	return &localHomie
}

func init() {
	thisShouldCrash := gimmeDat()
	fmt.Println(thisShouldCrash)
}
