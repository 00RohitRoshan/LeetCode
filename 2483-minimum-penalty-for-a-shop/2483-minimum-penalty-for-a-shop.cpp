class Solution {
public:
    int bestClosingTime(string customers) {
        int n = customers.size();
        int penalty = 0;// Closes at 0-1
        for(int i =0; i<n ; i++){
            if(customers[i] == 'Y') penalty++;
        }
        int minPenalty = penalty;
        int minHour = 0;
        for(int i =0; i<n ; i++){
            if(customers[i] == 'Y') --penalty;
            if(customers[i] == 'N') ++penalty;
            if(penalty < minPenalty){
                minPenalty = penalty;
                minHour = i+1;
            }
        }
        return minHour;
    }
};