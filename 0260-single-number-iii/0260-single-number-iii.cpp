class Solution {
public:
    vector<int> singleNumber(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        for(int i = 0; i < nums.size()-1; i++){      //255677
            if(nums[i]==nums[i+1]){
                nums.erase(nums.begin()+i,nums.begin()+i+2);
                i--;
            }
        }
        return nums;
    }
};