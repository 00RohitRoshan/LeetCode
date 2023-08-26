class Solution {
public:
    int myAtoi(string s) {
        int i =0 , n = s.size();

        // skipping space characters at the beginning
        while(s[i] == ' ') i++;

        int positive = 0 , negative = 0;

        // number of positive signs at the start in string
        if(s[i] == '+') {
            positive++;
            i++;
        }

        // number of negative signs at the start in string
        if(s[i] == '-'){
            negative++;
            i++;
        }

        double ans =0;

        while(i<n & s[i] >= '0' & s[i] <= '9'){
            ans = ans*10 + (s[i] - '0');// (s[i] - '0') is converting character to integer
            i++;
        }

        if(negative>0) ans = -ans;
        if(negative>0 & positive > 0) return 0 ;

        if(ans > INT_MAX) return INT_MAX;
        if(ans < INT_MIN) return INT_MIN;
        return (int)ans;
    }
};