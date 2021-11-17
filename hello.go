package main

import (
	"fmt"

	"github.com/longfei-zhang/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
