package stringset

var exists = struct{}{}

type StringSet struct {
	m map[string]struct{}
}

func NewStringSet() *StringSet {
	s := &StringSet{}
	s.m = make(map[string]struct{})
	return s
}

func (s *StringSet) Add(value string) {
	s.m[value] = exists
}

func (s *StringSet) Remove(value string) {
	delete(s.m, value)
}

func (s *StringSet) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func (s *StringSet) Length() int {
	return len(s.m)
}

func (s *StringSet) ToSlice() []string {
	keys := make([]string, 0, len(s.m))

	for k, _ := range s.m {
		keys = append(keys, k)
	}
	return keys
}

func (s *StringSet) Union(other *StringSet) *StringSet {
	union := NewStringSet()

	for v, _ := range s.m {
		union.Add(v)
	}

	for v, _ := range other.m {
		union.Add(v)
	}

	return union
}

func (s *StringSet) Intersection(other *StringSet) *StringSet {
	union := s.Union(other)
	intersection := NewStringSet()

	for v, _ := range union.m {
		if s.Contains(v) && other.Contains(v) {
			intersection.Add(v)
		}
	}

	return intersection
}
