package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/max.go gen "MaxType=int,int64,float32,float64"

type MaxType generic.Type

func MaxMaxType(a, b MaxType) MaxType {
	if a > b {
		return a
	}
	return b
}
