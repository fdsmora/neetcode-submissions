import "reflect"

// LC 15. 3Sum
func threeSum(nums []int) [][]int {
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					triplet := []int{nums[i], nums[j], nums[k]}
					res = addTriplet(res, triplet)
				}
			}
		}
	}
	return res
}

func addTriplet(list [][]int, triplet []int) [][]int {

	sort.Slice(triplet, func(x, y int) bool {
		return triplet[x] <= triplet[y]
	})

	for _, elem := range list {
		sort.Slice(elem, func(x, y int) bool {
			return elem[x] <= elem[y]
		})
		if reflect.DeepEqual(triplet, elem) {
			return list
		}
	}

	list = append(list, triplet)
	return list
}

