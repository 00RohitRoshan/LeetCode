class Solution {
public:
    bool canVisitAllRooms(vector<vector<int>>& rooms) {
        int n = rooms.size();
        vector<int> visited(n,0);
        visited[0] =1;
        dfs(visited,rooms,rooms[0]);
        for(int i =0; i<n; i++){
            cout<<visited[i];
        }
        for(int i=0;i<n;i++){
            if(visited[i] ==0) return false;
        }
        return true;
    }
    
    void dfs(vector<int>& visited,vector<vector<int>>& rooms,vector<int>& keys){
        for(auto it: keys){
            if(visited[it] ==0){
            visited[it] = 1;
            dfs(visited,rooms,rooms[it]);
            }
        }
    }
};