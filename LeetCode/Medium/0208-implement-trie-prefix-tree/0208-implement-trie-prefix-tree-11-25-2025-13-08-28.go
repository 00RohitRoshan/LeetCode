type Trie struct {
    children []*Trie
    isEnd    bool
}

func NewTrie() *Trie {
    return &Trie{children: make([]*Trie, 26)}
}

func Constructor() Trie { return Trie{children: make([]*Trie, 26)} }

func (t *Trie) Insert(word string) {
    curr := t
    for i := 0; i < len(word); i++ {
        idx := word[i] - 'a'
        if curr.children[idx] == nil {
            curr.children[idx] = NewTrie()
        }
        curr = curr.children[idx]
    }
    curr.isEnd = true
}

func (t *Trie) Search(word string) bool {
    node := t.walk(word)
    return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
    return t.walk(prefix) != nil
}

func (t *Trie) walk(s string) *Trie {
    curr := t
    for i := 0; i < len(s); i++ {
        idx := s[i] - 'a'
        if curr.children[idx] == nil {
            return nil
        }
        curr = curr.children[idx]
    }
    return curr
}
