class Solution {
public:
    int subarraySum(vector<int>& nums, int k) 
    {
         int ans=0;
        unordered_map<int,int> x;
        
        int sum=0;
      
        for(int i=0;i<nums.size();i++)
        {
            sum+=nums[i];
            if(sum==k) ans++;

            if(x.find(sum-k)!=x.end()) ans+=x[sum-k];

            x[sum]++;
        }

        return ans;
    }
};