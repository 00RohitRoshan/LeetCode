```Swift
class Solution {
    func breakPalindrome(_ palindrome: String) -> String {
        if palindrome.count == 1 { return "" }
        var arr = Array(palindrome)
        for i in 0..<palindrome.count / 2 {
            if arr[i] == "a"{ continue }
            arr[i] = "a"
            return String(arr)
        }
        arr[arr.count-1] = "b"
        return String(arr)
    }
}
```