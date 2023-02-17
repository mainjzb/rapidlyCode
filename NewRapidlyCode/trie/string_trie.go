package trie

type StringTrie struct {
	segmenter StringSegmenter // key segmenter, must not cause heap allocs
	value     interface{}
	children  map[string]*StringTrie
}

type Info struct {
	RawText       string // 一个单词，标点符号 等
	RawTextAppend string // 跟随在单词后面的空格等，在复原文本时有用
}

// NewStringTrie allocates and returns a new *PathTrie.
func NewStringTrie() *StringTrie {
	return &StringTrie{
		segmenter: BlackSegmenter,
	}
}

// newPathTrieFromTrie returns new trie while preserving its config
func (t *StringTrie) newStringTrie() *StringTrie {
	return &StringTrie{
		segmenter: t.segmenter,
	}
}

// Get returns the value stored at the given key. Returns nil for internal
// nodes or for nodes with a value of nil.
func (t *StringTrie) Get(key string) interface{} {
	node := t
	for part, i := t.segmenter(key, 0); part != ""; part, i = t.segmenter(key, i) {
		node = node.children[part]
		if node == nil {
			return nil
		}
	}
	return node.value
}

// Put inserts the value into the trie at the given key, replacing any
// existing items. It returns true if the put adds a new value, false
// if it replaces an existing value.
// Note that internal nodes have nil values so a stored nil value will not
// be distinguishable and will not be included in Walks.
func (t *StringTrie) Put(key string, value interface{}) bool {
	node := t
	for part, i := t.segmenter(key, 0); part != ""; part, i = t.segmenter(key, i) {
		child, _ := node.children[part]
		if child == nil {
			if node.children == nil {
				node.children = map[string]*StringTrie{}
			}
			child = t.newStringTrie()
			node.children[part] = child
		}
		node = child
	}
	// does node have an existing value?
	isNewVal := node.value == nil
	node.value = value
	return isNewVal
}

// Delete removes the value associated with the given key. Returns true if a
// node was found for the given key. If the node or any of its ancestors
// becomes childless as a result, it is removed from the trie.
func (t *StringTrie) Delete(key string) bool {
	var path []StringTrieNode // record ancestors to check later
	node := t
	for part, i := t.segmenter(key, 0); part != ""; part, i = t.segmenter(key, i) {
		path = append(path, StringTrieNode{part: part, node: node})
		node = node.children[part]
		if node == nil {
			// node does not exist
			return false
		}
	}
	// delete the node value
	node.value = nil
	// if leaf, remove it from its parent's children map. Repeat for ancestor path.
	if node.isLeaf() {
		// iterate backwards over path
		for i := len(path) - 1; i >= 0; i-- {
			parent := path[i].node
			part := path[i].part
			delete(parent.children, part)
			if !parent.isLeaf() {
				// parent has other children, stop
				break
			}
			parent.children = nil
			if parent.value != nil {
				// parent has a value, stop
				break
			}
		}
	}
	return true // node (internal or not) existed and its value was nil'd
}

func (t *StringTrie) Find(doc string) {
	// docNode := make([]DocNode, 0, 32)
	// for part, i := t.segmenter(doc, 0); part != ""; part, i = t.segmenter(doc, i) {
	// 	docNode = append(docNode, DocNode{Value: part, Index: i - 1 - len(part)})
	// }

}

// Walk iterates over each key/value stored in the trie and calls the given
// walker function with the key and value. If the walker function returns
// an error, the walk is aborted.
// The traversal is depth first with no guaranteed order.
func (t *StringTrie) Walk(walker WalkFunc) error {
	return t.walk("", walker)
}

// WalkPath iterates over each key/value in the path in trie from the root to
// the node at the given key, calling the given walker function for each
// key/value. If the walker function returns an error, the walk is aborted.
func (t *StringTrie) WalkPath(key string, walker WalkFunc) error {
	// Get root value if one exists.
	if t.value != nil {
		if err := walker("", t.value); err != nil {
			return err
		}
	}
	for part, i := t.segmenter(key, 0); ; part, i = t.segmenter(key, i) {
		if t = t.children[part]; t == nil {
			return nil
		}
		if t.value != nil {
			var k string
			if i == -1 {
				k = key
			} else {
				k = key[0:i]
			}
			if err := walker(k, t.value); err != nil {
				return err
			}
		}
		if i == -1 {
			break
		}
	}
	return nil
}

// StringTrieNode node and the part string key of the child the path descends into.
type StringTrieNode struct {
	node *StringTrie
	part string
}

func (t *StringTrie) walk(key string, walker WalkFunc) error {
	if t.value != nil {
		if err := walker(key, t.value); err != nil {
			return err
		}
	}
	for part, child := range t.children {
		if err := child.walk(key+part, walker); err != nil {
			return err
		}
	}
	return nil
}

func (t *StringTrie) isLeaf() bool {
	return len(t.children) == 0
}
