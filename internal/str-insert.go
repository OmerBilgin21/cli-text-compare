package internal

func StrInsert(arr []string, index int, value string) []string {
	if index < 0 || index > len(arr) {
		return append(arr, value)
	}

	arr = append(arr[:index], append([]string{value}, arr[index:]...)...)
	return arr
}
