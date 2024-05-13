// class Solution {
//     public int matrixScore(int[][] grid) {
//         int r = grid.length;
//         int c = grid[0].length;
//         System.out.println(r+" , "+c);
        
//         for(int i = 0 ; i<r; i++){
//             if (grid[i][0] == 0){
//                 negateColumn(grid,i);
//             }
//         }
        
//         for(int i = 0 ;)
        
        
//         return 0;
        
//     }
//     public void negateColumn(int[][] grid, int col) {
//         for (int i = 0; i < grid.length; i++) {
//             grid[i][col] = -grid[i][col];
//         }
//     }
//     public int countZeros(int[][] matrix, int col) {
//         int zeroCount = 0;

//         for (int i = 0; i < matrix.length; i++) {
//             if (matrix[i][col] == 0) {
//                 zeroCount++;
//             }
//         }

//         return zeroCount;
//     }
    
// }

class Solution {
    public int matrixScore(int[][] grid) {
        int n = grid.length, m = grid[0].length;
        int res = (1 << (m - 1)) * n;  

        for (int j = 1; j < m; ++j) {
            int val = 1 << (m - 1 - j);
            int set = 0;

            for (int i = 0; i < n; ++i) {
                if (grid[i][j] == grid[i][0]) {
                    set++;
                }
            }

            res += Math.max(set, n - set) * val;
        }

        return res;
    }
}