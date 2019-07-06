package talib

import (
	"fmt"
	"testing"
)

func Test_Acos(t *testing.T) {
	inReal := []float64{0.5, 0.6, 0.7}

	outReal := Acos(inReal)

	fmt.Printf("outReal=%v\n", outReal)
}

func Test_Max(t *testing.T) {
	//Initialize()

	inReal := []float64{1, 2, 3, 4, 5, 3, 43.2, 3.2, 3.13}

	optInTimePeriod := 5

	out := Max(inReal, optInTimePeriod)
	fmt.Printf("out=%v\n", out)

	//Shutdown()
}
