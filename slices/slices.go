package slices

// -----------------------------------------------------------------------------

// StringInArray checks if `sl` contains `s` and returns `true` in that case,
// false otherwise.
func StringInArray(sl []string, s string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false
}
