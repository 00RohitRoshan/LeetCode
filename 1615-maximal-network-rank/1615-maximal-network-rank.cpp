class Solution {
public:
    int maximalNetworkRank(int n, vector<vector<int>>& roads) {
        int count[101] = {0};
        int direct[101][101] = {0};
        for(auto road: roads){
            count[road[0]]++;
            count[road[1]]++;
            direct[road[0]][road[1]] = 1;
            direct[road[1]][road[0]] = 1;
        }
        
        int rank =0;
        for(int i =0; i<n; i++){
            for(int j = i+1; j<n; j++){
                rank = max(rank, count[i]+count[j]-direct[i][j]);
            }
        }
        return rank;
    }
};