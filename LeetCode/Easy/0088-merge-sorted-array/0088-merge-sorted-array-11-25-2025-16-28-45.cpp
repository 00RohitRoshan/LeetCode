class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        int i =nums1.size()-1;
        if(m == 0){
            nums1 = nums2;
        }
        while(n>0){
            nums1[i]=nums2[n-1];
            n--;
            i--;
        }
        sort(nums1.begin(),nums1.end());
        // while(m>0 && n>0 && i>=0){
        //     if(nums1[m-1]>=nums2[n-1]){
        //         nums1[i]=nums1[m-1];
        //         nums1[m-1]=0;
        //         m--;
        //         i--;
        //     }
        //     else if(nums1[m-1]<nums2[n-1]){
        //         nums1[i]=nums2[n-1];
        //         n--;
        //         i--;
        //     }
        //     // if(nums1[m]==nums1[n]){
        //     //     nums1[i]=nums1[m];
        //     //     m--;
        //     //     i--;
        //     // }
        // }
    }
};