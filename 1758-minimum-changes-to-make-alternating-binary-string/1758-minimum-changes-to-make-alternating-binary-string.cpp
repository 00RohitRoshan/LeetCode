class Solution {
public:
    int minOperations(string s) {
        
        int n = s.length();
        int res = 0;
        for(int i=0;i<n;i++){
            if(i%2 != s[i]-'0'){
                res++;
            }
        }
        return min(res,n-res);
    }
};