class Solution {
public:
    int buyChoco(vector<int>& p, int money) {
    // Use the nth_element function to partially sort the vector such that 
    // the first two elements are the smallest
        nth_element(begin(p), begin(p) + 1, end(p));

    // Check if the sum of the first two elements is less than or equal to the money
    // If true, return the remaining money after buying the chocolates
    // If false, return the original amount of money (since we can't afford any chocolates)
        return money < p[0] + p[1] ? money : money - p[0] - p[1];
    }
};