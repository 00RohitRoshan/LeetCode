class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        int l = 0;
        int r = nums.size()-1;
        int m = 0;
        while(l<=r){
            m = l+(r-l)/2;
            if(target<nums[m]){
                r=m-1;
            }else if(target == nums[m]){
                return m;
            }else{
                l=m+1;
            }
        }
        return l;
    }
};