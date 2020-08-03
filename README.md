[![Go Doc][godoc-image]][godoc-url]
[![Build Status][workflow-image]][workflow-url]
[![Go Report Card][goreport-image]][goreport-url]
[![Test Coverage][coverage-image]][coverage-url]
[![Maintainability][maintainability-image]][maintainability-url]

# algo

A collection of common data structures and algorithms for Go applications.

## Summary

  - **Algorithms**
    - Sorts
      - Selection Sort
      - Insertion Sort
      - Shell Sort
      - Merge Sort
      - Quick Sort
      - 3-Way Quick Sort
      - Heap Sort
     - Radix Sorts
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
      - Min Heap
      - Max Heap
      - Indexed Min Heap
      - Indexed Max Heap
    - Symbol Tables
      - Unordered
      - Ordered
        - BST
        - AVL Tree
        - Red-Black Tree
    - Graphs
      - Undirected Graph
      - Directed Graph
      - Weighted Undirected Graph
      - Weighted Directed Graph

## Development

| Command          | Purpose                                     |
|------------------|---------------------------------------------|
| `make test`      | Run unit tests                              |
| `make benchmark` | Run benchmarks                              |
| `make coverage`  | Run unit tests and generate coverage report |


[godoc-url]: https://pkg.go.dev/github.com/moorara/algo
[godoc-image]: https://godoc.org/github.com/moorara/algo?status.svg
[workflow-url]: https://github.com/moorara/algo/actions
[workflow-image]: https://github.com/moorara/algo/workflows/Main/badge.svg
[goreport-url]: https://goreportcard.com/report/github.com/moorara/algo
[goreport-image]: https://goreportcard.com/badge/github.com/moorara/algo
[coverage-url]: https://codeclimate.com/github/moorara/algo/test_coverage
[coverage-image]: https://api.codeclimate.com/v1/badges/48efddf545789eee4132/test_coverage
[maintainability-url]: https://codeclimate.com/github/moorara/algo/maintainability
[maintainability-image]: https://api.codeclimate.com/v1/badges/48efddf545789eee4132/maintainability
