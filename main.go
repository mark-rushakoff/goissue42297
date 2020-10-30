package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.FormatFloat(0.5, 'f', 1, 64)
	_, err := strconv.ParseFloat(s, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println("ok")
}
