package gotility

// The StringSet is a utility type to manage string sets.
type StringSet map[string]struct{}

// Set registers the string in this hash set
func (s StringSet) Set(value string) {
	s[value] = struct{}{}
}

// Contains returns true if the given value is contained in this string set.
func (s StringSet) Contains(value string) bool {
	_, exists := s[value]
	return exists
}
