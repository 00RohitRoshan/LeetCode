class Solution {
public:


    bool predicate(int min_cap,vector<int>& weights,int days){
        int day=1;
        int total=0;
        for(auto weight:weights){
            total+=weight;
            if(total>min_cap){
                total=weight;
                day++;
                if(day>days){
                    return false;
                }
            }
        }
        return true;
    }
    int shipWithinDays(vector<int>& weights, int days) {
        int l=*max_element(weights.begin(), weights.end());
        int r=accumulate(weights.begin(), weights.end(), 0);
        while(l<r){
            int m= l + (r-l)/2;
            if(predicate(m, weights,days)){
                r=m;
            }
            else{
                l=m+1;
            }
        }
        return l;
    }
};