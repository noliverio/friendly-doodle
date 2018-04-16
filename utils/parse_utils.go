package utils

func KVParse(input []byte) map[string]string {
	ampersandPos := make([]int, len(input)/3)
	equalPos := 0
	count := 0
	for pos, val := range input {
		if val == byte('&') {
			ampersandPos[count] = pos
			count++
		}
	}
	if count == 0 {
		for pos, val := range input {
			if val == byte('=') {
				equalPos = pos
			}
		}
	}

	returnMap := make(map[string]string, equalPos)

	return returnMap

}
