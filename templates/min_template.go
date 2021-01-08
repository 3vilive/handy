package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/min.go gen "MinType=int,int64,float32,float64"

type MinType generic.Type

func MinMinType(values ...MinType) MinType {
	var min MinType
	if len(values) == 0 {
		return min
	} else {
		min = values[0]
	}
	for i := range values {
		if values[i] < min {
			min = values[i]
		} 
	}
	return min
}
