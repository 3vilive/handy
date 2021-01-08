package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/filter.go gen "FilterType=int,int64,byte,string,float32,float64"

type FilterType generic.Type

func FilterFilterType(elements []FilterType, fn func(FilterType) bool) []FilterType {
	var filtered []FilterType
	for _, elem := range elements {
		if fn(elem) {
			filtered = append(filtered, elem)
		}
	}
	return filtered
}
