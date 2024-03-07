package collection

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

func FindFirst[T any](slice []T, predicate func(t *T) bool) *T {
	if slice == nil {
		return nil
	}

	for _, t := range slice {
		if predicate(&t) {
			return &t
		}
	}

	return nil
}

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func ForEach[T any](slice []T, eachFunc func(t T)) {
	for _, t := range slice {
		eachFunc(t)
	}
}

func Map[O any, C any](slice []O, mapFunc func(original O) C) []C {
	mapSlice := make([]C, 0, len(slice))
	for _, o := range slice {
		mapSlice = append(mapSlice, mapFunc(o))
	}

	return mapSlice
}

func Range(start, end int) []int {
	if start <= end {
		nums := make([]int, 0, end-start+1)
		for i := start; i <= end; i++ {
			nums = append(nums, i)
		}

		return nums
	}

	nums := make([]int, 0, start-end+1)
	for i := start; i >= end; i-- {
		nums = append(nums, i)
	}

	return nums
}

func Reduce[T any](list []T, reducer func(prev T, curr T) T) T {
	copyList := make([]T, 0, len(list))
	copyList = append(copyList, list...)

	for i := 1; i < len(copyList); i += 1 {
		prev := copyList[i-1]
		curr := copyList[i]
		copyList[i] = reducer(prev, curr)
	}

	return copyList[len(copyList)-1]
}
