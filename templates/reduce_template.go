package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/reduce.go gen "ReduceType=int,int64,float32,float64"

type ReduceType generic.Type

func ReduceReduceType(elements []ReduceType, fn func(accm ReduceType, curr ReduceType) ReduceType) ReduceType {
	var accm ReduceType
	for _, elem := range elements {
		accm = fn(accm, elem)
	}
	return accm
}
