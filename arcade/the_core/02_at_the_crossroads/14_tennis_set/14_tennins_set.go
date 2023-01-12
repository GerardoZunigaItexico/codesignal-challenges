package _4_tennis_set

func solution(score1 int, score2 int) bool {
	max, min := maxAndMin(score1, score2)

	if max == 6 && min < 5 {
		return true
	}

	return false
}

func maxAndMin(score1 int, score2 int) (max, min int) {
	if score1 > score2 {
		return score1, score2
	}
	return score2, score1
}
