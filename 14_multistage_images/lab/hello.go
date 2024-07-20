package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello, world!")                          // Print simple text on screen
	fmt.Println(os.Getenv("USER"), ", Let's be friends!") // Read Linux $USER environment variable

	duration := time.Duration(120) * time.Second // Pause for 30 seconds
	time.Sleep(duration)
}
