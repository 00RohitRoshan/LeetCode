class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        
        vector<int> a(2,-1);
        int n = nums.size();
        for(int i =0; i<n; i++){
            if(nums[i] == target){
                if(a[0]  == -1) a[0] = i;
                a[1] = i;
            }
        }
        return a;
    }
};