class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int n= nums.size() , i =0 , j =0 ;
        // int i,j = 0,0;
        
        while( i < n){
            while (j < n-1){
                if (nums[j] == nums[j+1]){
                    j++;
                }else{
                    nums[i+1] = nums[j+1];
                    i++;
                    j++;
                }
            }
            break;
            
        }
         return i+1;
        
    }
};