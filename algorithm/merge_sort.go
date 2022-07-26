package algorithm


func MergeSort(arr []int) []int {

	//	fmt.Printf("%v %v %v\n", arr, lowerBound, upperBound)
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2

	r := MergeSort(arr[:mid])
	q := MergeSort(arr[mid:])

	return Merge(r, q)
}

func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

