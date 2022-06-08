package util

var exists = struct{}{}

type StringSet map[string]struct{}

func NewStringSet(is []string) StringSet {
	s := StringSet{}
	s.Adds(is)
	return s
}

func (s StringSet) Add(i string) {
	s[i] = exists
}

func (s StringSet) Adds(rs []string) {
	for _, r := range rs {
		s[r] = exists
	}
}

func (s StringSet) Contains(i string) bool {
	_, ok := s[i]
	return ok
}

func (s StringSet) Delete(i string) {
	delete(s, i)
}

func (s StringSet) Copy() StringSet {
	n := StringSet{}
	for k := range s {
		n[k] = exists
	}

	return n
}

func (s StringSet) ToSlice() []string {
	rs := []string{}
	for k := range s {
		rs = append(rs, k)
	}

	return rs
}
