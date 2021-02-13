# talib
A Golang wrapper for TA-LIB (Windows x64 & Linux)

## Example
```go
package main

import (
	"fmt"
	"math"

	"github.com/frankrap/talib"
)

func main() {
	fmt.Println(talib.Sin([]float64{0, math.Pi / 2}))
	// => [0 1]
}
```

##### Windows
## Build "talib.dll" use VS2013
Open ta-lib\c\ide\vs2013\dll_proj\talib.sln and Build. That's All.

##### Linux

Install from your package manager or install from source.

On Arch Linux `ta-lib` is available from the AUR.

To compile first download [ta-lib-0.4.0-src.tar.gz](http://prdownloads.sourceforge.net/ta-lib/ta-lib-0.4.0-src.tar.gz) and:
```
$ wget http://prdownloads.sourceforge.net/ta-lib/ta-lib-0.4.0-src.tar.gz
$ tar -xzf ta-lib-0.4.0-src.tar.gz
$ cd ta-lib
$ ./configure --prefix=/usr LDFLAGS="-lm"
$ make
$ sudo make install
```

> If you build ``TA-Lib`` using ``make -jX`` it will fail but that's OK!
> Simply rerun ``make -jX`` followed by ``[sudo] make install``.
