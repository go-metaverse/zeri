package main

import (
	"fmt"
	"time"

	"github.com/go-metaverse/zeri/routine"
)

func main() {
	routine.Run(func(msg string) {
		fmt.Println(msg)
	}, "Hello, world!")

	time.Sleep(time.Second) // wait for the goroutine to finish
}
