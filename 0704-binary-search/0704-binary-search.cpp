class Solution {
public:
    int search(vector<int>& nums, int target) {
        int i = 0;
        int j = nums.size()-1;
        int mid = 0;
        // if(nums[i]==target){
        //     return i;
        // }else
        // if(nums[j]==target){
        //     return j;
        // }
        while(i<=j){
            // if(nums[i]==target){
            //     return i;
            // }else
            // if(nums[j]==target){
            //     return j;
            // }
            mid = i+ (j-i)/2;
            if(nums[mid]<target){
                i=mid+1;
            }else
            if(nums[mid]>target){
                j=mid-1;
            }else
            if(nums[mid]==target){
                return mid;
            }
        }
        return -1;

    }
};