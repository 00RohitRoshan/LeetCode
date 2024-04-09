class Solution {
    public int timeRequiredToBuy(int[] tickets, int k) {
        int ans = 0;
        
           while(true){
                for (int i = 0; i < tickets.length; i++) {
                if(tickets[k]>0){
                    if (tickets[i] - 1 >= 0) {
                        ans++;
                        tickets[i] = tickets[i] - 1;
                    }
                }
                else{
                    return ans;
                }
            }
           }
        
        // return ans;
    }
}
