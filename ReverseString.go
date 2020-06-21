func reverseString(s []byte) []string {
	result := recurse(0, s)
	fmt.Print(arraySplit)
	return arraySplit
}

func recurse(index int, str []byte) string {
	if index >= len(str) {
			result := ""
			return result
	}
	result := recurse(index + 1, str)
	result = result + string(str[index])
	return result
}