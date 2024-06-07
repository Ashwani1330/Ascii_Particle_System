package main

import (
	"fmt"
	particles "mymodule"
	"time"
)

func main() {
    coffee := particles.NewCoffee(5, 3)
    coffee.Start()

    timer := time.NewTicker(100 * time.Millisecond)
    for {
        <-timer.C
        fmt.Print("\033[H\033[2J")
        coffee.Update()
        fmt.Println(coffee.Display())
    }
}
