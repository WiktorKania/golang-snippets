package snippets

import "fmt"

func changeMe(x *int) {
	(*x) += 3
}

func copyMe(x int) {
	x += 3
}

func init() {
	num := 5
	copyMe(num)
	fmt.Println(num)
	changeMe(&num)
	fmt.Println(num)
}
