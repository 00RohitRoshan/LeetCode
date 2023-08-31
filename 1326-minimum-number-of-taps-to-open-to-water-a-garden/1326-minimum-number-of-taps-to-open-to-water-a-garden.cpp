class Solution {
public:
    int minTaps(int n, vector<int>& ranges) 
    {
        int taps = 0;
        int maxi = 0, mini = 0; // initially both are zero
        int lastTap = 0;

        //we don't want to calculate when our max range is past the range
        while(maxi < n)
        {
            for(int i = lastTap; i <= n; i++)
            {
                //majority of the time for loop is not doing much
                //we are just straching the range as much as possible
                //when the below criteria satisfies
                if(i - ranges[i] <= mini && i + ranges[i] > maxi)
                {
                    lastTap = i;
                    maxi = i + ranges[i];
                }
            }
            // if garden cannot be watered
            if(maxi == mini)
                return -1;
            //Before the while loop exits the tap gets incrimented
            //And mini (min range) gets updated
            taps++; 
            mini = maxi;
        }
        return taps;
    }
};