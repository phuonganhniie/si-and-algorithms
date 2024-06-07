package helper

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

// Initializes a new Trie node
func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

// Initializes a new Trie
func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// Inserts a word into the Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exist := node.children[ch]; !exist {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch] // move to the child node
	}
	node.isEnd = true
}

/*
Bắt đầu từ nút gốc và duyệt qua từng ký tự của từ. Với mỗi ký tự:
- Nếu nút hiện tại được đánh dấu là kết thúc của một từ, chúng ta trả về tiền tố đã tích lũy.
- Nếu nút con cho ký tự đó tồn tại, chúng ta di chuyển đến nút đó và thêm ký tự vào tiền tố.
- Nếu nút con cho ký tự đó không tồn tại, chúng ta dừng lại.
- Nếu duyệt hết các ký tự mà không tìm thấy gốc, chúng ta trả về từ gốc ban đầu.
*/
// Find shortest root for a word
func (t *Trie) Search(word string) string {
	node := t.root
	prefix := ""
	for _, ch := range word {
		if node.isEnd {
			return prefix
		}
		if _, exist := node.children[ch]; exist {
			node = node.children[ch]
			prefix += string(ch)
		} else {
			break
		}
	}

	if node.isEnd {
		return prefix
	}
	return word
}

// visualizeNode recursively visualizes the Trie node
func visualizeNode(node *TrieNode, prefix string) {
	for ch, child := range node.children {
		fmt.Printf("%s -> %c", prefix, ch)
		if child.isEnd {
			fmt.Print("*")
		}
		fmt.Println()
		visualizeNode(child, prefix+"    ")
	}
}

// Visualize visualizes the entire Trie
func (t *Trie) Visualize() {
	visualizeNode(t.root, "")
}
