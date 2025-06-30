func minBitFlips(start int, goal int) int {
    flip := start^goal;
    ans := 0;
    for flip >0 {
        if(flip & 1 == 1) {ans+=1;}
        flip >>=1;
    }
    return ans;
}