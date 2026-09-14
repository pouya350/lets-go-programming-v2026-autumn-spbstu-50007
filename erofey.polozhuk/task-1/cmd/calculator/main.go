package main

import (
	"fmt"
	"strconv"
)

func Divide(a, b int) (int, error) {
	if b != 0 {
		return a / b, nil
	}
	return 0, fmt.Errorf("Division by zero")

func main() {
	
	
}
