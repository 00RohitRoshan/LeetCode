class Solution {
public:
    bool halvesAreAlike(string s) {
        int i=0, j=s.length()-1, a=0, b=0;
        
        while(j>i){
            if(s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || 
   s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U') a++;

            if(s[j] == 'a' || s[j] == 'e' || s[j] == 'i' || s[j] == 'o' || s[j] == 'u' || 
   s[j] == 'A' || s[j] == 'E' || s[j] == 'I' || s[j] == 'O' || s[j] == 'U') b++;

            
            i++;
            j--;
        }
        
        return a==b;
        
    }
};