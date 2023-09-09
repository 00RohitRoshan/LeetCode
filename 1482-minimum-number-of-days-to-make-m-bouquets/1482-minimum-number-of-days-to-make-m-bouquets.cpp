class Solution {

bool canMake(vector<int>&v, int mid, int m, int k ){ 
    long long bc=0, adcnt=0;
    for(int i=0; i<v.size(); i++){
        if(v[i]<=mid){
            adcnt++;
            if(adcnt==k){
                bc++;
                if(bc==m)return true;
                adcnt=0;
            }
        }
        else adcnt=0;
    } 
    return false;
}

public:
    int minDays(vector<int>& bloomDay, int m, int k) {
        int s = 1, e = *max_element(bloomDay.begin(),bloomDay.end());
        int ans=-1;
        while(s<=e){
            int mid = s+(e-s)/2;
            if(canMake(bloomDay, mid, m, k)){
                ans = mid;
                e = mid-1;
            }
            else s = mid+1;
        }
        return ans; 
    }
};