package intset

var exists = struct{}{}

type intSet struct {
	m map[int]struct{}
}

func NewIntSet() *intSet {
	s := &intSet{}
	s.m = make(map[int]struct{})
	return s
}

func (s *intSet) Add(value int) {
	s.m[value] = exists
}

func (s *intSet) Remove(value int) {
	delete(s.m, value)
}

func (s *intSet) Contains(value int) bool {
	_, c := s.m[value]
	return c
}
