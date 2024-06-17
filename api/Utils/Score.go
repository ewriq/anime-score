package Utils

func CalculateRating(fivestar, fourstar, threestar, twostar, onestar int) float64 {
	totalScore := (5 * fivestar) + (4 * fourstar) + (3 * threestar) + (2 * twostar) + (1 * onestar)
	totalVote := fivestar + fourstar + threestar + twostar + onestar
	var rating float64
	if totalVote > 0 {
		rating = float64(totalScore) / float64(totalVote)
	} else {
		rating = 0
	}
	return rating
}
