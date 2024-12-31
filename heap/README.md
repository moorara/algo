# Heap

A heap is a tree-based data structure that satisfies the **heap property**:

  - **Max-Heap Property**: For any given node n,
    the value of n is greater than or equal to the values of its children.
    This means the largest value is at the root.
  - **Min-Heap Property**: For any given node n,
    the value of n is less than or equal to the values of its children.
    This means the smallest value is at the root.

The heap is one maximally efficient implementation of a **pririoty queue** ADT (*abstract datat type*).
In fact, heaps are often considered equivalent to priority queues.

| Heap                   | Description                                        |
|------------------------|----------------------------------------------------|
| Binary Heap            | Min/Max priority queue.                            |
| Binomial Heap          | Min/Max priority queue.                            |
| Fibonacci Heap         | Min/Max priority queue.                            |
| Indexed Binary Heap    | Min/Max priority queue with adjustable priorities. |
| Indexed Binomial Heap  | Min/Max priority queue with adjustable priorities. |
| Indexed Fibonacci Heap | Min/Max priority queue with adjustable priorities. |

## Comparison

| **Operation** | **Binary Heap** | **Binomial Heap** | **Fibonacci Heap** |
|---------------|:---------------:|:-----------------:|:------------------:|
| Insert        | O(log n)        | O(1)ᵃᵐᵒʳᵗⁱᶻᵉᵈ     | O(1)               |
| Peek          | O(1)            | O(1)              | O(1)               |
| Delete        | O(log n)        | O(log n)          | O(log n)ᵃᵐᵒʳᵗⁱᶻᵉᵈ  |
| ChangeKey     | O(log n)        | O(log n)          | O(1)ᵃᵐᵒʳᵗⁱᶻᵉᵈ      |
| Merge         | N/A             | O(log n)          | O(1)               |

## Quick Start

  - Use `generic.NewCompareFunc()` for creating a minimum heap.
  - Use`generic.NewReverseCompareFunc()` for creating a maximum heap.

## Binary Heap

A binary heap is a **complete binary tree** that satisfies the heap property.

Since a binary heap is a compelete tree, it can be implemented using an array (slice).

  - The root of tree is stored at index `1` (index `0` is left unused).
  - The parent of node at index `i` is located at index `i / 2`.
  - The left and right children of node node at index `i` are located at indices `2i` and `2i + 1`.

## Binomial Heap

Binomial heap is an implementation of the **mergeable** heap ADT, a priority queue supporting merge operation.
A binomial heap is implemented as a set of **binomial trees** that satisfy the binomial heap properties.

  - **Heap Property**: Every binomial tree in a binomial heap satisfies the min-heap or max-heap property.
  - **Structural Property**: The heap contains at most one binomial tree of any given order.

A binomial tree `Bₖ` of order `k` is defined recursively:

  - A binomial tree `B₀` of order `0` is a single node.
  - A binomial tree `Bₖ` of order `k` is formed by linking two binomial trees of orders `k-1` together,
    making the root of one tree a child of the root of the other tree.
    Equivalently, a binomial tree `Bₖ` has a root node whose children are roots of binomial trees of orders `k-1`, `k-2`, ..., `1`, `0`.

Here are some properties of binomial trees:

  - The height of a `Bₖ` tree is `k`.
  - The number of nodes in a `Bₖ` tree is `2ᵏ`.
  - The root of a `Bₖ` tree has `k` children.
  - The children of the root of a `Bₖ` tree are the roots of `B₀`, `B₁`, ..., `Bₖ₋₁` trees.
  - A binomial tree `Bₖ` of order `k` has `C(k, d)` nodes at depth `d`, a **binomial coefficient**.

## Fibonacci Heap
