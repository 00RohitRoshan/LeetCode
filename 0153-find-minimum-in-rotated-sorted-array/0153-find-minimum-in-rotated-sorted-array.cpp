class Solution {
public:
    int findMin(vector<int>& nums) {
        int i = 0, j = nums.size()-1, m=0, mini=INT_MAX;
        while(i<=j){
            m = i + (j-i)/2;
            if(nums[m] < mini) mini = min(mini,nums[m]);
            if (nums[i] == nums[m] && nums[m] == nums[j]) {
                i = i + 1;
                j = j - 1;
                continue;
            }
            if(nums[i]<=nums[m]){
                mini = min(mini,nums[i]);
                i = m+1;
            }else{
                mini = min(mini,nums[m]);
                j = m-1;
            }
        }
        return mini;
    }
};