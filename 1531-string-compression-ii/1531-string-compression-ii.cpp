int dp[101][27][101][101];
class Solution {
public:
    unordered_set<int> possibleLen = {1, 9, 99};
    int solve(string& s, int idx, char lastChar, int lastCharCount, int k) {
        if (k < 0) return INT_MAX / 4;
        if (idx == s.length()) return 0;

        int& ans = dp[idx][lastChar - 'a'][lastCharCount][k];
        if(ans != -1) return ans; 

        int keepChar;
        int deleteChar = solve(s, idx + 1, lastChar, lastCharCount, k - 1);

        if (s[idx] == lastChar) {
            keepChar = solve(s, idx + 1, lastChar, lastCharCount + 1, k) + (possibleLen.count(lastCharCount) ? 1 : 0);
        } else {
            keepChar = solve(s, idx + 1, s[idx], 1, k) + 1;
        }

        ans = min(keepChar, deleteChar);
        return ans;
    }

    int getLengthOfOptimalCompression(string s, int k) {
        memset(dp,-1,sizeof(dp));
        return solve(s, 0, 'z' + 1, 0, k);
    }
};