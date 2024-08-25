//go:build go1.23
// +build go1.23

package gocitiesjson

import "iter"

func All() iter.Seq[*City] {
	return func(yield func(*City) bool) {
		for _, city := range Cities {
			if !yield(city) {
				return
			}
		}
	}
}

func Filter(filter func(*City) bool) iter.Seq[*City] {
	return func(yield func(*City) bool) {
		for _, city := range Cities {
			if filter(city) {
				if !yield(city) {
					return
				}
			}
		}
	}
}
