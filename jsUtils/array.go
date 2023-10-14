package jsUtils

import (
	"errors"
)

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

// Reduce applies a reducer function to a list of elements and returns a single result.
func Reduce[T comparable](arr []T, reducer func(accumulator T, currentValue T) (T, error), initialValue T) (T, error) {
	accumulator := initialValue
	for _, value := range arr {
		result, err := reducer(accumulator, value)
		if err != nil {
			return accumulator, err
		}
		accumulator = result
	}
	return accumulator, nil
}

// LastIndexOf finds the index of the last occurrence of a value in a slice.
func lastIndexOf(arr []interface{}, value interface{}) (int, error) {
	if value == nil {
		return -1, errors.New("element not found: value is nil")
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if value == arr[i] {
			return i, nil
		}
	}
	return -1, errors.New("element not found")
}
