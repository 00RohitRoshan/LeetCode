class Solution {
public:
    bool makeEqual(vector<string>& words) {
        map<char, int> mp;
        int n = words.size();
        for(int i = 0; i<n ; i++){
            string str = words[i];
            int m = str.size();
            for(int j = 0; j<m; j++){
                mp[str[j]]++;
            }
        }
        for (auto& pair : mp) {
            if (pair.second % n != 0) {
                return false;
            }
        }
        return true;
    }
};