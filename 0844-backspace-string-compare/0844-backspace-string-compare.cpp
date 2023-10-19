class Solution {
public:
    string build(string str) {
        string result;
        for (char c : str) {
            if (c != '#')
                result.push_back(c);
            else if (!result.empty())
                result.pop_back();
        }
        return result;
    }

    bool backspaceCompare(string s, string t) {
        return build(s) == build(t);
    }
};