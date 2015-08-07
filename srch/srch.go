package srch

// CompareFunc compare an obligation to a desired value
// < 0 means the Obligation < value
// 0 means the Obligation == value
// > 0 means the Obligation > value
type CompareFunc func(int) int

// This function is the callback to tell
// the caller to store the result at the
// index past back
type StoreFunc func(int)

// Find a Single Obligation using binary search O(log(n))
// returns the *Obligation,
// and the index of *Obligation
func singleBinarySearch(A, B int, CF CompareFunc) int {
	// Base case A is greater than B
	if A >= B {
		if CF(B) == 0 {
			return B
		}
		return -1
	}

	m := (A+B) / 2
	r := CF(m)

	// Current Obligation is too small move right
	if r < 0 {
		return singleBinarySearch(m+1, B, CF)
	}

	// Current Obligation is too large move left
	if r > 0 {
		return singleBinarySearch(A, m, CF)
	}

	return m
}

// Works for always increasing lists
func SingleBinarySearch(L int, CF CompareFunc) int {
	if L < 1 {
		return -1
	}
	return singleBinarySearch(0, L-1, CF)
}

// Find a Multiple Obligations using binary search O(log(n) + k)
// k is the number of results found
// returns a list of Obligations
// Unlink singleBinarySearch it works for simply any sorted
// list
// The correct design is to add results inside the compare function
func MultBinarySearch(L int, CF CompareFunc, SF StoreFunc) {
	if L < 1 {
		return
	}
	idx := singleBinarySearch(0, L-1, CF)

	if idx < 0 {
		return
	}
	SF(idx)

	// Search Left
	i := idx - 1
	for {
		if i < 0 {
			break
		}

		r := CF(i)
		if r != 0 {
			break
		}
		SF(i)

		i--
	}

	// Search Right
	i = idx + 1
	for {
		if i >= L {
			break
		}

		r := CF(i)
		if r != 0 {
			break
		}
		SF(i)

		i++
	}
}
