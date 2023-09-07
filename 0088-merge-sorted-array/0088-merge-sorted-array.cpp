class Solution {
public:
    void merge(vector<int>& nums1, int n, vector<int>& nums2, int m) {
            int right=0,left=n-1;
            while(right<m && left>=0)
            {
                if(nums1[left]>nums2[right])
                {
                    swap(nums1[left],nums2[right]);
                    right++;
                    left--;
                }
                else
                break;
            }
            int ptr=n;
        for(int i=0;i<m;i++)
        {
            nums1[ptr]=nums2[i];
            ptr++;
        }  
        sort(nums1.begin(),nums1.end());
        
    }
};