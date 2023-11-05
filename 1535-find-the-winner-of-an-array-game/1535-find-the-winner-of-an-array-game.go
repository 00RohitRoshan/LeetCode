package main

func getWinner(arr []int, k int) int {
    if k == 1 {
        return max(arr[0], arr[1])
    }
    if k >= len(arr) {
        return maxArr(arr)
    }

    current_winner := arr[0]
    consecutive_wins := 0

    for i := 1; i < len(arr); i++ {
        if current_winner > arr[i] {
            consecutive_wins++
        } else {
            current_winner = arr[i]
            consecutive_wins = 1
        }

        if consecutive_wins == k {
            return current_winner
        }
    }

    return current_winner
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxArr(arr []int) int {
    m := arr[0]
    for _, val := range arr {
        if val > m {
            m = val
        }
    }
    return m
}