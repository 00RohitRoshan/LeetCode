class Solution {
public:
    int maxOperations(vector<int>& nums, int k) {
        unordered_map<int, int> mp;
        int res = 0;
        for (auto num : nums) {
            if (mp[k - num] > 0) {
                res++;
                mp[k - num]--;
            } else {
                mp[num]++;
            }
        }
        return res;
    }
};
