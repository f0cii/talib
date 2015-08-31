package talib

import (
	"testing"
	"fmt"
)

func Test_Max(t *testing.T) {
	inReal := []float64{1, 2, 3, 4, 5, 3, 43.2, 3.2, 3.13}

	optInTimePeriod := 5

	out := Max(inReal, optInTimePeriod)
	fmt.Printf("out=%v\n", out)
}
