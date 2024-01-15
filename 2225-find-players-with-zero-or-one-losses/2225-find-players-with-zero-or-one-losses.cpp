class Solution {
public:
    vector<vector<int>> findWinners(vector<vector<int>>& matches) {
        
        map<int, int> mp;
        for(int i = 0; i<matches.size(); i++){
            if(mp.find(matches[i][0]) == mp.end()) mp[matches[i][0]] = 0;
            mp[matches[i][1]]++;
        }
        
        vector<int> zeroLost, oneLost ;
        for(auto count : mp){
            if(count.second == 0) zeroLost.push_back(count.first);
            if(count.second == 1) oneLost.push_back(count.first);
        }
        
        return {zeroLost,oneLost};
        
    }
};