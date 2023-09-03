class Solution {
public:
    int minExtraChar(string s, vector<string>& dictionary) {
        unordered_set<string> dictionary_set(dictionary.begin(), dictionary.end());
        vector<int> dp(s.length() + 1, -1);
        return minExtraCharHelper(s, 0, dictionary_set, dp);
    }
    
    int minExtraCharHelper(string s, int start, unordered_set<string>& dictionary_set, vector<int>& dp) {
        if (start == s.length()) {
            return 0;
        }
        if (dp[start] != -1) {
            return dp[start];
        }
        int max_val = s.length() + 1;
        dp[start] = max_val;
        for (int i = start + 1; i <= s.length(); ++i) {
            if (dictionary_set.find(s.substr(start, i - start)) != dictionary_set.end()) {
                dp[start] = min(dp[start], minExtraCharHelper(s, i, dictionary_set, dp));
            }
        }
        dp[start] = min(dp[start], minExtraCharHelper(s, start + 1, dictionary_set, dp) + 1);
        return dp[start];
    }
};
