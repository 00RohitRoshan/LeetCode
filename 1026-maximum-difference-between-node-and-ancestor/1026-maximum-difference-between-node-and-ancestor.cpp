/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int maxDiff(TreeNode* node, int maxval, int minval){
        if(node == nullptr) return maxval-minval;
        
        maxval = max(maxval, node->val);
        minval = min(minval, node->val);
        
        int maxDiffLeft = maxDiff(node->left, maxval, minval);
        int maxDiffRight = maxDiff(node->right, maxval, minval);
        
        return max(maxDiffLeft,maxDiffRight);
    }
    int maxAncestorDiff(TreeNode* root) {
        if(root == nullptr) return 0;
        return maxDiff(root, root->val, root->val);
    }
};