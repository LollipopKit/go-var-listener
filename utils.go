package govarlistener

func Contains[T string | int | int64 | float64](a []T, b T) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}
