package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/sum.go gen "SumType=int,int64,float32,float64"

type SumType generic.Type

func SumSumType(values ...SumType) SumType {
	return ReduceSumType(values, func(accm SumType, curr SumType) SumType { return accm + curr })
}
