class Solution {
public:
    int romanToInt(string s) {
        map<char, int> m = {
            {'M', 1000},
            {'D', 500},
            {'C', 100},
            {'L', 50},
            {'X', 10},
            {'V', 5},
            {'I', 1}
        };
        int p = INT_MIN;
        int ans = 0;
        for(int i = s.size(); i>=0 ; i--){
            if(m[s[i]] >= p){
                ans += m[s[i]];
            }else{
                ans -= m[s[i]];
            }
            p=m[s[i]];
        }
        return ans;
    }
};