class Solution {
private:
    bool checkAllCharsSame(string &s, int &stringLen){
       
        for(int i = 1; i < stringLen; i++)
        {
            if(s[i] != s[i-1])
                return false;
        }
        
        // all characters are same
        return true;
    }
    
int findMinLenOfOptimalCompression(int idx, int &stringLen, string &s, int k, char prev, int freqCount, int dp[][27][11][101]){
    
    // return max value, if k is less than zero as we cannot delete char further
    if(k < 0) return 1e9;
    
    // we reached end of string
    if(idx >= stringLen)
        return 0;
    
    // consider freqCount as 10, because, until 10 .. 99, string length will be 2,
    if(freqCount >= 10)
        freqCount = 10;
    
    if(dp[idx][prev - 'a'][freqCount][k] != -1)
        return dp[idx][prev - 'a'][freqCount][k];
    
    int minLen = INT_MAX;
    // possibilites:
    // case 1: if k > 0, we can delete cur char and move to next index
    minLen = min(minLen, findMinLenOfOptimalCompression(idx + 1, stringLen, s, k-1, prev, freqCount, dp));
    
    // case 2 : if our cur char is not same as prev, then our len will be increased and freqCount starts with 1
    if(s[idx] != prev)
    minLen = min(minLen, 1 + findMinLenOfOptimalCompression(idx + 1, stringLen, s, k, s[idx], 1, dp));
    
    else{
    // case 3: If our char matches with prev,
    // case 3.1: if freqcount is 1/9/99, then our freqcount is going to increase + 1, res will also be increased by 1
        if(freqCount == 1 || freqCount == 9 || freqCount == 99)
            minLen = min(minLen, 1 + findMinLenOfOptimalCompression(idx+1, stringLen, s, k, prev, freqCount+1, dp));
        
    // case 3.2: we move to next index with freqcount + 1
        else{
            minLen = min(minLen, findMinLenOfOptimalCompression(idx+1, stringLen, s, k, prev, freqCount+1, dp));
        }
    }
    
    
    return dp[idx][prev - 'a'][freqCount][k] = minLen;
}
    
public:
    int getLengthOfOptimalCompression(string s, int k) {
        
        int stringLen = s.size();
        
        if(stringLen <= k)
            return 0;
        
        bool allCharsSame = checkAllCharsSame(s, stringLen);
        
        if(allCharsSame){
            
            // remove k characters from s
            stringLen -= k;
            
            // Given Statement: Notice that in this problem, we are not adding '1' after single characters.
            if(stringLen == 1)
                return 1;
            
            int numOfDigits = log10(stringLen) + 1;
            
            // + 1 for cur Char
            return numOfDigits + 1;
        }
        
        // need to take care of 4 states
        // => idx : 0 ... stringLen
        // => prev : ranges from 'a' to 'z'
        // => freqCount : count of cur char 
        // => k : we need to try possibilities for k -> 0 to 100
        int dp[stringLen + 1][27][11][101];
        memset(dp, -1, sizeof(dp));
        
        // If we are reaching this statement, then all characters are not same, so try all possibility
        return findMinLenOfOptimalCompression(0, stringLen, s, k, 'z' + 1, 0, dp);
    }
};