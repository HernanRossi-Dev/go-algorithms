func reverseString(s []byte) {
	length := len(s)
	swapCount := length / 2
	
	for i := 0; i < swapCount; i++ {
			temp := s[i]
			s[i] = s[length - i -1]
			s[length - 1 - i] = temp
	}
}