package Week_01

import "reflect"

func isAnagram(s string, t string) bool {
	h1 := map[int32]int{}
	h2 := map[int32]int{}

	for _, i := range s {
		if _, ok := h1[i]; ok {
			h1[i] += 1
		} else {
			h1[i] = 1
		}
	}

	for _, i := range t {
		if _, ok := h2[i]; ok {
			h2[i] += 1
		} else {
			h2[i] = 1
		}
	}

	return reflect.DeepEqual(h1, h2)
}
