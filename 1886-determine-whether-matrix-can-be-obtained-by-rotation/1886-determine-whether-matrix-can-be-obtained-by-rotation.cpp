class Solution {
public:
      void rotate(vector<vector<int>> &matrix){
                  vector<vector<int>> a=matrix;

            int n=matrix.size();
            int s=0;
            for(int i=0;i<n;i++){
                for(int j=0;j<n;j++){
                matrix[i][j]=a[j][s];
                     }
                 s++;
                    }
                for(int i=0;i<n;i++){
                    reverse(matrix[i].begin(),matrix[i].end());

                }
      }
    bool findRotation(vector<vector<int>>& mat, vector<vector<int>>& target) {
       for(int i=0;i<4;i++){
           rotate(mat);
           if(mat==target) return true;
       }
   return false;
    }
};