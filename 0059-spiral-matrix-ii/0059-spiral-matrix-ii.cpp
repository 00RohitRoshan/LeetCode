class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
        int count = 1, i=0, j=n-1, k=0, l=n-1;
        vector<vector<int>> ans(n, vector<int>(n,0));
        while(i<=j & k<=l){
            for(int x=i; x<=j; x++){
                ans[k][x] = count++;
            }
            k++;
            for(int x=k; x<=l ; x++){
                ans[x][j]=count++;
            }
            j--;
            for(int x=j; x>=i; x--){
                ans[l][x]=count++;
            }
            l--;
            for(int x=l; x>=k; x--){
                ans[x][i]=count++;
            }
            i++;
        }
        return ans;
    }
};