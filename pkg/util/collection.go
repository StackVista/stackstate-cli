package util

func ContainsString(l []string, s string) bool {
	for _, n := range l {
		if s == n {
			return true
		}
	}

	return false
}

func StringSliceIndex(l []string, s string) int {
	for i, n := range l {
		if s == n {
			return i
		}
	}

	return -1
}
