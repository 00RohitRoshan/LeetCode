class Solution {
public:
    int projectionArea(vector<vector<int>>& grid) {
        int n = grid.size();
        int area = 0;
        int la = 0;
        int lb = 0;
        for(int i =0;i<n;i++){
            int maxl = 0;
            int maxb = 0;
            for(int j = 0;j<n;j++){
                maxl = max(maxl,grid[i][j]);
                maxb = max(maxb,grid[j][i]);
                if(grid[i][j] != 0){
                    area++;
                }
            }
            la+=maxl;
            lb+=maxb;
        }
        return area+la+lb ;
    }
};