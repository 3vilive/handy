package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/max.go gen "MaxType=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"

type MaxType generic.Type

func MaxMaxType(a, b MaxType) MaxType {
	if a > b {
		return a
	}
	return b
}
