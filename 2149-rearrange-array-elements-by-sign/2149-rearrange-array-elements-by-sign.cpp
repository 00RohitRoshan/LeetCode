class Solution {
public:
    vector<int> rearrangeArray(vector<int>& nums) {
        int n=nums.size(),pos=0,neg=1,k=0;
        vector<int> ans(n,0);
        while(k<n){
            if(nums[k]>0){
                ans[pos]=nums[k];
                pos+=2;
            }
            if(nums[k]<0){
                ans[neg]=nums[k];
                neg+=2;
            }
            k++;
        }
        return ans;
    }
};