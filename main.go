package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	numbers := "one two three four five six seven eight nine ten"
	chan1 := make(chan string)
	chan2 := make(chan string)
	go goroutineReformat(chan1, numbers)
	go goroutineFilter(chan1, chan2)
	goroutinePrint(chan2)
}
func goroutineReformat(downstream chan string, numbers string) {
	stringSlice := strings.Fields(numbers)
	for value := range stringSlice {
		downstream <- stringSlice[value]
	}
	close(downstream)
}

func goroutineFilter(upstream, downstream chan string) {
	for value := range upstream {
		if value != "" {
			downstream <- value
		}

	}
	close(downstream)
}

func goroutinePrint(upstream chan string) {
	for value := range upstream {
		fmt.Println(value)
		time.Sleep(1 * time.Second)
	}
}
