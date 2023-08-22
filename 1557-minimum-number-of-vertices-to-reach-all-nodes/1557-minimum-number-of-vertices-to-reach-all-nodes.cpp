// This problem can be solved using a topological sort algorithm. The idea is to find the vertices that have no incoming edges.
// These vertices form the smallest set of vertices from which all nodes in the graph are reachable.
class Solution {
public:
    vector<int> findSmallestSetOfVertices(int n, vector<vector<int>>& edges) {
        unordered_map<int, int> incoming_edges;
    
        // Iterate over the edges and update the incoming edges count for each vertex
        for (const auto& edge : edges) {
            incoming_edges[edge[1]]++;
        }

        // Find the vertices with no incoming edges
        vector<int> result;
        for (int i = 0; i < n; i++) {
            if (incoming_edges[i] == 0) {
                result.push_back(i);
            }
        }

        return result;
    }
};