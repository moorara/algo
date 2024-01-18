
| Sort      | Best Case         | Average Case      | Worst Case        | Memory | Comment                                       |
|-----------|:-----------------:|:-----------------:|:-----------------:|:------:|-----------------------------------------------|
| Selection | N<sup>2</sup> / 2 | N<sup>2</sup> / 2 | N<sup>2</sup> / 2 | 1      | N exchanges                                   |
| Insertion | N<sup>2</sup> / 2 | N<sup>2</sup> / 2 | N<sup>2</sup> / 2 | 1      | **Stable**, N exchanges, Suitable for small N |
| Shell     | N                 | ?                 | ?                 | 1      | Tight code, Subquadratic                      |
| Merge     | NlgN              | NlgN              | NlgN              | N      | **Stable**, **NlgN guarantee**, Extra memory  |
| MergeRec  | NlgN              | NlgN              | NlgN              | N      | **Stable**, **NlgN guarantee**, Extra memory  |
| Heap      | NlgN              | NlgN              | NlgN              | 1      | **NlgN guarantee**                            |
| Quick     | NlgN              | 2NlnN             | N<sup>2</sup> / 2 | clogN  | **NlgN probabilistic guarantee**              |
| Quick3Way | N                 | 2NlnN             | N<sup>2</sup> / 2 | clogN  | Faster in presence of duplicate keys          |

By running benchmarks, you can compare the performance of these algorithms with each other and also with built-in sort algorithm of go.
