const (
	base int = iota
	curr
)
func checkInclusion(s1 string, s2 string) bool {
	if len(s1)>len(s2){
		return false
	}
	if s1==s2{
		return true
	}

	//hm := map[rune][]int{}

	var counter []int
	var sum int
	counter, sum = getCounter(s1)

/*	for _, c := range s1{
		counter[c-'a']+=1
		sum+=1
/*
		hm[c]=[]int{0,0}
		hm[c][base]+=1
		hm[c][curr]+=1
		sum+=1

	} */

	front := 0
	origSum := sum

	for front+len(s1) <= len(s2){
		for i:=front; i<front+len(s1); i++ {
			c := s2[i]
			if counter[c-'a']==0 {
				break
			} 
			counter[c-'a']--
			sum--
		}
		if sum == 0 {
			return true
		}
		if sum != origSum {
			counter, sum = getCounter(s1)
		}
		front++
	}
	return false

/*	for front+len(s1) <= len(s2){
		for i:=front; i<front+len(s1); i++ {
			c := rune(s2[i])
			if _,ok := hm[c]; ok {
				if hm[c][curr]==0{
					break
				}
				hm[c][curr]--
				sum--
			}
		}
		if sum == 0 {
			return true
		}
		reset(hm)
		sum = origSum
		front++
	}
	return sum == 0

/*
	i := 0
	origSum := sum
	for i<len(s2) {
		c := rune(s2[i])
		if _, ok := hm[c]; ok {
			if hm[c][curr]==0{
				reset(hm)
				sum = origSum
			}else{
				hm[c][curr]--
				sum--
				if sum == 0 {
					return true
				}
				i++
			}
		}else {
			if sum != origSum{
				reset(hm)
				sum = origSum
			}
			i++
		}
	}
	return false
	*/
}

func getCounter(s1 string) ([]int, int) {
	var (
		sum int
		counter = make([]int, 26)
	)
	for _, c := range s1{
		counter[c-'a']+=1
		sum+=1
	}
	return counter, sum
}
/*
func reset(hm map[rune][]int) {
	for _, v := range hm {
		v[curr]=v[base]
	}
}
*/