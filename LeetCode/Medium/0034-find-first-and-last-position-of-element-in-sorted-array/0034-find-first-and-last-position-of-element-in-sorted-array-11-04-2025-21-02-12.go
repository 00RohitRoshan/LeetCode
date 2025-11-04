func searchRange(nums []int, target int) []int {
    i,j := -1,-1
    k := find(nums,target)
    if k == -1 {
        return []int{i,j}
    }
    i,j = k,k
    l,m := k,k
    for l<len(nums) && nums[l] == target {
        j=l
        l++
    }
    for m>-1 && nums[m] == target {
        i=m
        m--
    }
    return []int{i,j}
}

func find(nums []int,target int)  int {
    i,j := 0,len(nums)-1
    for i<=j {
        m:= i+((j-i)/2)
        // fmt.Println(m,nums[m])
        if nums[m] > target {
            j = m-1
        }else if nums[m] < target{
            i = m+1
        }else{
            return m
        }
    }
    return -1
}