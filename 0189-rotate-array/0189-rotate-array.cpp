class Solution {
public:
    void rotate(vector<int>& nums, int k) {
        int n = nums.size();
        int r = k%n;
        reverse(nums,0,n-1-r);
        reverse(nums,n-r,n-1);
        reverse(nums,0,n-1);
        
    }
    
    private:
    void reverse(vector<int>& arr,int i,int j){
        // cout<<"hit";
        while(i<j){
            // cout<<i,j,"\n";
            // swap(arr[i] , arr[j]);
            int temp = arr[i];
            arr[i] = arr[j];
            arr[j] = temp;
            i++;
            j--;
        }
    }
};