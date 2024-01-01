class Solution {
public:
    bool isVowel(char c) {
        c = tolower(c);
        return (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u');
    }

    std::string reverseVowels(std::string s) {
        int i = 0, j = s.size() - 1;
        while(i < j) {
            if(!isVowel(s[i])) {
                ++i;
                continue;
            }
            if(!isVowel(s[j])) {
                --j;
                continue;
            }
            std::swap(s[i++], s[j--]);
        }
        return s;
    }
};