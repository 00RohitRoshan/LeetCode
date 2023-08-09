class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int n = nums.size();
        int x = 0;
        for (int i =0; i<n ; i++){
            // cout<<x;
            x = x^i;
            // cout<<x;
            x = x^nums[i];
            // cout<<x;
        }
        return x^n;
    }
    
};