func combinationSum(nums []int, target int) [][]int {
	res:=[][]int{}

	var dfs func([]int, int, int) 

	dfs = func(A []int, sum, j int){
		if sum > target {
			return
		}
		if sum == target { 
			tmp := make([]int, len(A))
			copy(tmp, A)
			res = append(res, tmp)
			return
		}
		for i:=j; i<len(nums); i++ {
			v := nums[i]
			A = append(A, v)
			sum += v
			// debug
			//fmt.Println(A)
			//fmt.Println(sum)
			dfs(A, sum, i)
			A = A[:len(A)-1]
			sum -= v
		}
	}
	
	B := []int{}
	dfs(B, 0, 0)
	return res
}
