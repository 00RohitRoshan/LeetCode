class Solution {
public:
    int minSteps(string s, string t) {
        unordered_map<char, int> mp ;
        
        for(int i=0; i<s.size(); i++){
            mp[s[i]]++;
            mp[t[i]]--;
        }
        
        int ans = 0;
        for(auto p : mp){
            if(p.second > 0) ans+=p.second ;
        }
        
        return ans;
    }
};