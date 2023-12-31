class Solution {
public:
    vector<bool> kidsWithCandies(vector<int>& candies, int extraCandies) {  
        int n = candies.size();
        vector<bool> res(n,0);
        int mx = 0;
        for(int i = 0; i<n; i++){
            mx = max(mx,candies[i]);
        }
        for(int i = 0; i<n; i++){
            res[i] = (extraCandies+candies[i] >= mx);
        }
        return res;
    }
};