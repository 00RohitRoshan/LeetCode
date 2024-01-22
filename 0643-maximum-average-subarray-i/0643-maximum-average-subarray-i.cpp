class Solution {
public:
    double findMaxAverage(vector<int>& nums, int k) {
        if (nums.size() == 1) {
            return static_cast<double>(nums[0]);
        }
        
        int start = 0;
        int end = k;
        double average = 0.0;
        
        // Calculate the average of the first window
        for (int i = 0; i < k; i++) {
            average += static_cast<double>(nums[i]) / k;
        }
        
        // Initialize the maximum average
        double max_average = average;
        
        // Slide the window and update the maximum average
        while (end < nums.size()) {
            average = average - static_cast<double>(nums[start]) / k;
            average = average + static_cast<double>(nums[end]) / k;
            max_average = max(max_average, average);
            start++;
            end++;
        }
        
        return max_average;
    }
};
