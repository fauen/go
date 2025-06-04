// A friendly message and current time.
package main

import (
	"fmt"
	"time"
)

// This will greet and display the current time.
func greeting() string {
	return "Hello buddy! The current time is: " + time.Now().String()
}

func main() {
	fmt.Println(greeting())
}
