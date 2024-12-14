class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        int l =0, r=nums.size()-1, m=0, i=-1, j=-1;
        while(l<=r){
            m = l + (r-l)/2;
            if(target<nums[m]){
                r=m-1;
            }else if(target>nums[m]){
                l=m+1;
            }else{
                i=m;
                while(i<nums.size()-1 && nums[i+1]==target)i++;
                j=m;
                while(j>0 && nums[j-1]==target)j--;
                break;
            }
        }
        return {j,i};
    }
};