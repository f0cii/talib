# talib
A Golang wrapper for TA-LIB

## Example
```go
package main

import (
	"fmt"
	"math"

	"github.com/nvsoft/talib"
)

func main() {
	fmt.Println(talib.Sin([]float64{0, math.Pi / 2}))
	// => [0 1]
}
```
