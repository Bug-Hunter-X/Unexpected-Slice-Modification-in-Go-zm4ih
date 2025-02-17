# Unexpected Slice Modification in Go

This repository demonstrates a common but subtle error in Go involving slice manipulation.  When a slice is created using slicing (`[:]`), it shares the underlying array with the original slice.  Appending to a slice that shares an underlying array can modify the original slice unexpectedly, leading to bugs that are hard to track.

The `bug.go` file showcases this issue, while `bugSolution.go` provides a corrected version.

## How to Reproduce

1. Clone the repository.
2. Run `go run bug.go` to observe the unexpected behavior.
3. Run `go run bugSolution.go` to see the correct solution.

## Solution

The solution involves creating a copy of the slice using `append([]int(nil), y...)` before appending to it.  This creates a new slice with its own underlying array, thereby preventing the unexpected modification of the original slice.  See `bugSolution.go` for details.