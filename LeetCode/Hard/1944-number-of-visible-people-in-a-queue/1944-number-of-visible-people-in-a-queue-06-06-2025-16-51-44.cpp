class Solution {
public:
    vector<int> canSeePersonsCount(vector<int>& heights) {
        // vector<int> ans;
        // for(int i = 0; i<heights.size()-1; i++){
        //     int h = heights[i+1];
        //     int h2 = heights[i+1];
        //     int c = 1;
        //     for(int j = i+2; j<heights.size(); j++){
        //         if(
        //             // heights[j]>heights[j-1] && 
        //             heights[j] > h
        //             && h <= heights[i] ){
        //             c++;
        //             h2 = h;
        //             h = heights[j];
        //         }
        //     }
        //     ans.push_back(c);
        // }
        // ans.push_back(0);
        // return ans;

        vector<int> ans(heights.size(),0);
        stack<int> st;
        for(int i = heights.size()-1; i>=0; i--){
            while (!st.empty() && heights[i]>st.top()){
                st.pop();
                ans[i]++;
            }
            if(!st.empty()) ans[i]++;
            st.push(heights[i]);
        }
        // stack<int> st1;
        // for(int i = 0; i<heights.size(); i++){
        //     while (!st1.empty() && heights[i]>st1.top()){
        //         st1.pop();
        //         ans[i]++;
        //     }
        //     if(!st1.empty()) ans[i]++;
        //     st1.push(heights[i]);
        // }
        return ans;
    }
};