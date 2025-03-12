package internal

func StrInsert(arr []string, index int, value string) []string {
	if index < 0 || index > len(arr) {
		arr = append(arr, value)
	}

	// Insert value at index by slicing and using append
	arr = append(arr[:index], append([]string{value}, arr[index:]...)...)
	return arr
}
