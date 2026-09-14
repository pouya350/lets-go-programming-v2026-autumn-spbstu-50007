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
}
func Summa(a, b int) (int, error) {
	return a + b, nil
}
func Diff(a, b int) (int, error) {
	return a - b, nil
}
func Multip(a, b int) (int, error) {
	return a * b, nil
}
func Operation(a ,b int ,opertion string)(int, error){
	var res int
	var err error
	switch opertion{
	case "+":
		res, err = Summa(a,b)
		return res, err
	case "-":
		res, err = Diff(a,b)
		return res, err
	case "*":
		res, err = Multip(a,b)
		return res, err
	case "/":
		res,err = Divide(a,b)
		return res, err
	default:
		return 0, fmt.Errorf("Invalid operation")
	}
}

func main() {
	
	
}
