# Heap

| Heap                | Description                                        |
|---------------------|----------------------------------------------------|
| Binary Heap         | Min/Max priority queue.                            |
| Binomial Heap       | Min/Max priority queue.                            |
| Fibonacci Heap      | Min/Max priority queue.                            |
| Indexed Binary Heap | Min/Max priority queue with adjustable priorities. |

  - Use `generic.NewCompareFunc()` for creating a minimum heap.
  - Use`generic.NewReverseCompareFunc()` for creating a maximum heap.

## Comparison

| **Operation** | **Binary Heap** | **Binomial Heap** | **Fibonacci Heap** |
|---------------|:---------------:|:-----------------:|:------------------:|
| Insert        | O(log n)        | O(log n)          | O(1)*              |
| Peek          | O(1)            | O(log n)          | O(1)*              |
| Delete        | O(log n)        | O(log n)          | O(log n)           |
| ChangeKey     | O(log n)        | O(log n)          | O(1)*              |
| Merge         | N/A             | O(log n)          | O(1)*              |
