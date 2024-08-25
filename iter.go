//go:build go1.23
// +build go1.23

package gocitiesjson

import (
	"iter"
	"math/rand"
)

func All(shuffle bool) iter.Seq[*City] {
	return func(yield func(*City) bool) {
		var idxes []int
		if shuffle {
			idxes = make([]int, len(idx))
			copy(idxes, idx)
			rand.Shuffle(len(idxes), func(i, j int) { idxes[i], idxes[j] = idxes[j], idxes[i] })
		} else {
			idxes = idx
		}
		for _, i := range idxes {
			if !yield(Cities[i]) {
				return
			}
		}
	}
}

func Filter(fn func(*City) bool, shuffle bool) iter.Seq[*City] {
	return func(yield func(*City) bool) {
		for city := range All(shuffle) {
			if fn(city) {
				if !yield(city) {
					return
				}
			}
		}
	}
}
