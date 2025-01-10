[![Go Doc][godoc-image]][godoc-url]
[![Build Status][workflow-image]][workflow-url]
[![Go Report Card][goreport-image]][goreport-url]
[![Test Coverage][codecov-image]][codecov-url]

# algo

A collection of common data structures and algorithms for Go applications.

## Summary

  - **Algorithms**
    - Comparative Sorts
      - Selection Sort
      - Insertion Sort
      - Shell Sort
      - Merge Sort
      - Quick Sort
      - 3-Way Quick Sort
      - Heap Sort
    - Non-Comparative Sorts
      - Least Significant Digit
      - Most Significant Digit
      - 3-Way Quick Sort
    - Misc
      - Shuffle
      - Quick Select
  - **Data Structures**
    - Lists
      - Queue
      - Stack
    - Heaps
      - Binary Heaps
      - Binomial Heaps
      - Fibonacci Heaps
    - Sets
      - Union
      - Intersection
      - Difference
      - Powerset,
      - Partitions
    - Symbol Tables
      - Unordered
        - Separate Chaining Hash Table
        - Linear Probing Hash Table
        - Quadratic Probing Hash Table
        - Double Hashing Hash Table
      - Ordered
        - BST
        - AVL Tree
        - Red-Black Tree
        - Tries
          - Binary Trie
          - Patricia Trie
    - Graphs
      - Undirected Graph
      - Directed Graph
      - Weighted Undirected Graph
      - Weighted Directed Graph
    - Automata
      - DFA
      - NFA
    - Grammars
      - Context-Free Grammar
        - Chomsky Normal Form
        - Left Recursion Elimination
        - Left Factoring
        - FIRST and FOLLOW
    - Parsers
      - Two-Buffer Input Reader
      - Parser Combinators

## Development

| Command          | Purpose                                     |
|------------------|---------------------------------------------|
| `make test`      | Run unit tests                              |
| `make benchmark` | Run benchmarks                              |
| `make coverage`  | Run unit tests and generate coverage report |


[godoc-url]: https://pkg.go.dev/github.com/moorara/algo
[godoc-image]: https://pkg.go.dev/badge/github.com/moorara/algo
[workflow-url]: https://github.com/moorara/algo/actions
[workflow-image]: https://github.com/moorara/algo/workflows/Go/badge.svg
[goreport-url]: https://goreportcard.com/report/github.com/moorara/algo
[goreport-image]: https://goreportcard.com/badge/github.com/moorara/algo
[codecov-url]: https://codecov.io/gh/moorara/algo
[codecov-image]: https://codecov.io/gh/moorara/algo/branch/main/graph/badge.svg
