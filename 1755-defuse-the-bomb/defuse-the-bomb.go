func decrypt(code []int, k int) []int {
	length := len(code)
	result := make([]int, length)

	kIndex := func(startIndex, i int) int {
		return (startIndex + i + length) % length
	}

	prefixWindow, i := 0, k

	if k == 0 {
		return result
	}

	for i != 0 {
		prefixWindow += code[kIndex(0, i)]
		if i > 0 {
			i--
		} else {
			i++
		}
	}

	result[0] = prefixWindow

	for j := 1; j < length; j++ {
		if k > 0 {
			prefixWindow += code[kIndex(j, k)]
		} else {
			prefixWindow += code[kIndex(j, -1)]
		}
		if k > 0 {
			prefixWindow -= code[kIndex(j, 0)]
		} else {
			prefixWindow -= code[kIndex(j-1, k)]
		}
		result[j] = prefixWindow
	}

	return result
}