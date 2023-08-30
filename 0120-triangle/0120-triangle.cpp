class Solution {
public:
    int minimumTotal(vector<vector<int>>& triangle) {
        int n = triangle.size();
        vector<int> res(n,-1);
        for(int i = n-1; i>=0 ; i--){
            for(int j = 0; j< triangle[i].size(); j++){
                if(i<n-1){
                    res[j] = min(triangle[i][j]+res[j],triangle[i][j]+res[j+1]);
                }else{
                    res[j] = triangle[i][j];
                }
            }
        }
        return res[0];
    }
};