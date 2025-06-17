/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    // ListNode* reverseList(ListNode* head) {
    //     ListNode* n = NULL ;
    //     ListNode* t ;
    //     while(head){
    //         t = head->next;
    //         head->next = n;
    //         n = head;
    //         head = t;
    //     }
    //     return n;
    // }
    ListNode* reverseList(ListNode* head) {
        if (!head || !head->next) return head;

        ListNode* n = reverseList(head->next);
        head->next->next = head;
        head->next = NULL;
        return n;
    }
};