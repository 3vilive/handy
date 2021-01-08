package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/divmod.go gen "DivModType=int,int64"

type DivModType generic.Type

func DivModDivModType(a, b DivModType) (DivModType, DivModType) {
	return a / b, a % b
}
