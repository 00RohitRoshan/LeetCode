```Kotlin
class Solution {
    fun breakPalindrome(palindrome: String): String {
        val n = palindrome.length
        if (n == 1) return ""

        for (i in 0 until (n / 2)) {
            if (palindrome[i] != 'a') {
                return palindrome.replaceFirst(palindrome[i], 'a')
            }
        }

        return palindrome.substring(0, n - 1) + "b"
    }
}
```