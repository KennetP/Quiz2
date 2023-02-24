package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 3; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}
func printLetters() {
	for i := 'a'; i <= 'c'; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}

func main() {
	go printNumbers()
	go printLetters()

	time.Sleep(3 * time.Second)
}
