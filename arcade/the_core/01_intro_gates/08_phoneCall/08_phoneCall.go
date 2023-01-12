package _8_phoneCall

func solution(min1 int, min2_10 int, min11 int, s int) int {
	sMin1 := s - min1
	if sMin1 < 0 {
		return 0
	}
	if sMin1 == 0 {
		return 1
	}
	sMin2_10 := sMin1 - (9 * min2_10)
	//Insufficient minutes to accomplish 10
	if sMin2_10 < 0 {
		return (sMin1 / min2_10) + 1
		// Only 10 minutes of call
	} else if sMin2_10 == 0 {
		return 10
	}

	// More than 10 minutes of call
	sMin11 := sMin2_10 / min11
	return 10 + sMin11
}
