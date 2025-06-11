class Solution {
public:
    int beautySum(string s) {
        int a = 0;
        for(int i = 0; i<s.size(); i++){
            unordered_map<char,int> f;
            for(int j = i; j<s.size(); j++ ){
                int n = INT_MAX;
                int x = INT_MIN;
                f[s[j]]++;
                for (auto& [ch, count] : f) {
                    x = max(x, count);
                    n = min(n, count);
                }
                a += (x-n);
                cout<<f[s[j]]<<' '<<x<<' '<<n<<' '<<a<<'\n';
            }
            cout<<'\n';
        }
        return a;
    }
};