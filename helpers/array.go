package helpers

import (
	"errors"
	"strings"
)

func Map[T any, U any](items []T, apply func(item T) U) []U {
	results := []U{}
	for _, item := range items {
		results = append(results, apply(item))
	}
	return results
}

func Reduce[T any, U any](items []T, apply func(result U, item T) U, initial U) U {
	result := initial
	for _, item := range items {
		result = apply(result, item)
	}

	return result
}

func SplitMap[T any](input string, split string, apply func(s string) (T, error)) ([]T, error) {
	slice := strings.Split(input, split)
	results := []T{}

	for _, part := range slice {
		result, err := apply(part)
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

func FindIndex[T any](items []T, predicate func(item T) bool) (int, error) {
	for i, item := range items {
		if predicate(item) {
			return i, nil
		}
	}
	return -1, errors.New("Coulnd't find matching item")
}
