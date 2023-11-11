package util

func Filter[T any](slice []T, predicate func(t T) bool) []T {
	if slice == nil {
		return slice
	}

	result := make([]T, 0, len(slice))
	for _, value := range slice {
		if predicate(value) {
			result = append(result, value)
		}
	}

	return result
}

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}
