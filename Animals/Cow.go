package animals

import (
	"fmt"
)

func Mew() {
	fmt.Println("Moooo!!!")
}

func Bark() {
	Mew()
	fmt.Println("Bark bark!!!")
}
