package fn

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../fn/map.go gen "MapInputType=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,uintptr,string,float32,float64,bool MapOutputType=int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,uintptr,string,float32,float64,bool"

type MapInputType generic.Type
type MapOutputType generic.Type

func MapMapInputTypeToMapOutputType(elements []MapInputType, fn func(MapInputType) MapOutputType) []MapOutputType {
	out := make([]MapOutputType, 0, len(elements))
	for _, elem := range elements {
		out = append(out, fn(elem))
	}
	return out
}
