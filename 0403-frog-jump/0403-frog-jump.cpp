class Solution {
public:
    int n; // Number of stones
    vector<vector<int>> dp; //2D dp matrix (index for stone, jump units)

    // Recursive function to check if frog can cross the river using given stones
    // i: current stone index, k: jump units
    bool f(vector<int>& stones, int i, int k) {
        if (i == n - 1) return 1; // If frog reaches the last stone, return true
        if (i >= n) return 0;    // If frog goes beyond the last stone, return false
        if (dp[i][k] !=-1) return dp[i][k]; // Return stored result if available

        bool ans=0;
        for (int jump : {k - 1, k, k + 1}) { //Possible jump units: k-1, k, k+1
            if (jump==0) continue; // Skip 0 jump
            // Use binary search, as stones[i] is in ascending order
            int next = lower_bound(stones.begin()+(i+1), stones.end(), stones[i]+jump)- stones.begin();
            if (next==n ||stones[next]!=stones[i] + jump) continue; // Jump not found
            ans = ans || f(stones, next, jump); // Recursively check the next jump
        }
        return dp[i][k] = ans; // Store and return the result
    }

    bool canCross(vector<int>& stones) {
        n = stones.size(); // Set the number of stones
        dp.assign(n + 1, vector<int>(n + 1, -1)); // Initialize the 2D dp matrix with -1
        return f(stones, 0, 0); // Start the recursive checking from the first stone
    }
};