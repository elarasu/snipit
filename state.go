package main

import (
	"fmt"
	"time"
)

func me(data chan string) {
	for {
		time.Sleep(1000 * time.Millisecond)
		data <- "ME!"
	}
}
func nome(data chan string) {
	for {
		time.Sleep(1000 * time.Millisecond)
		data <- "NO ME"
	}
}

func StateMonitor() chan string {
	state := make(chan string)
	var str string
	ticker := time.Tick(1000 * time.Millisecond)
	go func() {
		for {
			select {
			case c := <-state:
				str = c
				fmt.Println(str)
			case <-ticker:
				fmt.Println("Tick")
			}
		}
	}()
	return state
}

func main() {
	state := StateMonitor()
	go me(state)
	go nome(state)
	/*
		    go func() {
		        for changed := range state {
		            fmt.Println(changed, "state change!")

		}

	*/
	var blah string
	fmt.Scanln(&blah)
}
