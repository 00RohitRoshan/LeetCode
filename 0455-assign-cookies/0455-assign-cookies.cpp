class Solution {
public:
    int findContentChildren(vector<int>& g, vector<int>& s) {
        sort(g.begin(),g.end());
        sort(s.begin(),s.end());
        
        int ans = 0, i = 0;
        for(int j = 0; j<s.size(); j++){
            if(i<g.size() && g[i] <= s[j]){
                ans++;
                i++;
            }
        }
        return ans;
    }
};