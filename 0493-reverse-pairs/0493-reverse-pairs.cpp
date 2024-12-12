class Solution {
public:
    int reversePairs(vector<int>& nums) {
        return ms(nums,0,nums.size()-1);
    }
    int ms(vector<int>& nums, int l, int r){
        if(l>=r) return 0;
        int m = l + (r-l)/2;
        int a = ms(nums,l,m);
        int b = ms(nums,m+1,r);
        int c = count(nums,l,m,r);
        merge(nums,l,m,r);
        return a+b+c;
    }
    int count(vector<int>& nums, int l, int m, int r) {
        int j = m + 1;
        int count = 0;
        for (int a = l; a <= m; a++) {
            while (j <= r && nums[a] > 2 * (long long)nums[j]) {
                j++;
            }
            count += (j - (m + 1));  // Fix the count calculation
        }
        return count;
    }
    void merge(vector<int>& nums, int l, int m, int r){
        int i = l;
        int j = m+1;
        vector<int> temp;
        while(i<=m && j<=r){
            if(nums[i]<=nums[j]){
                temp.push_back(nums[i]);
                i++;
            }else{
                temp.push_back(nums[j]);
                j++;
            }
        }
        while(i<=m){
            temp.push_back(nums[i]);
            i++;
        }
        while(j<=r){
            temp.push_back(nums[j]);
            j++;
        }
        for(int x =0; x<temp.size();x++){
            nums[l+x]=temp[x];
        }
    }
};