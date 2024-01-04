class Solution {
public:
    int minOperations(vector<int>& nums) {
        unordered_map<int,int> mp;
        for(int i = 0; i<nums.size(); i++) mp[nums[i]]++;
        int ans = 0;
        for(auto [ele,cnt] : mp){
            if(cnt == 1) return -1;
            ans += cnt/3 + (cnt%3 > 0);
        }
        return ans;
    }
};