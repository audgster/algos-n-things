package algos

func isPermutation(a string, b string) bool {
	// permutations have the multiset of characters
	// The must be the same length, and all the characters in them must be the same
	if len(a) != len(b) {
		return false
	}

	// Approaches:
	// - sort both strings. If equal, they are permutations
	// - count characters in each string and ensure they have the same counts

	var characters [256]int

	for _, character := range a {
		characters[int(character)]++
	}
	for _, character := range b {
		characters[int(character)]--
	}

	for _, count := range characters {
		if count != 0 {
			return false
		}
	}

	return true
}
