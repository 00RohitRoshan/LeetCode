class Solution {
public:
    int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {
        if(obstacleGrid[0][0]) return 0;
        for(int i =0; i<obstacleGrid.size(); i++){
            for(int j=0; j<obstacleGrid[0].size(); j++){
                if(obstacleGrid[i][j]){
                    obstacleGrid[i][j]=0;
                }else{
                    if(i==0 | j==0){
                        if(i==0 & j==0){obstacleGrid[i][j]=1;}
                        else if(i==0){
                            obstacleGrid[i][j]=obstacleGrid[i][j-1];
                        }
                        else if(j==0){
                            obstacleGrid[i][j]=obstacleGrid[i-1][j];
                        }
                    }else{
                        obstacleGrid[i][j]=obstacleGrid[i-1][j]+obstacleGrid[i][j-1];
                    }
                }
            }
        }
        return obstacleGrid[obstacleGrid.size()-1][obstacleGrid[0].size()-1];
    }
};