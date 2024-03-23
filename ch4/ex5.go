package ch4

func RemoveDups(arr []string) []string {
	n := len(arr)

	for i := 0; i < n; i++ {
		str := arr[i]
		j := i + 1
		for ; j < n && arr[j] == str; j++ {
		}

		copy(arr[i+1:], arr[j:])
		n -= j - i - 1

		arr = arr[0:n]
	}

	return arr
}
