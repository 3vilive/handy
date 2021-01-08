package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/min.go gen "MinType=int,int64,float32,float64"

type MinType generic.Type

func MinMinType(a, b MinType) MinType {
	if a < b {
		return a
	}
	return b
}
