package set

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=../collections/set.go gen "SetType=int,int64,string"

type SetType generic.Type

type SetTypeSet struct {
	m map[SetType]struct{}
}

func NewSetTypeSet(elements ...SetType) *SetTypeSet {
	set := SetTypeSet{
		m: make(map[SetType]struct{}),
	}
	set.Add(elements...)
	return &set
}

func (s *SetTypeSet) Add(elements ...SetType) {
	for _, element := range elements {
		s.m[element] = struct{}{}
	}
}

func (s *SetTypeSet) Contain(element SetType) bool {
	_, exists := s.m[element]
	return exists
}

func (s *SetTypeSet) Elements() []SetType {
	elements := make([]SetType, 0, len(s.m))
	for element := range s.m {
		elements = append(elements, element)
	}

	return elements
}
