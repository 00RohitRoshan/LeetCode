class Solution {
public:
    int minAddToMakeValid(string s) {
        int l=0,r=0;
        for(const char c:s){
            if(c == '(') l++;
            if(c == ')'){
                if(l == 0){
                    r++; 
                }
                else{ l-- ;}
            }
        }
        
        return l+r;
    }
};