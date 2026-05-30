func search(nums []int, target int) int {
    s, e := 0, len(nums)-1

    for s <= e {
        m := (s+e)/2
        if nums[m]>=nums[s]{
            if target == nums[m]{
                return m
            }
            if target == nums[s]{
                return s
            }
            if target > nums[m] || target < nums[s] {
                s = m+1
            }else{
                e = m-1
            }
        }else{
            if target == nums[m] {
                return m
            }
            if target == nums[e] {
                return e
            }
            if target > nums[e] || target < nums[m]{
                e = m-1
            }else{
                s = m+1
            }
        }
    } 
    return -1
}

/*func search(nums []int, target int) int {
    s, e := 0, len(nums)-1

    if nums[s] <= nums[e]{
        return binarySearch(nums, s, e, target)
    }

    for s < e {
        m := (s+e)/2
        if nums[m]>=nums[s]{
            s = m+1
        } else {
            e = m-1
        }
    }
    pivot := e 
    if target == nums[pivot] {
        return pivot
    }else if nums[0] < target && target < nums[pivot]{
        return binarySearch(nums, 0, pivot, target)
    }else{
        return binarySearch(nums, s+1, len(nums)-1, target)
    }

    return -1
}

func binarySearch(nums []int, s, e, target int) int {
    for s <= e {
        m := (s+e)/2
        if nums[m]==target {
            return m
        }
        if target < nums[m]{
            e = m-1
        }else{
            s = m+1
        }
    }
    return -1
}*/
