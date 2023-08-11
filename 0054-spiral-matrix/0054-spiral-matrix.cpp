class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        int n = matrix.size();
        int m = matrix[0].size();
        vector<int> ans;
        int i=0 , j=m-1 , k=0 , l=n-1 ;
        while(i<=j & k<=l){
            for(int x=i;x<=j;x++){
                ans.push_back(matrix[k][x]);
            }
            k++;
            for(int x=k;x<=l;x++){
                ans.push_back(matrix[x][j]);
            }
            j--;
            if(k<=l){
                for(int x=j;x>=i;x--){
                    ans.push_back(matrix[l][x]);
                }
                l--;
            }
            if(i<=j){
                for(int x=l;x>=k;x--){
                    ans.push_back(matrix[x][i]);
                }
                i++;
            }
        }
        return ans;
    }
};