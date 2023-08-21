class Solution {
public:
    int minPartitions(string n) {
        int i = n.size();
        int mx = 0;
        for(int j =0; j<i; j++){
            if(mx<n[j]-'0') mx = n[j]-'0' ;
        }
        return mx;
    }
};