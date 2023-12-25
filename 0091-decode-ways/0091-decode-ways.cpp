// Memoized
// class Solution {
// public:
//     int numDecodings(string s) {
//         vector<int> dp(s.size()+1,-1);
//         return s.empty() ? 0 : numDecodings(0,s,dp); 
//     }
//     int numDecodings(int p, string &s, vector<int> &dp){
//         if (dp[p] > -1) return dp[p];
//         if (s[p] == '0') return dp[p] = 0;
//         int n = s.size();
//         if (p == n) return dp[p] = 1;
//         int res = numDecodings(p+1,s,dp);
//         if (p<n-1 && (s[p] == '1'|| (s[p] == '2' && s[p+1] < '7'))) res += numDecodings(p+2,s,dp);
//         return dp[p] = res;
//     }
// };

// Bottom Up
class Solution {
public:
    int  numDecodings(string s) {
        if (s.empty()) return 0;
        int n = s.size();
        vector<int> dp(n+1,0);
        dp[n] = 1;
        for(int i = n-1; i >=0;i--){
            if (s[i]!='0') dp[i] = dp[i+1];
            if (i < n-1 && (s[i] == '1' || (s[i] == '2' && s[i+1] < '7'))) dp[i]=dp[i]+dp[i+2];
        }
        return dp[0];
    }
};
