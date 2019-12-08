package list

// Queue represents a queue abstract data type.
type Queue interface {
	Size() int
	IsEmpty() bool
	Enqueue(interface{})
	Dequeue() interface{}
	Peek() interface{}
	Contains(interface{}) bool
}

type arrayQueue struct {
	listSize       int
	nodeSize       int
	frontNodeIndex int
	rearNodeIndex  int
	frontNode      *arrayNode
	rearNode       *arrayNode
	cmp            CompareFunc
}

// NewQueue creates a new array-list queue.
func NewQueue(nodeSize int, cmp CompareFunc) Queue {
	return &arrayQueue{
		listSize:       0,
		nodeSize:       nodeSize,
		frontNodeIndex: -1,
		rearNodeIndex:  -1,
		frontNode:      nil,
		rearNode:       nil,
		cmp:            cmp,
	}
}

// Size returns the number of items in queue.
func (q *arrayQueue) Size() int {
	return q.listSize
}

// IsEmpty returns true if queue is empty.
func (q *arrayQueue) IsEmpty() bool {
	return q.listSize == 0
}

// Enqueue adds a new item to queue.
func (q *arrayQueue) Enqueue(item interface{}) {
	if q.frontNode == nil && q.rearNode == nil {
		q.frontNodeIndex = 0
		q.rearNodeIndex = 0
		q.frontNode = newArrayNode(q.nodeSize, nil)
		q.rearNode = q.frontNode
	}

	q.listSize++
	q.rearNode.block[q.rearNodeIndex] = item
	q.rearNodeIndex++

	if q.rearNodeIndex == q.nodeSize {
		q.rearNodeIndex = 0
		q.rearNode.next = newArrayNode(q.nodeSize, nil)
		q.rearNode = q.rearNode.next
	}
}

// Dequeue removes an item from queue.
func (q *arrayQueue) Dequeue() interface{} {
	if q.listSize == 0 {
		return nil
	}

	q.listSize--
	item := q.frontNode.block[q.frontNodeIndex]
	q.frontNodeIndex++

	if q.frontNodeIndex == q.nodeSize {
		q.frontNodeIndex = 0
		q.frontNode = q.frontNode.next
	}

	return item
}

// Peek returns the next item in queue without removing it from queue.
func (q *arrayQueue) Peek() interface{} {
	if q.listSize == 0 {
		return nil
	}

	return q.frontNode.block[q.frontNodeIndex]
}

// Contains returns true if a given item is already in queue.
func (q *arrayQueue) Contains(item interface{}) bool {
	n := q.frontNode
	i := q.frontNodeIndex

	for n != nil && (n != q.rearNode || i <= q.rearNodeIndex) {
		if q.cmp(n.block[i], item) == 0 {
			return true
		}

		i++
		if i == q.nodeSize {
			n = n.next
			i = 0
		}
	}

	return false
}
