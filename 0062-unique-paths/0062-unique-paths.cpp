class Solution {
public:
    int uniquePaths(int m, int n) {
        // vector<vector<int>> dp(m,vector<int>(n,-1));
        // return f(m-1,n-1,dp);
        
        // Tabulation
        // int dp[m][n];
        // for(int i = 0;i<m;i++){
        //     for(int j = 0;j<n;j++){
        //         if(i == 0 & j==0) {
        //             dp[i][j]=1;
        //         }
        //         else{
        //             if(i==0){
        //                 dp[i][j]=dp[i][j-1];
        //             }
        //             else if(j==0){
        //                 dp[i][j]=dp[i-1][j];
        //             }
        //             else{
        //                 dp[i][j]=dp[i-1][j]+dp[i][j-1];
        //             }
        //         }
        //     }
        // }
        // return dp[m-1][n-1];
        
        //Space opimization
        vector<int> dp(n,0);
        for(int i=0;i<m;i++){
            for(int j=0;j<n;j++){
                if(i==0 | j==0) {dp[j]=1;}
                else{
                    dp[j]=dp[j-1]+dp[j];
                }
            }
        }
        return dp[n-1];
    }
    //Recursive Function Memoized
    // int f(int i,int j, vector<vector<int>> &dp){
    //     if(i == 0 | j == 0){
    //         if(i==j) {return 1;}
    //     } 
    //     if (i<0 | j<0) return 0;
    //     if(dp[i][j] != -1) return dp[i][j];
    //     int up = f(i,j-1,dp);
    //     int left = f(i-1,j,dp);
    //     return dp[i][j] = up+left;
    // }
};

//Today 3rd Sep 2023 is the first time when a previously solved problem appears on leetcode daily.