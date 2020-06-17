package filter

// Node 敏感词结点
type Node struct {
	next  map[rune]*Node
	isEnd bool
}

// NewNode 创建新节点
func NewNode() *Node {
	return &Node{}
}

// add new word on filter tree
func (n *Node) add(words []rune) {
	if len(words) == 0 {
		return
	}
	if n.next == nil {
		n.next = map[rune]*Node{}
	}
	key := words[0]
	node, bHas := n.next[key]
	if !bHas {
		node = &Node{}
	}

	if len(words) == 1 {
		node.isEnd = true
	} else {
		node.add(words[1:])
	}
	n.next[key] = node
}

// remove delete key word from tree
func (n *Node) remove(words []rune) {
	if len(words) == 0 || n.next == nil {
		return
	}
	key := words[0]
	node, bHas := n.next[key]
	if !bHas {
		return
	}
	if len(words) == 1 {
		if len(node.next) == 0 {
			delete(n.next, key)
		} else {
			node.isEnd = false
			n.next[key] = node
		}
		return
	}
	node.remove(words[1:])
	if len(node.next) == 0 && !node.isEnd {
		delete(n.next, key)
	}
}

// mlength return max match length,pos: start position
func (n *Node) mlength(text []rune, pos int) (end int) {
	if len(text) == 0 || n.next == nil {
		return 0
	}
	key := text[0]
	node, bHas := n.next[key]
	if !bHas {
		return 0
	}
	if len(text) == 1 {
		if node.isEnd {
			return pos + 1
		}
		return 0
	}
	end = node.mlength(text[1:], pos+1)
	if node.isEnd && end == 0 {
		return pos + 1
	}
	return end
}

// contains return true for any match
func (n *Node) contains(text []rune) bool {
	if len(text) == 0 || n.next == nil {
		return false
	}
	key := text[0]
	node, bHas := n.next[key]
	if !bHas {
		return false
	}
	if node.isEnd {
		return true
	}
	return node.contains(text[1:])
}
