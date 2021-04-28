package main

import (
	"errors"
	"fmt"
)

const pi = 3.1415

func main() {
	printCircleArea(-3.0)
}

func printCircleArea(rad float64) {
	circleArea, err := calculateArea(rad)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Radius: %f64 sm", rad)
	fmt.Println("Folmula: S = pi*rad^2")
	fmt.Printf("Area: %f32", circleArea)
}

func calculateArea(rad float64) (float64, error) {
	if rad <= 0 {
		return float64(0), errors.New("Radius must be greater than 0!")
	}
	return float64(rad) * float64(rad) * pi, nil
}
