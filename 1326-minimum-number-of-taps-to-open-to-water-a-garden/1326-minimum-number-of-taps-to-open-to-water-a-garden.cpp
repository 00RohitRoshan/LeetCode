class Solution {
public:
    int minTaps(int n, std::vector<int>& ranges) {
        
        vector<int> maxReach(n + 1, 0);
        //Creating the range array
        for (int i = 0; i < ranges.size(); ++i) {
            int s = std::max(0, i - ranges[i]);
            int e = i + ranges[i];
            maxReach[s] = e;
        }
        
        int tap = 0;
        int currEnd = 0;
        int nextEnd = 0;
        // Can't use while loop as it will cause infinity loop for 0 case as the range doesn't move
        for (int i = 0; i <= n; ++i) {
            if (i > nextEnd) {
                return -1;
            }
            if (i > currEnd) {
                tap++;
                currEnd = nextEnd;
            }
            nextEnd = max(nextEnd, maxReach[i]);
        }
        
        return tap;
    }
};