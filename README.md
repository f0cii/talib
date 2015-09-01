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

## Build "talib.dll" use VS2013
Open ta-lib\c\ide\vs2013\dll_proj\talib.sln and Build. That's All.
