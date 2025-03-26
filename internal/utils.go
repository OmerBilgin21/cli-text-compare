package internal

func GetBiggerArray(one []string, two []string) ([]string, []string) {
	if len(one) > len(two) {
		return one, two
	}
	return two, one
}
