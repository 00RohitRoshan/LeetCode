class Solution {
public:
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        vector<vector<int>> res;
        sort(intervals.begin(),intervals.end());
        vector<int> temp = intervals[0];
        // for(int i=0; i<intervals.size(); i++){
        //     cout<<"["<<intervals[i][0]<<","<<intervals[i][1]<<"]"<<",";
        // }
        int i =1;
        while(i<intervals.size()){
            // cout<<i<<","<<"["<<temp[0]<<","<<temp[1]<<"]"<<",";
            if(intervals[i][0] <= temp[1]){
                temp[1]=max(intervals[i][1],temp[1]);
                i++;
            }else if(intervals[i][0] > temp[1]){
                res.push_back(temp);
                temp = intervals[i];
                i++;
            }
        }
        res.push_back(temp);
        // cout<<endl;
        return res;
    }
};