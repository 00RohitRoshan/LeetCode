// #include <vector>
// using namespace std;

class Solution {
public:
    vector<vector<int>> allPathsSourceTarget(vector<vector<int>>& graph) {
        vector<vector<int>> res;
        vector<int> path{0};
        dfs(graph, 0, path, res);
        return res;
    }
private:
    void dfs(vector<vector<int>>& graph, int node, vector<int>& path, vector<vector<int>>& res) {
        if (node == graph.size() - 1) {
            res.push_back(path);
        } else {
            for (int nei : graph[node]) {
                path.push_back(nei);
                dfs(graph, nei, path, res);
                path.pop_back();
            }
        }
    }
};
