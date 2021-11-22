package utils

func ArrayContains(arr []string, v string) bool {
	for _, a := range arr {
		if v == a {
			return true
		}
	}

	return false
}
