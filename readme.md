# handy

a handy module for go

# usage

```
go get github.com/3vilive/handy
```

```go
package main

import (
	"fmt"

	"github.com/3vilive/handy/fn"
)

func main() {
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// filter odd numbers
	intArr = fn.FilterInt(intArr, func(i int) bool { return i%2 != 0 })
	fmt.Printf("%#v\n", intArr)

	// double it
	intArr = fn.MapIntToInt(intArr, func(i int) int { return i * 2 })
	fmt.Printf("%#v\n", intArr)

	// reverse it
	fn.ReverseInt(intArr)
	fmt.Printf("%#v\n", intArr)

	// use reduce to sum
	sumOfArr := fn.ReduceInt(intArr, func(accm int, curr int) int { return accm + curr })
	fmt.Printf("sumOfArr: %d\n", sumOfArr)

	// sum
	sumOfArr2 := fn.SumInt(intArr...)
	fmt.Printf("sumOfArr2: %d\n", sumOfArr2)

	// min
	minOfArr := fn.MinInt(intArr...)
	fmt.Printf("minOfArr: %d\n", minOfArr)

	// max
	maxOfArr := fn.MaxInt(intArr...)
	fmt.Printf("maxOfArr: %d\n", maxOfArr)
}

```

## development

install genny:

```
go get github.com/cheekybits/genny
```

generate:

```
cd templates && go generate
```

## types

common types:

```
int,int64,byte,string,float32,float64
```

numberic:

```
int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64
```

stirng & byte:

```
string,byte
```

other types:

```
bool,uintptr
```