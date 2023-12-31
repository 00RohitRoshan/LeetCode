class Solution {
public:
    int maxLengthBetweenEqualCharacters(string s) {
        map<char, int> mp;
        int n = s.size();
        int mx = -1;
        for(int i = 0; i<n; i++){
            if (mp.find(s[i]) != mp.end()){
                mx = max(mx, (i-mp[s[i]]-1));
            }else{
                mp[s[i]] = i;
            }
        }
        return mx;
    }
};