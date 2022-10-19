package ordenar

import (
	"sort"
)

func Ordenar(s []int) []int {
	//sort.Ints(s)
	sort.Slice(s, func(i, j int) bool {
		return s[j] > s[i]
	})
	return s
}
