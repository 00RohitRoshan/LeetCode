`Barclays`

Given a string S1 of length L1 consisting of Latin uppercase alphabets only and a string S2 of length L2 consisting of characters 'T' and 'F'only.

Generate a lexicographically smallest string S of length (L1+ L2 - 1) such that a substring of length 21 in string S starting at index/ (0 ≤i< L2) is equal to S1 if and only if ith element of S2 is 'T (without quotes) else not.
If no such string can be generated, print "-1" (without quotes).

Notes:
A string a is lexicographically smaller than a string b if and only if one of the following holds:
• a is a prefix of b. but a!=b.
• in the first position where a and b differ, the string a has a letter that appears earlier in the alphabet than the corresponding letter in b.

Find the Texicographically smallest string S which satisfies the given condition.
```
Input
S1 = "ABCA"
S2 = "TFTF"
Output
S = -1
```

```
Input
S1 = "ABCA"
S2 = "TFFF"
Output
S = "ABCAAAA"
```
```
Input
S1 = "ABC"
S2 = "TFTT"
Output
S = -1
```
