class Solution {
public:
    int orangesRotting(vector<vector<int>>& grid) {
        int m = grid.size(); // Number of rows in the grid
        int n = grid[0].size(); // Number of columns in the grid
        
        // Create a copy of the grid called "visited" to track the state of oranges
        vector<vector<int>> visited = grid;
        
        // Create a queue to store the coordinates of rotten oranges
        queue<pair<int, int>> q;
        
        int countFreshOrange = 0; // Count of fresh oranges
        
        // Iterate through the grid to populate the queue and count fresh oranges
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (visited[i][j] == 2) {
                    q.push({i, j}); // Add rotten oranges to the queue
                }
                if (visited[i][j] == 1) {
                    countFreshOrange++; // Count fresh oranges
                }
            }
        }
        
        // If there are no fresh oranges, return 0 as they are already rotten
        if (countFreshOrange == 0)
            return 0;
        
        // If there are no rotten oranges initially, but there are fresh oranges, 
        // we cannot make them all rotten, so return -1
        if (q.empty())
            return -1;
        
        int minutes = -1; // Initialize the time counter
        
        // Define the four possible directions to move: up, down, left, right
        vector<pair<int, int>> dirs = {{1, 0}, {-1, 0}, {0, -1}, {0, 1}};
        
        // Start the BFS traversal of the grid
        while (!q.empty()) {
            int size = q.size();
            while (size--) { // Process each level of oranges separately
                auto [x, y] = q.front(); // Get the coordinates of the current rotten orange
                q.pop();
                
                // Check all four directions for neighboring fresh oranges
                for (auto [dx, dy] : dirs) {
                    int i = x + dx; // New row coordinate
                    int j = y + dy; // New column coordinate
                    
                    // Check if the new coordinates are within the grid boundaries
                    // and if the orange at the new coordinates is fresh
                    if (i >= 0 && i < m && j >= 0 && j < n && visited[i][j] == 1) {
                        visited[i][j] = 2; // Rot the fresh orange
                        countFreshOrange--; // Decrement the count of fresh oranges
                        q.push({i, j}); // Add the newly rotten orange to the queue
                    }
                }
            }
            minutes++; // Increment the time counter after processing each level
        }
        
        // If there are still fresh oranges left, return -1, as they couldn't be rotten
        if (countFreshOrange > 0)
            return -1;
        
        return minutes; // Return the total time taken to rot all oranges
    }
};