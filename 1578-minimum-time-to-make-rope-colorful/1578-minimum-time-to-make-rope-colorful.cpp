class Solution {
public:
    int minCost(string colors, vector<int>& cost) {
        int n = colors.length();
        int sum = 0, maxCost = 0, ans = 0;
        for(int i = 0; i < n; i++){
            //count for a color
            if(i > 0 && colors[i] != colors[i - 1]){
                ans += sum - maxCost;
                sum = 0;
                maxCost = 0;
            }
            sum += cost[i];
            maxCost = max(cost[i], maxCost);
        }
        ans += sum - maxCost;
        return ans;
    }
};