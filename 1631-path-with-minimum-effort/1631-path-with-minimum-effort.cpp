class Solution {
public:
    int minimumEffortPath(vector<vector<int>>& heights) {
        int row = heights.size(), col = heights[0].size();
        int dx[4] = {-1, +1, 0, 0}, dy[4] = {0, 0, -1, +1};
        priority_queue<pair<int, pair<int, int>>, vector<pair<int, pair<int, int>>>, greater<pair<int, pair<int, int>>>> pq;
        vector<vector<int>> dist(row, vector<int>(col, 1e9));
        pq.push({0, {0, 0}});
        dist[0][0] = 0;

        while(pq.size())
        {
            int weight = pq.top().first;
            int i = pq.top().second.first;
            int j = pq.top().second.second;
            pq.pop();

            //Traversing through its adjacent nodes.
            for(int k = 0; k < 4; k++)
            {
                int newi = i + dx[k];
                int newj = j + dy[k];
                //Checking for valid condition and updating the effort and pushing into priority queue if the previous effort is larger than present effort of the neighbour nodes.
                if(newi >= 0 && newi < row && newj >= 0 && newj < col && dist[newi][newj] > max(abs(heights[newi][newj] - heights[i][j]), weight))
                {
                    dist[newi][newj] = max(abs(heights[newi][newj] - heights[i][j]), weight);
                    pq.push({dist[newi][newj], {newi, newj}});
                }
            }
        }
        //Returning the minimum effort to reach the end point (row-1, col-1).
        return dist[row-1][col-1];
    }
};