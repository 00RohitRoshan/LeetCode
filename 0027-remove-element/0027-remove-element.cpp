class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int n = nums.size();
        int i=0,j=0;
        while(j<n){
            if(nums[i]==val){
                if(nums[j]!=val){
                    swap(nums[i],nums[j]);
                    i++;
                    j++;
                }else{j++;}
            }else{
                i++;
                j++;
            }
        }
        return i;
    }
};