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
  - Use `generic.NewReverseCompareFunc()` for creating a maximum heap.

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

Fibonacci heap is an implementation of the **mergeable** heap ADT, a priority queue supporting merge operation.
A Fibonacci heap is implemented as a collection of heap-ordered trees.
It has a better amortized running time than binary and binomial heaps.

Fibonacci heaps are more flexible than binomial heaps, as their trees do not have a predetermined shape.
In the extreme case, a Fibonacci heap can have every item in a separate tree.
This flexibility allows some operations to be executed in a lazy manner,
postponing the work for later operations.

To maintain the desired running time, order must eventually be introduced.
Node degrees (the number of direct children) are kept low, with each node having a degree of at most `logᵩn`.
Additionally, a node with degree `k` has a subtree size of at least `Fₖ₊₂`,
where `Fᵢ` is the i-th Fibonacci number.
This is enforced by allowing at most one child to be cut from a non-root node.
If a second child is cut, the node itself is also cut and becomes a new tree root.

### Proof of Degree Bounds

The Fibonacci sequence is defined as:

	F₀ = 0           n = 0
	F₁ = 1           n = 1
	Fₙ = Fₙ₋₁ + Fₙ₋₂  n ≥ 2

The amortized performance of a Fibonacci heap depends on the fact that the degree of any node
is bounded by `O(logn)`, where `n` is the total number of items in the heap.
This bound ensures that key operations, such as `Delete`, remain efficient.

In a Fibonacci heap, the size of the subtree rooted at any node `x` of degree `k` is at least `Fₖ₊₂`.
It can be shown (via direct proof or induction) that `Fₖ₊₂ ≥ φᵏ` for all `𝑘 ≥ 2`,
where `φ = (1 + √5) / 2` is the golden ratio.

	Fₖ₊₂ = (φᵏ⁺² - (-φ)⁻⁽ᵏ⁺²⁾) / √5
	Fₖ₊₂ ≈ φᵏ⁺² / √5
	Fₖ₊₂ ≈ φᵏ (φ² / √5)
	Fₖ₊₂ ≈ φᵏ 1.17082039324993680829
	Fₖ₊₂ ≥ φᵏ

This exponential growth guarantees that the degree of a node remains logarithmic
in terms of the size of the heap, supporting the efficient amortized performance of Fibonacci heaps.

	n ≥ Fₖ₊₂ ≥ φᵏ
	k ≤ logᵩn

Let `x` be an arbitrary node in a Fibonacci heap.
Define `size(x)` to be the number of nodes in the subtree rooted at `x`.
We aim to prove by induction that `size(x) ≥ Fₖ₊₂`, where `k` is the degree of `x`.

#### Base Case

If `x` has height `0`, then `k = 0` (no children), and `size(x) = 1`.
This satisfies `Fₖ₊₂ = F₂ = 1`, completing the base case.

#### Inductive Case

Suppose `x` has degree `k`. Let `y₁`, `y₂`, ..., `yₖ` denote the children of `x` in the order
they were added (chronological order), and let `d₁`, `d₂`, ..., `dₖ` be their respective degrees.

##### Claim

For each `i`, `dᵢ ≥ i - 2`. 

**Proof of claim:** Just before `yᵢ` was added as a child of `x`, `x` already had `i - 1` children.
Since merging occurs only when roots have equal degrees, `yᵢ` must have had degree at least `i - 1`, at that time.
The deletion algorithm ensures `yᵢ` loses at most one child afterward, so `dᵢ ≥ i - 2`.

##### Applying The Inductive Hypothesis

Since the height of each `yᵢ` is less than that of `x`, we apply the inductive hypothesis:

$$\text{size}(x) \ge F_{d_i+2} \ge F_{(i-2)+2} = F_i$$

##### Combining Results

The nodes `x` and `y` each contribute at list `1` to size(x), and so we have:

$$
\begin{align*}
\text{size}(x) &\ge 2 + \sum_{i=2}^{k} \text{size}(y_i) \\
               &\ge 2 + \sum_{i=2}^{k} F_i = 1 + \sum_{i=0}^{k} F_i = F_{k+2}
\end{align*}
$$

This completes the inductive step. Hence, by induction, `size(x) ≥ Fₖ₊₂` for all `x`.
