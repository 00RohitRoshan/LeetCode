class Solution {
public:
    bool canPlaceFlowers(vector<int>& flowerbed, int n) {
        int m = flowerbed.size();
        int k = n;
        for(int i = 0; i<m && k>0; i++){
            if(i == 0 && flowerbed[i] == 0 && (i+1 == m || flowerbed[i+1] == 0)) { k--; i++; }
            else if(i == m-1 && flowerbed[i] == 0 && flowerbed[i-1] == 0) { k--; i++; }
            else if(i > 0 && i < m-1 && flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0) { k--; i++; }
        }
        return k <= 0;
    }
};
