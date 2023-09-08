class Solution {
public:
    int findMin(vector<int>& nums) {
        int l1=-1,r1=nums.size()-1,l2=-1,r2=nums.size();
        while(r1-l1>2 && r2-l2>2)
        {
            if(r1-l1>2){
            int m1=l1+(r1-l1)/3;
            int m2=r1-(r1-l1)/3;
            if(nums[m1]<nums[m2])
            {
                l1=m1;
            }
            else{
                r1=m2;
            }
            }
            if(r2-l2>2){
            int m1=l2+(r2-l2)/3;
            int m2=r2-(r2-l2)/3;
            if(nums[m1]>nums[m2])
            {
                l2=m1;
            }
            else{
                r2=m2;
            }
            }
        }
        return min(nums[r1],nums[l2+1]);
    }
};