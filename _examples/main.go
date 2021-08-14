package main

import (
	"fmt"
	"time"

	timeago "github.com/caarlos0/timea.go"
)

func main() {
	fmt.Println(timeago.Of(time.Now().Add(-5 * time.Second)))
}
