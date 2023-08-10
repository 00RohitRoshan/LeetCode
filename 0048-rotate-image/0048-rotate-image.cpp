class Solution {
public:
    void rotate(vector<vector<int>>& matrix) {
        int n = matrix.size();
        for(int i = 0; i<n ; i++){
            for(int j =i; j<n ; j++){
                swap(matrix[i][j],matrix[j][i]);
            }
        }
        for(int i = 0; i<n ; i++){
            int j=0;
            int k=n;
            while(j<k){
                swap(matrix[i][j],matrix[i][k-1]);
                j++;
                k--;
            }
        }
    }
};