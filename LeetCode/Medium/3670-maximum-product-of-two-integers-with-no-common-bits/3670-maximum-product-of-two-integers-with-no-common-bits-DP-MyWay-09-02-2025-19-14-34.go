func maxProduct(nums []int) int64 {
	// operations = 0
	maxN := 0
	Pairs := make(map[int]int)
	for _, x := range nums {
		maxN = max(maxN, x)
		Pairs[x] = x
	}
	msb := int(math.Log2(float64(maxN)))
	maxMask := (1 << (msb + 1)) - 1
	q := make([]int,maxMask+1)
	for _, v := range Pairs {
		q [v] = v
	}
	for i,v := range q {
		if v != 0 {
			for b := range msb+1 {
				// operations++
				index := i|(1<<b)
				if q[index] < v {
					q[index] = v
					// for _,x := range q {
					// 	fmt.Fprintf(f,"%2d ",x)
					// }
					// fmt.Fprintf(f,"\n")
				}
			}
		}
	}
	// for i := range len(q) {
	// 	fmt.Fprintf(f,"%2d ",i)
	// }
	// fmt.Fprintf(f,"\n")

	maxProduct := 0
	for _, i := range Pairs {
		key := maxMask^i
		// logger.Printf("i %d %b key %d %b v %d",i,i,key,key,q[key])
		maxProduct = max(maxProduct, i*q[key])
	}
	// logger.Printf("input length %d uniqueNum %d SC %d TC %d", len(nums), len(Pairs), len(Pairs), operations)
	return int64(maxProduct)
}