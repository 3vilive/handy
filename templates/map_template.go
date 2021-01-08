package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/map.go gen "MapInputType=int,int64,byte,string,float32,float64,bool MapOutputType=int,int64,byte,string,float32,float64,bool"

type MapInputType generic.Type
type MapOutputType generic.Type

func MapMapInputTypeToMapOutputType(elements []MapInputType, fn func(MapInputType) MapOutputType) []MapOutputType {
	out := make([]MapOutputType, 0, len(elements))
	for _, elem := range elements {
		out = append(out, fn(elem))
	}
	return out
}
