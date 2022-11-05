package accenter

import (
	"math/rand"

	wiki "example.com/accenter/pkg/wiki"
)

// given a map of InfoWord
// pick one according to some logic
func ExtractWord(m map[wiki.Word]*InfoWord) wiki.Word {
	return pick(m)
}

// Pick a random key in a map.
//
// https://www.reddit.com/r/golang/comments/kiees6/comment/ggs5z6l
func pick[K comparable, V any](m map[K]V) K {
	i := rand.Intn(len(m))
	for k := range m {
		if i == 0 {
			return k
		}
		i--
	}
	panic("unreachable")
}
