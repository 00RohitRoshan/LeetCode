class Solution {
public:
    bool backspaceCompare(string s, string t) {
        int i = s.length() - 1, j = t.length() - 1;
        int backspace_s = 0, backspace_t = 0;

        while (true) {
            while (i >= 0 && (backspace_s || s[i] == '#')) {
                backspace_s += s[i] == '#' ? 1 : -1;
                i--;
            }
            while (j >= 0 && (backspace_t || t[j] == '#')) {
                backspace_t += t[j] == '#' ? 1 : -1;
                j--;
            }
            if (!(i >= 0 && j >= 0 && s[i] == t[j])) {
                return i == -1 && j == -1;
            }
            i--; j--;
        }
    }
};