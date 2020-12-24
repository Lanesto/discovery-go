package exercise5

import (
	"fmt"
	"strings"
)

// MultiSet is simple type alias to map string->int
type MultiSet map[string]int

// NewMultiSet returns new MultiSet instance
func NewMultiSet() MultiSet {
	return make(MultiSet)
}

// Insert insert item to multiset
func Insert(m MultiSet, val string) {
	m[val]++
}

// Erase remove item from multiset
// noop if val is not in multiset
func Erase(m MultiSet, val string) {
	if cnt, ok := m[val]; ok {
		m[val] = cnt - 1
		if m[val] < 0 {
			m[val] = 0
		}
	}
}

// Count returns number of items
func Count(m MultiSet, val string) int {
	return m[val]
}

// String returns formatted string wrapping elements,
// space delimted between { }
func String(m MultiSet) string {
	var str string
	for k, v := range m {
		str += strings.Repeat(fmt.Sprintf(" %s", k), v)
	}
	ret := fmt.Sprintf("{%s }", str)
	return ret
}
