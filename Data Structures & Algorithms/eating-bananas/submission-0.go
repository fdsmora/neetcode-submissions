import "slices"

func minEatingSpeed(piles []int, h int) int {
	max := slices.Max(piles)
	min := 1
	var k, res int
	for min <= max {
		k = (min+max)/2
		total := sum(k, piles)
		if total <= h {
			res = k
			max = k -1
		} else {
			min = k +1
		}
	}
	return res
}

func sum(k int, piles []int) int {
	total := 0.0
	for _, v := range piles {
		total += math.Ceil(float64(v)/float64(k))
	}
	return int(total)
}
