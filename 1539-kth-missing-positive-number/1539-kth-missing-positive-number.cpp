class Solution {
public:
    int findKthPositive(vector<int>& arr, int k) {
        
        if(arr.size() > arr.back() || arr.front() > k){
            return arr.front() > k ? k : k + arr.size();
        }
        int h = arr.size()-1, l = 0;
        while(l <= h){
            int m = (l + h)>>1;
            if(arr[m]-m-1 < k){
                l = m+1;
            }else{
                h = m-1;
            }
        }
        return l + k;
    }
};