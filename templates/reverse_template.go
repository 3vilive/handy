package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/reverse.go gen "ReverseType=int,int64,byte,string,float32,float64"

type ReverseType generic.Type

func ReverseReverseType(elements []ReverseType) {
	for i, j := 0, len(elements)-1; i < j; i, j = i+1, j-1 {
		elements[i], elements[j] = elements[j], elements[i]
	}
}
