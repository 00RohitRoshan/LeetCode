class Solution {
public:
    int minFallingPathSum(vector<vector<int>>& matrix) {
        int n = matrix.size();
        vector<int> cur = matrix[0], prev = cur;

        for (int i = 1; i < n; ++i) {
            for (int j = 0; j < n; ++j) {
                // Handle the edge case for the first and last column
                if (j == 0) {
                    cur[j] = matrix[i][j] + min(prev[j], prev[j + 1]);
                } else if (j == n - 1) {
                    cur[j] = matrix[i][j] + min(prev[j], prev[j - 1]);
                } else {
                    cur[j] = matrix[i][j] + min({prev[j - 1], prev[j], prev[j + 1]});
                }
            }
            swap(cur, prev); // The current row becomes the previous row in the next iteration
        }

        // The minimum element in `prev` is the minimum sum of a falling path
        return *min_element(prev.begin(), prev.end());
    }
};
