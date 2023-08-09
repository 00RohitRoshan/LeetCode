class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int leftmin = prices[0], profit = 0, n=prices.size();
        for(int i =0;i<n;i++){
            profit = max(prices[i]-leftmin,profit);
            leftmin = min(leftmin,prices[i]);
        }
        return profit;
    }
};