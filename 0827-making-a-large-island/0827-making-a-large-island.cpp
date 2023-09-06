class disjointSet{
    public:

    vector<int>parent,size;

    disjointSet(int v){
        size.resize(v,1);
        parent.resize(v);
        for(int i=0;i<v;i++){
            parent[i]=i;
        }
    }


    int findUPar(int node){
        if(node==parent[node])return node;

        return parent[node]=findUPar(parent[node]);
    }    


    void unionBySize(int u, int v){
        int ulp_u=findUPar(u);
        int ulp_v=findUPar(v);
        if(ulp_u==ulp_v)return;

        if(size[ulp_v]>size[ulp_u]){
            parent[ulp_u]=ulp_v;
            size[ulp_v]+=size[ulp_u];
        }
        else{
            parent[ulp_v]=ulp_u;
            size[ulp_u]+=size[ulp_v];
        }
    }
};

class Solution {
public:
    int largestIsland(vector<vector<int>>& grid) {
        int n=grid.size();
       
        disjointSet s(n*n);
        vector<vector<int>>visited(n,vector<int>(n,0));
        for(int i=0;i<n;i++){
            for(int j=0;j<n;j++){
                if(grid[i][j] && !visited[i][j]){
                    visited[i][j]=1;
                    int dx[]={-1,0,1,0};
                    int dy[]={0,-1,0,1};
                    for(int k=0;k<4;k++){
                        int x=i+dx[k];
                        int y=j+dy[k];
                        if(x<0 || x>=n || y<0 || y>=n)continue;
                        if(grid[x][y]==1){
                            if(s.findUPar(n*x + y)!=s.findUPar(n*i + j)){
                                s.unionBySize(n*x + y, n*i + j);
                            }
                        }
                    }
                }
            }
        }
        int ans=0;
        for(int i=0;i<n*n;i++){
            ans=max(ans,s.size[i]);
        }

        for(int i=0;i<n;i++){
            for(int j=0;j<n;j++){
                if(grid[i][j]==0){
                    int count=1;
                    vector<int>ultimate;
                    int dx[]={-1,0,1,0};
                    int dy[]={0,-1,0,1};
                    for(int k=0;k<4;k++){
                        int x=i+dx[k];
                        int y=j+dy[k];
                        if(x<0 || x>=n || y<0 || y>=n)continue;
                        if(visited[x][y]){
                            int ultP=s.findUPar(x*n + y);
                            if(std::find(ultimate.begin(),ultimate.end(),ultP)==ultimate.end()){
                                ultimate.push_back(ultP);
                                count+=s.size[ultP];
                            }
                        }
                    }
                    ans=max(ans,count);
                }
            }
        }

        return ans;
    }
};