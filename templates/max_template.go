package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/max.go gen "MaxType=int,int64,float32,float64"

type MaxType generic.Type

func MaxMaxType(values ...MaxType) MaxType {
	var max MaxType

	if len(values) == 0 {
		return max
	} else {
		max = values[0]
	}

	for i := range values {
		if values[i] > max {
			max = values[i]
		}
	}

	return max
}
