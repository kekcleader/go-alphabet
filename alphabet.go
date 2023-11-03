package alphabet

import "sort"

// If r is a small Cyrillic letter, converts it to upper-case and returns it,
// otherwise just retuns r.
func up(r rune) rune {
  if 'а' <= r && r <= 'я' {
    r = r - 'а' + 'А'
  }

  if r == 'ё' {
    return 'Ё'
  }

  return r
}

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

// Returns true if runes are equal, case-insensitive if fold equals true.
func EqualRunesEx(a, b rune, fold bool) bool {
  return !fold && a == b ||
          fold && up(a) == up(b)
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

// CompareRunesFold returns true if a < b, otherwise false, case-insensitively.
func CompareRunesFold(a, b rune) bool {
  a, b = up(a), up(b)
  if a == b {
    return false
  }

  if IsCyr(a) && IsCyr(b) {
    if a == 'Ё' {
      return b > 'Е'
    } else if b == 'Ё' {
      return a <= 'Е'
    }
  }

  return a < b
}

// CompareRunesEx returns true if a < b, otherwise false.
// fold means case-insensitivity.
func CompareRunesEx(a, b rune, fold bool) bool {
  if fold {
    return CompareRunesFold(a, b)
  }
  return CompareRunes(a, b)
}

// Returns true if a < b, otherwise false.
// fold means case-insensitivity.
func CompareStringsEx(a, b string, fold bool) bool {
  ar, br := []rune(a), []rune(b)

  for i, x := range ar {
    if i > len(br) {
      return false
    }

    y := br[i]
    if !EqualRunesEx(x, y, fold) {
      return CompareRunesEx(x, y, fold)
    }
  }

  if len(br) > len(ar) {
    return true
  }

  return false
}

// Returns true if a < b, otherwise false.
func CompareStrings(a, b string) bool {
  return CompareStringsEx(a, b, false)
}

// Returns true if a < b, otherwise false, case-insensitively.
func CompareStringsFold(a, b string) bool {
  return CompareStringsEx(a, b, true)
}

// Sorts the slice x.
// fold means case-insensitivity.
func SortEx(x []string, fold bool) {
  sort.Slice(x, func(i, j int) bool {
    return CompareStringsEx(x[i], x[j], fold)
  })
}

// Sorts the slice x
func Sort(x []string) {
  SortEx(x, false)
}

// Sorts the slice x case-insensitively
func SortFold(x []string) {
  SortEx(x, true)
}
