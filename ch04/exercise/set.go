package eval

// Set for set
type Set map[interface{}]struct{}

// NewSet returns new Set instance
func NewSet(elems ...interface{}) Set {
	ret := Set{}
	for _, elem := range elems {
		ret[elem] = struct{}{}
	}
	return ret
}
