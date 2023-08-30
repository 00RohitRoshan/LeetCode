class Solution {
public:
    int hIndex(vector<int>& citations) {
        int n = citations.size();
        int l = 0 , h = n-1;
        int ans = 0 ;
        while( l <= h ){
            int mid = l + (h-l)/2;
            int idx = n - mid ;
            if(citations[mid] >= idx){
                ans = idx ;
                h = mid-1;
            }
            else{
                l = mid+1;
            }
        }
        return ans;
    }
};