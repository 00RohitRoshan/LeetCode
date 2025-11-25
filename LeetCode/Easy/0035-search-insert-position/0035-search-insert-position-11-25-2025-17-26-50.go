func searchInsert(nums []int, target int) int {
    i := 0
    j := len(nums)-1
    for i <= j {
        mid := i + (j-i)/2
        fmt.Println(i,mid,j)
        if nums[mid] > target {
            j = mid-1
        }else if nums[mid] < target  {
            i = mid+1
        }else {
            return mid
        }
    }
    return j+1
}