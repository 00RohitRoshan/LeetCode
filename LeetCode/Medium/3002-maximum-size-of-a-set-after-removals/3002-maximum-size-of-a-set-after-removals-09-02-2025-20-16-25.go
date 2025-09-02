func maximumSetSize(nums1 []int, nums2 []int) int {
    m1 := make(map[int]bool)
    m2 := make(map[int]bool)
    m3 := make(map[int]bool)
    for i := range len(nums1){
        m1[nums1[i]] = true
        m2[nums2[i]] = true
        m3[nums1[i]] = true
        m3[nums2[i]] = true
    }
    n1 := len(m1)
    n2 := len(m2)
    n3 := len(m3)
    if n1 > len(nums1)/2 {
        n1 = len(nums1)/2
    }
    if n2 > len(nums2)/2 {
        n2 = len(nums2)/2
    }
    if n3 < n1+n2 {
        return n3
    }

    return n1+n2
}