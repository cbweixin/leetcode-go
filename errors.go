package main

import (
	"errors"
	"fmt"
	"math"
)

var errNotFound error = errors.New("Not found error")

func Sqrt1(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}
	// implementation of Sqrt
	return math.Sqrt(f), nil
}

func main() {

	fmt.Printf("error: %v", errNotFound)
	if _, err := Sqrt1(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

}
