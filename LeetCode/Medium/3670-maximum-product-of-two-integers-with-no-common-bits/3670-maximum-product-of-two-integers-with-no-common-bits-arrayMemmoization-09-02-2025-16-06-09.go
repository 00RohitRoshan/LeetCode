func maxProduct(nums []int) int64 {
	// logger.Printf("input:\n %v", nums)
	// find maximum number in nums
	maxN := 0
	availablePairs := make(map[int]int)
	for _, x := range nums {
		maxN = max(maxN, x)
		availablePairs[x] = x
	}
	// logger.Printf("MaxN: %d %b", maxN, maxN)

	// most significant bit
	loop := []int{}
	msb := int(math.Log2(float64(maxN)))
	maxMask := (1 << (msb + 1)) - 1
	memArray := make([]int, maxMask+1)
	for i := range memArray {
		memArray[i] = -1  // or any default value you want
	}
	memArray[0] = 0
	for _, v := range availablePairs {
		loop = append(loop, v)
		memArray[v] = v
	}
	// logger.Printf("msb %d maxMask %b", msb, maxMask)

	// uniqPairs := len(availablePairs)
	maxProduct := 0
	for _, v := range loop {
		maxPosiblePair := maxMask ^ v
		// logger.Printf("i %d %b compliment %d %b", v, v, maxPosiblePair, maxPosiblePair)
		maxProduct = max(maxProduct, v*findMaxPair(maxPosiblePair, msb, memArray))
		// fmt.Fprintln(f, memArray)
	}

	// logger.Printf("input length %d availablePairs %d memoization %d operations %d", len(nums), uniqPairs, len(availablePairs), operations)
	// fmt.Fprintln(f, memArray)
	// fmt.Fprintln(f, loop)
	return int64(maxProduct)
}
func findMaxPair(compliment int, msb int, available []int) int {
	// fmt.Fprintln(f, available)
	// logger.Printf("find %d %b",compliment,compliment)
	if available[compliment] > -1 {
		return available[compliment] 
	} 
	// operations++
	maxPair := 0
	// logger.Printf("compliment %d %b ", compliment, compliment)
	for i := range msb + 1 {
		if compliment&(1<<i) != 0 {
			submask := compliment &^ (1 << i)
			// logger.Printf("compliment %b submask %b",compliment, submask)
			maxPair = max(maxPair, findMaxPair(submask, msb, available))
		}
	}
	if available[compliment] < maxPair {
		available[compliment] = maxPair
		// logger.Printf("memoized at %d",compliment)
	}
	// logger.Printf("returning %d %b", maxPair, maxPair)
	return available[compliment]

}