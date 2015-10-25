package main

func InsertionSort(vals []int, beg, end int) {
	size := end - beg
	for i := beg; i < beg+size; i++ {
		for j := i + 1; j <= size+beg; j++ {
			if vals[i] > vals[j] {
				vals[i], vals[j] = vals[j], vals[i]
			}
		}
	}
}

func partition(vals []int, beg, end int) int {
	vals[beg], vals[(beg+end)/2] = vals[(beg+end)/2], vals[beg]

	total := end - beg
	low := 0
	high := 0

	for low+high < total {
		for low+high < total && vals[beg+1+low] <= vals[beg] {
			low++
		}
		for low+high < total && vals[end-high] >= vals[beg] {
			high++
		}
		if low+high < total {
			vals[beg+1+low], vals[end-high] = vals[end-high], vals[beg+1+low]
		}
	}

	vals[beg+low], vals[beg] = vals[beg], vals[beg+low]
	return beg + low
}

func quicksort(vals []int, beg, end int) {
	if end-beg <= 8 && end-beg >= 3 {
		InsertionSort(vals, beg, end)
		return
	}

	part := partition(vals, beg, end)

	if part-1-beg > 2 {
		quicksort(vals, beg, part-1)
	}

	if end-part+1 > 2 {
		quicksort(vals, part+1, end)
	}
}
