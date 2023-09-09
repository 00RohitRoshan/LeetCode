class Solution {
public:
    // bool binary_search(vector<int> &a, int target){
    //     int l=0;
    //     int r=a.size()-1;
    //     while(l<=r){
    //         int mid=l+(r-l)/2;
    //         if(a[mid]==target){return true;}
    //         else if(a[mid]<target){l=mid+1;}
    //         else{r=mid-1;}
    //     }
    //     return false;
    // }
    bool searchMatrix(vector<vector<int>>& a, int target) {
        //Brute_Force........
        // int m = a.size();
        // int n= a[0].size();
        // bool ans = false;
        // for(int i =0; i<m; i++){
        //     for(int j=0; j<n; j++){
        //         if(a[i][j]==target){
        //             ans=true;
        //             break;
        //         }
        //     }
        // }
        // return ans;


        //Better Approach.......
        // int m=a.size();
        // int n=a[0].size();
        // bool ans= false;
        // for(int i=0; i<m; i++){
        //     if(a[i][0]<=target && a[i][n-1]>=target){
        //         ans=binary_search(a[i], target);
        //         if(ans==true){break;}
        //     }
        // }
        // return ans;

        //Optimal Aproach....
        int m=a.size();
        int n=a[0].size();
        int i=0;
        int j=n-1;
        while(i<m && j>=0){
            if(a[i][j]==target){
                return true;
            }
            else if(a[i][j]>target){j--;}
            else{i++;}
        }
        return false;
    }
};