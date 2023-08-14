class Solution {
public:
    int subarraySum(vector<int>& nums, int k) 
    {
         int ans=0,sum=0;
        unordered_map<int,int> x;
        for(int i=0;i<nums.size();i++)
        {
            sum+=nums[i];
            if(sum==k) ans++; 
            ans+=x[sum-k];
            x[sum]++;
        }
        return ans;
    }
};