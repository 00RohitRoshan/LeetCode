// ---------- State definition ----------
type State struct {
	sum int
	idx []int
}

// ---------- MinHeap definition ----------
type MinHeap []State

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(State)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ---------- Helper to stringify indices (for visited set) ----------
func idxKey(idx []int) string {
	parts := make([]string, len(idx))
	for i, v := range idx {
		parts[i] = strconv.Itoa(v)
	}
	return strings.Join(parts, ",")
}

// ---------- Main function ----------
func kthSmallest(mat [][]int, k int) int {
	n := len(mat)

	// initial state: all indices = 0
	startIdx := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		sum += mat[i][0]
	}

	start := State{sum: sum, idx: startIdx}

	h := &MinHeap{start}
	heap.Init(h)

	seen := make(map[string]bool)
	seen[idxKey(startIdx)] = true

	for popped := 0; popped < k; popped++ {
		cur := heap.Pop(h).(State)

		// if this is the k-th pop, return sum
		if popped == k-1 {
			return cur.sum
		}

		// generate neighbors
		for i := 0; i < n; i++ {
			if cur.idx[i]+1 < len(mat[i]) {
				newIdx := append([]int(nil), cur.idx...)
				newIdx[i]++
				key := idxKey(newIdx)

				if !seen[key] {
					newSum := cur.sum - mat[i][cur.idx[i]] + mat[i][newIdx[i]]
					heap.Push(h, State{sum: newSum, idx: newIdx})
					seen[key] = true
				}
			}
		}
	}

	return -1 // should never reach
}