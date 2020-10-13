package utils

func StringRemoveItemFromArray(arr []string, removeEl string) []string {
	var s []string

	for _, el := range arr {
		if removeEl == el {
			continue
		}

		s = append(s, el)
	}

	return s
}
