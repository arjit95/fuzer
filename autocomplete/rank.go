package autocomplete

const MIN_RANK = -10000

func rank(originalPattern []rune, lowerPattern []rune, originalWord []rune, lowerWord []rune) (int, []int) {
	score := 0

	patternLength := len(originalPattern)
	patternIdx := 0
	lastMatch := -1

	highlights := make([]int, 0)

	if patternLength > len(originalWord) {
		return MIN_RANK, highlights
	}

	for i, c := range lowerWord {
		if patternIdx >= patternLength {
			break
		}

		if c == lowerPattern[patternIdx] {
			score += 100
			if lastMatch+1 == i { // adjacency bonus
				score += 150
			}

			if originalPattern[patternIdx] == originalWord[i] { // case bonus
				score += 50
			}

			if patternIdx == i { // position bonus
				score += 150
			}

			highlights = append(highlights, i)
			patternIdx++
			lastMatch = i
		} else if lastMatch > -1 {
			score -= (i - lastMatch) * 25
		}
	}

	// Negate the score if no matches are found or complete pattern is not found upto 2 typos
	if lastMatch == -1 || patternIdx+2 < patternLength {
		return MIN_RANK, highlights
	}

	if patternIdx >= patternLength { // Complete pattern match bonus
		score = score + 200 - len(originalWord) // A smaller word should have higher priority
	} else {
		score -= 10 * (patternLength - patternIdx)
	}

	return score, highlights
}
