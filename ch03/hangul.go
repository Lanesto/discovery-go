package hangul

var (
	start = rune(44032)
	end   = rune(55204)
)

const numConsonants = 28

// HasConsonantSuffix returns true if s has Hangual consonant at the end.
func HasConsonantSuffix(s string) bool {
	var last rune
	for _, r := range s {
		last = r
	}
	if start <= last && last < end {
		idx := int(last - start)
		return (idx % numConsonants) != 0
	}
	return false
}
