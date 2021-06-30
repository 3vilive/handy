package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/pow.go gen "PowType=int,int64"

type PowType generic.Type

func PowPowType(x, y PowType) PowType {
	var (
		z PowType = 1
		i PowType = 0
	)
	for ; i < y; i++ {
		z = z * x
	}
	return z
}
