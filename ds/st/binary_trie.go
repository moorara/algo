package st

type trieNode struct {
	key      interface{}
	value    interface{}
	terminal bool
	left     *trieNode
	right    *trieNode
}

type binaryTrie struct {
	root   *trieNode
	cmpKey CompareFunc
}

// NewBinaryTrie creates a new binary Trie tree.
//
// Binary Trie is ...
func NewBinaryTrie(cmpKey CompareFunc) OrderedSymbolTable {
	return &binaryTrie{
		root:   nil,
		cmpKey: cmpKey,
	}
}

func (t *binaryTrie) verify() bool {

}

func (t *binaryTrie) Size() int {

}

func (t *binaryTrie) Height() int {

}

func (t *binaryTrie) IsEmpty() bool {

}

func (t *binaryTrie) Put(interface{}, interface{}) {

}

func (t *binaryTrie) Get(interface{}) (interface{}, bool) {

}

func (t *binaryTrie) Delete(interface{}) (interface{}, bool) {

}

func (t *binaryTrie) KeyValues() []KeyValue {

}

func (t *binaryTrie) Min() (interface{}, interface{}) {

}

func (t *binaryTrie) Max() (interface{}, interface{}) {

}

func (t *binaryTrie) Floor(interface{}) (interface{}, interface{}) {

}

func (t *binaryTrie) Ceiling(interface{}) (interface{}, interface{}) {

}

func (t *binaryTrie) Rank(interface{}) int {

}

func (t *binaryTrie) Select(int) (interface{}, interface{}) {

}

func (t *binaryTrie) DeleteMin() (interface{}, interface{}) {

}

func (t *binaryTrie) DeleteMax() (interface{}, interface{}) {

}

func (t *binaryTrie) RangeSize(interface{}, interface{}) int {

}

func (t *binaryTrie) Range(interface{}, interface{}) []KeyValue {

}

func (t *binaryTrie) Traverse(TraverseOrder, VisitFunc) {

}

func (t *binaryTrie) Graphviz() string {

}
