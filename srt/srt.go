package srt

// QuickSort ...
func QuickSort(A []int) {
	L := len(A)

	QuickSortRec(A, 0, L-1)
}

// QuickSortRec ...
func QuickSortRec(A []int, a, b int) {
	if a >= b {
		return
	}

	m := MedianIndex(A, a, b)
	QuickSortRec(A, a, m)
	QuickSortRec(A, m+1, b)
}

// MedianIndex ...
func MedianIndex(A []int, a, b int) int {
	return 1
}

// MergeSort is a handy wrapper for
// MergeSortRec. It will sort the
// entire array A.
func MergeSort(A []int) {
	L := len(A)

	C := make([]int, L)

	MergeSortRec(A, C, 0, L-1)
}

// MergeSortRec sorts a given array
// between the points A[a, b] it uses
// C for copying values in Merge
func MergeSortRec(A, C []int, a, b int) {
	if (b - a) < 2 {
		if A[b] < A[a] {
			t := A[a]
			A[a] = A[b]
			A[b] = t
		}
		return
	}

	m := (a + b) / 2

	MergeSortRec(A, C, a, m)
	MergeSortRec(A, C, m+1, b)

	Merge(A, C, a, m, b)
}

// Merge merges two sorted arrays located
// at A[a, m] and A[m + 1, b] the result
// being a sorted A array at A[a, b]
// Merge uses C to store the sorted
// array before copying it over to the
// original array, A
func Merge(A, C []int, a, m, b int) {
	ai := 0
	bi := 0
	L := b - a + 1

	for i := 0; i < L; i++ {
		if (a + ai) > m {
			C[i] = A[m+1+bi]
			bi++
			continue
		}
		if (m + 1 + bi) > b {
			C[i] = A[a+ai]
			ai++
			continue
		}

		if A[a+ai] < A[m+1+bi] {
			C[i] = A[a+ai]
			ai++
		} else {
			C[i] = A[m+1+bi]
			bi++
		}
	}

	for i := 0; i < L; i++ {
		A[a+i] = C[i]
	}
}
