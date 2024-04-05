package utils

func Ternary(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}

	return b
}
