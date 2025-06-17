class Solution {
public:
    bool isPowerOfTwo(int n) {
        if (n <= 0) return false;
        int i = 1;
        while (i < n) {
            i <<= 1; // shift and assign
        }
        return i == n;
    }
};