package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/any.go gen "AnyType=int,int64,string"

type AnyType generic.Type

func AnyAnyType(values []AnyType, check func(AnyType) bool) bool {
	for i := range values {
		if check(values[i]) {
			return true
		}
	}

	return false
}
