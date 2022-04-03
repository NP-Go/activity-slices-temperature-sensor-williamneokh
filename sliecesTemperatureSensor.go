package main

import (
	"fmt"
)

func main() {
	//slices template
	room1 := []float64{20, 21, 23, 25, 22}
	room2 := []float64{27, 23, 25, 20, 30, 24}
	room3 := []float64{22, 23, 24, 22}

	var HdbSlice [][]float64

	HdbSlice = append(HdbSlice, room1, room2, room3)

	for eachRoom := range HdbSlice {
		totalTemp := 0.0
		for i := range HdbSlice[eachRoom] {

			totalTemp += HdbSlice[eachRoom][i]

		}
		avg := totalTemp / float64(len(HdbSlice[eachRoom]))
		fmt.Printf("Room %v average temp is %.2f\n", eachRoom+1, avg)
	}

}
