package main

func Quicksort(arr *[]int) {
	qs(arr, 0, len(*arr)-1)
}

func qs(arrP *[]int, l int, r int) {
	if l < r {
		part := partition(arrP, l, r)
		qs(arrP, l, part-1)
		qs(arrP, part+1, r)
	}
}

func partition(arrP *[]int, l int, r int) int {
	arr := *arrP
	swap := l
	for i := l; i < r; i++ {
		if arr[i] < arr[r] {
			arr[swap], arr[i] = arr[i], arr[swap]
			swap++
		}
	}
	arr[swap], arr[r] = arr[r], arr[swap]
	return swap
}
