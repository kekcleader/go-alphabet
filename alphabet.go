package alphabet

import "sort"

// Returns true if r is a Cyrillic letter Yo (big or small).
func IsYo(r rune) bool {
  return r == 'ё' || r == 'Ё'
}

// Returns true if r is a Russian letter.
func IsCyr(r rune) bool {
  return 'а' <= r && r <= 'я' ||
         'А' <= r && r <= 'Я' ||
         IsYo(r)
}

// CompareRunes returns true if a < b, otherwise false.
func CompareRunes(a, b rune) bool {
  if a == b {
    return false
  }

  if IsCyr(a) && IsCyr(b) {
    if a == 'ё' {
      return b > 'е'
    } else if a == 'Ё' {
      return b > 'Е'
    } else if b == 'ё' {
      return a <= 'е'
    } else if b == 'Ё' {
      return a <= 'Е'
    }
  }

  return a < b
}

// Returns true if a < b, otherwise false.
func CompareStrings(a, b string) bool {
  ar, br := []rune(a), []rune(b)

  for i, x := range ar {
    if i > len(br) {
      return false
    }

    y := br[i]
    if CompareRunes(x, y) {
      return true
    }
  }

  return false
}

// Sorts the slice x
func Sort(x []string) {
  sort.Slice(x, func(i, j int) bool {
    return CompareStrings(x[i], x[j])
  })
}
