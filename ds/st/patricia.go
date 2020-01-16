package st

type bitString struct {
}

type patriciaNode struct {
	key    interface{}
	value  interface{}
	bitPos int
	left   *patriciaNode
	right  *patriciaNode
}

type patricia struct {
	root   *patriciaNode
	cmpKey CompareFunc
}

// NewPatricia creates a new Patricia tree.
//
// Patricia is ...
func NewPatricia(cmpKey CompareFunc) OrderedSymbolTable {
	return &patricia{
		root:   nil,
		cmpKey: cmpKey,
	}
}

func (t *patricia) verify() bool {

}

func (t *patricia) Size() int {

}

func (t *patricia) Height() int {

}

func (t *patricia) IsEmpty() bool {

}

func (t *patricia) Put(interface{}, interface{}) {

}

func (t *patricia) Get(interface{}) (interface{}, bool) {

}

func (t *patricia) Delete(interface{}) (interface{}, bool) {

}

func (t *patricia) KeyValues() []KeyValue {

}

func (t *patricia) Min() (interface{}, interface{}) {

}

func (t *patricia) Max() (interface{}, interface{}) {

}

func (t *patricia) Floor(interface{}) (interface{}, interface{}) {

}

func (t *patricia) Ceiling(interface{}) (interface{}, interface{}) {

}

func (t *patricia) Rank(interface{}) int {

}

func (t *patricia) Select(int) (interface{}, interface{}) {

}

func (t *patricia) DeleteMin() (interface{}, interface{}) {

}

func (t *patricia) DeleteMax() (interface{}, interface{}) {

}

func (t *patricia) RangeSize(interface{}, interface{}) int {

}

func (t *patricia) Range(interface{}, interface{}) []KeyValue {

}

func (t *patricia) Traverse(TraverseOrder, VisitFunc) {

}

func (t *patricia) Graphviz() string {

}
