gorph
=====

[![Build Status](https://drone.io/github.com/olanmatt/gorph/status.png)](https://drone.io/github.com/olanmatt/gorph/latest)

A Go graph library

Usage
-----

To come soon.


Runtimes
--------

Storage space is O(\\V\\ + 2\\E\\).

| Operation | Best Case | Worst Case        | Average Case               |
| --------- | :-------: | :---------------: | :------------------------: |
| Add Node  | O(\\V\\)  | O(\\V\\)          | O(\\V\\)                   |
| Del Node  | O(1)      | O(\\V\\)          | O(\\V\\ / 2)               |
| Add Edge  | O(1)      | O(2\\V\\ + \\E\\) | O(\\V\\ + \\E\\ / \\V\\)   |
| Del Edge  | O(1)      | O(\\V\\ + \\E\\)  | O(\\V\\ / 2 + \\E\\ / 2)   |
| Adjacent  | O(1)      | O(\\E\\)          | O(\\E\\ / 2)               |
| Neighbors | O(\\E\\)  | O(\\E\\)          | O(\\E\\)                   |
| Search    | O(1)      | O(\\V\\)          | O(\\V\\ / 2)               |

*\\V\\ and \\E\\ denotes size of node and edge arrays respectively*
