We can solve this problem by building two strings and comparing them.
This solution has a time complexity of O(N) where N is the total length of strings s and t. It also has a space complexity of O(N) for storing the final transformed strings.
```
The idea behind this optimized solution is to compare the strings s and t in reverse order, i.e., from the end to the beginning. This is because the backspace operation affects the characters that come after it, so it’s easier to start from the end.
​
We use two pointers, i and j, to iterate through s and t respectively. At each step, we first handle the backspace characters:
​
If the current character is a backspace (i.e., ‘#’), we increment a counter (backspace_s or backspace_t) and move the pointer one step to the left.
If the current character is not a backspace, but there are some backspaces to apply (i.e., the counter is greater than 0), we decrement the counter and move the pointer one step to the left.
We repeat these steps until we find a character that is not a backspace and there are no more backspaces to apply.
Once we have done this for both s and t, we compare the characters at pointers i and j. If they are not equal, we return False, because that means s and t are different. If they are equal, we move both pointers one step to the left and repeat the process.
​
If we have iterated through both strings without returning False, that means all non-deleted characters in s and t are equal, so we return True.
​
This solution is more efficient than building new strings because it only requires constant extra space to store the counters and pointers, regardless of how long s and t are. It still requires linear time because in the worst case we need to iterate through all characters in s and t. However, it’s faster in practice because we can stop as soon as we find a pair of unequal characters.
```
​