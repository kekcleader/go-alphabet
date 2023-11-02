package alphabet

import "sort"

// Sorts the slice x
func Sort(x []string) {
  sort.Slice(x, func(i, j int) bool {
    return x[i] < x[j]
  })
}
