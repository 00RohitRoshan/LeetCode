class Solution {
public:
    bool search(vector<int>& nums, int target) {
        int i = 0, j = nums.size()-1, m=0;
        while(i<=j){
            m = i + (j-i)/2;
            if(nums[m] == target) return true;
            if (nums[i] == nums[m] && nums[m] == nums[j]) {
                i = i + 1;
                j = j - 1;
                continue;
            }
            if(nums[i]<=nums[m]){
                if(target >= nums[i]  && target <= nums[m]){
                    j = m-1;
                }else{
                    i = m+1;
                }
            }else{
                if(target >= nums[m]  && target <= nums[j]){
                    i = m+1;
                }else{
                    j = m-1;
                }
            }
        }
        return false;
    }
};