package main

import (
	"fmt"
	"time"
)

func main() {
	TN := time.Now()
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from Docker")
		if i == 4 {
			break
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println(time.Since(TN))
}
