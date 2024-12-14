class Solution {
public:
    int search(vector<int>& nums, int k) {
        int l = 0, r = nums.size() - 1;
        while (l <= r) {
            int m = l + (r - l) / 2;
            
            // Check if we found the element
            if (nums[m] == k) return m;

            // Handle duplicates by incrementing l or decrementing r
            if (nums[l] == nums[m] && nums[m] == nums[r]) {
                l++;  // Skip duplicate elements on both ends
                r--;  // Skip duplicate elements on both ends
            } else if (nums[l] <= nums[m]) {  // Left part is sorted
                if (k >= nums[l] && k < nums[m]) {
                    r = m - 1;  // Search in the left half
                } else {
                    l = m + 1;  // Search in the right half
                }
            } else {  // Right part is sorted
                if (k > nums[m] && k <= nums[r]) {
                    l = m + 1;  // Search in the right half
                } else {
                    r = m - 1;  // Search in the left half
                }
            }
        }
        return -1;  // Element not found
    }
};
