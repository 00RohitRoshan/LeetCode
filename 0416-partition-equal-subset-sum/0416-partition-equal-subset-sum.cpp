class Solution {
public:
    bool canPartition(vector<int>& nums) {
        int tosum = 0;
        for(int i=0; i<nums.size(); i++){
            tosum += nums[i];
        }
        if(tosum%2) return false;
        int target = tosum/2;
        vector<vector<int>> dp(nums.size(),vector<int>(target+1,-1));
        return f(target,nums.size()-1,nums,dp);
    }
    bool f(int target,int index,vector<int>& arr,vector<vector<int>> &dp){
        if(target == 0) return true;
        if(index == 0) return (target == arr[0]);
        if(dp[index][target] != -1) return dp[index][target];
        
        bool notTake = f(target,index-1,arr,dp);
        bool take = false;
        if(target > arr[index]) take = f(target-arr[index],index-1,arr,dp);
        return (dp[index][target]=take|notTake) ;
    }
};