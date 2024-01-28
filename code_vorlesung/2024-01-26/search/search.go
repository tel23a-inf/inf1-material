package searchsort

func Find(list []int, value int) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}
