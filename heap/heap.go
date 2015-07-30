package heap

import ()

const (
	AlmostInf int = 0xFFFFFFFFFFFFFFF
)

// Only works on integer arrays
// containing no zeroes
type ArrayHeap struct {
	es []int
}

func Array() ArrayHeap {
	return ArrayHeap{}
}

func (t *ArrayHeap) Add(x int) {
	l := len(t.es)
	L := t.L()

	if L < l {
		t.es[L] = x
	} else {
		t.es = append(t.es, x)
	}

	t.PrioritizeUp()
}

func (t *ArrayHeap) PrioritizeUp() {
	i := t.L() - 1
	e := t.es[i]

	PrioritizeUpRec(t.es, i, e)
}

func PrioritizeUpRec(es []int, i int, e int) {
	pI := ParentI(i)
	if pI < 0 {
		return
	}

	pE := es[pI]

	if pE < e {
		return
	} else {
		// Swap current element in current index
		// for the parent element
		es[i] = pE
		es[pI] = e

		// Move up to parent position
		PrioritizeUpRec(es, pI, e)
	}
}

func ParentI(x int) int {
	if x == 0 {
		return -1
	}

	return (x - 1) / 2
}

func LeftI(x int) int {
	return (2 * x) + 1
}

func RightI(x int) int {
	return (2 * x) + 2
}

func (t *ArrayHeap) Pop() int {
	m := t.es[0]
	L := t.L()
	i := L - 1

	e := t.es[i]
	t.es[0] = e
	t.es[i] = AlmostInf

	t.PrioritizeDown(L)

	return m
}

func (t *ArrayHeap) L() int {
	L := len(t.es)

	// Finding Actual
	for (L > 0) && t.es[L-1] == AlmostInf {
		L--
	}

	return L
}

func (t *ArrayHeap) PrioritizeDown(L int) {
	e := t.es[0]
	PrioritizeDownRec(t.es, 0, e, L)
}

func PrioritizeDownRec(es []int, i, e, L int) {
	iL := LeftI(i)
	iR := RightI(i)

	// Both iL and iR are valid
	if iL < L && iR < L {
		eL := es[iL]
		eR := es[iR]

		if eL > eR {
			// e < eR < eL
			if e < eR {
				return
			} else {
				es[i] = eR
				es[iR] = e

				PrioritizeDownRec(es, iR, e, L)
			}
		} else {
			// e < eL < eR
			if e < eL {
				return
			} else {
				es[i] = eL
				es[iL] = e

				PrioritizeDownRec(es, iL, e, L)
			}
		}
		return
	}

	if iL < L {
		eL := es[iL]
		if e < eL {
			return
		} else {
			es[i] = eL
			es[iL] = e

			PrioritizeDownRec(es, iL, e, L)
		}
	}
}
