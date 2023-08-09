class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int n= nums.size() , i =0 , j =1 ;
        // int i,j = 0,0;
        
        while( i < n){
            while (j < n){
                if (nums[i] != nums[j]){
                    nums[i+1] = nums[j];
                    i++;
                    j++;
                }else{
                    j++;
                }
            }
            break;
            
        }
         return i+1;
        
    }
};