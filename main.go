package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().String())
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(<-timer.C)
}
