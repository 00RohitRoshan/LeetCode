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
// class Solution {
// public:
//      int numDecodings(string s) {
//         int n = s.size();
//         vector<int> dp(n+1);
//         dp[n] = 1;
//         for(int i=n-1;i>=0;i--) {
//             if(s[i]=='0') dp[i]=0;
//             else {
//                 dp[i] = dp[i+1];
//                 if(i<n-1 && (s[i]=='1'||s[i]=='2'&&s[i+1]<'7')) dp[i]+=dp[i+2];
//             }
//         }
//         return s.empty()? 0 : dp[0];   
//      }
// };

// Space optimization
class Solution {
public:
    int numDecodings(string s) {
        // p = previous
        int p = 1, pp, n = s.size();
        for(int i=n-1;i>=0;i--) {
            int cur = s[i]=='0' ? 0 : p;
            if(i<n-1 && (s[i]=='1'||s[i]=='2'&&s[i+1]<'7')) cur+=pp;
            pp = p;
            p = cur;
        }
        return s.empty()? 0 : p;   
    }
};
