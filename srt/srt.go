package srt

func MergeSort(A []int) {
	L := len(A)

	C := make([]int, L)

	MergeSortRec(A, C, 0, L - 1)
}

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
	MergeSortRec(A, C, m + 1, b)

	Merge(A, C, a, m, b)
}

func Merge(A, C []int, a, m, b int) {
	ai := 0
	bi := 0
	L := b-a+1

	for i := 0; i < L; i++ {
		if (a + ai) > m {
			C[i] = A[m + 1 + bi]
			bi++
			continue
		}
		if (m + 1 + bi) > b {
			C[i] = A[a + ai]
			ai++
			continue
		}

		if A[a + ai] < A[m + 1 + bi] {
			C[i] = A[a + ai]
			ai++
		} else {
			C[i] = A[m + 1 + bi]
			bi++
		}
	}

	for i := 0; i < L; i++ {
		A[a + i] = C[i]
	}
}
