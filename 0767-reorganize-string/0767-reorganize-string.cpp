// This problem can be solved using a greedy approach with the help of a priority queue or max heap data structure. 
// The idea is to put the highest frequency character first and then use a series of swaps to rearrange the characters
class Solution {
public:
    string reorganizeString(string s) {
        return rearrangeString( s);
    }
    string rearrangeString(string s) {
        // Create a frequency map of characters
        unordered_map<char, int> freq;
        for (char c : s) freq[c]++;

        // Create a max heap of pairs (frequency, character)
        priority_queue<pair<int, char>> maxHeap;
        for (auto it : freq) maxHeap.push({it.second, it.first});

        string result = "";
        pair<int, char> prev = {-1, '#'};

        while (!maxHeap.empty()) {
            // Pop the top element from the max heap
            pair<int, char> curr = maxHeap.top();
            maxHeap.pop();

            // Add the character to the result string
            result += curr.second;

            // Decrease the frequency of the character
            curr.first--;

            // Push the previous character back into the max heap if its frequency is greater than 0
            if (prev.first > 0) maxHeap.push(prev);

            // Update the previous character
            prev = curr;
        }

        // Check if the length of the result string is equal to the length of the input string
        if (result.length() != s.length()) return "";

        return result;
    }
};
// time complexity of this solution is O(n log n)
// space complexity is O(n)