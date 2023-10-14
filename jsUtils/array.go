package jsUtils

import "errors"

// LastElement returns the last element of the slice.
func lastElement(arr []interface{}) interface{} {
	return arr[len(arr)-1]
}

// Filter returns a new slice with elements that pass the predicate.
func Filter(arr []interface{}, predicate func(interface{}) bool) []interface{} {
	var newArr []interface{}
	for _, value := range arr {
		if predicate(value) {
			newArr = append(newArr, value)
		}
	}
	return newArr
}

// T indexOf returns the index of a given element
func indexOf[T comparable](arr []T, value T) (int, error) {
	for index, value := range arr {
		if value == value {
			return index, nil
		}
	}
	return -1, errors.New("element not found")
}
