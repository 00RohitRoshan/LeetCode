func lemonadeChange(bills []int) bool {
    count := make(map[int]int)
    for _,bill := range bills {
        if bill == 5 {
            count[5]++
        }
        if bill == 10 {
            if count[5] > 0 {
                count[5]--
                count[10]++
            }else{
                return false
            }
        }
        if bill == 20 {
            if count[5] > 0 && count[10] > 0 {
                count[5]--
                count[10]--
            }else if count[5] > 2 {
                count[5] -= 3
            }else {
                return false
            }
        }
    }
    return true
}