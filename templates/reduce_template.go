package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/reduce.go gen "ReduceType=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,uintptr,string,float32,float64,bool"

type ReduceType generic.Type

func ReduceReduceType(elements []ReduceType, fn func(accm ReduceType, curr ReduceType) ReduceType) ReduceType {
	var accm ReduceType
	for _, elem := range elements {
		accm = fn(accm, elem)
	}
	return accm
}
