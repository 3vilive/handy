package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/min.go gen "MinType=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,float32,float64"

type MinType generic.Type

func MinMinType(a, b MinType) MinType {
	if a < b {
		return a
	}
	return b
}
