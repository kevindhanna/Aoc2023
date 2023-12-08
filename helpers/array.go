package helpers

import (
	"errors"
	"strings"
)

func Reverse[T any](items []T) []T {
	itemsCopy := make([]T, len(items))
	copy(itemsCopy, items)
	for i, j := 0, len(itemsCopy)-1; i < j; i, j = i+1, j-1 {
		itemsCopy[i], itemsCopy[j] = itemsCopy[j], itemsCopy[i]
	}
	return itemsCopy
}

func Map[T any, U any](items []T, apply func(item T, i int) U) []U {
	results := []U{}
	for i, item := range items {
		results = append(results, apply(item, i))
	}
	return results
}

func Reduce[T any, U any](items []T, apply func(result U, item T, i int) U, initial U) U {
	result := initial
	for i, item := range items {
		result = apply(result, item, i)
	}

	return result
}

func SplitMap[T any](input string, split string, apply func(s string, i int) (T, error)) ([]T, error) {
	slice := strings.Split(input, split)
	results := []T{}

	for i, part := range slice {
		result, err := apply(part, i)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}

	return results, nil
}

func Filter[T any](items []T, predicate func(item T) bool) []T {
	results := []T{}
	for _, item := range items {
		if predicate(item) {
			results = append(results, item)
		}
	}
	return results
}

func FindIndexComparable[T comparable](items []T, needle T) (int, error) {
	return FindIndex(items, func(item T) bool {
		return needle == item
	})
}

func FindIndex[T any](items []T, predicate func(item T) bool) (int, error) {
	for i, item := range items {
		if predicate(item) {
			return i, nil
		}
	}
	return -1, errors.New("Coulnd't find matching item")
}

func Every[T any](items []T, predicate func(item T) bool) bool {
	for _, item := range items {
		if !predicate(item) {
			return false
		}
	}
	return true
}
