class Solution {
public:
    int findTheCity(int n, vector<vector<int>>& edges, int distanceThreshold) {
        
        vector<vector<int>> dist(n,vector<int>(n,1e9));
        //creating asjancy matrix
        for(auto node : edges){
            dist[node[0]][node[1]] = node[2];
            dist[node[1]][node[0]] = node[2];

        }
        //setting distance of node to itself = 0 
        for(int i=0;i<n;i++) dist[i][i]=0;
        //applying Floyd Warshall algorithm to find min distance between each node
        for(int k=0;k<n;k++){
            for(int i=0;i<n;i++){
                for(int j=0;j<n;j++){
                 if((dist[i][j] > dist[i][k] + dist[k][j]) && dist[i][k] !=1e9 && dist[k][j] !=1e9){
                        dist[i][j] = dist[i][k] + dist[k][j];
                    }
                }
            }
        }
        vector<int> Number_of_reachable_Cities(n,0);
        
        //traversing the distance 2D array to find the number of cities whose  
        //min distance  from city i is less than and equal to distanceThreshold
        for(int i=0;i<n;i++){
            for(int j=0;j<n;j++){
                if(i==j) continue;
                if(dist[i][j]<=distanceThreshold){
                     Number_of_reachable_Cities[i]++;  
                }
            }
        }

        int mn=1e9;
        // find the minimum in Number_of_reachable_Cities this wil help to find
        // smallest all the cities 
        for(int i=0;i<n;i++){ 
            mn = min(mn,Number_of_reachable_Cities[i]);
        }
        int mx=-1;
        //finding  city with the greatest number If there are multiple such cities
        for(int i=0;i<n;i++){
            if(Number_of_reachable_Cities[i] == mn) mx = max(mx,i);  
        }
        return mx;
    }
};