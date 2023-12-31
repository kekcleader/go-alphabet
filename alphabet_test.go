package alphabet

import (
  "testing"
  "reflect"
  "fmt"
)

func TestCompareRunes(t *testing.T) {
  // equal
  x := 'б'
  if CompareRunes(x, x) {
    t.Errorf("CompareRunes('%c', '%c') = true; want false", x, x)
  }

  x = 'ж'
  if CompareRunes(x, x) {
    t.Errorf("CompareRunes('%c', '%c') = true; want false", x, x)
  }

  x = 'Ю'
  if CompareRunes(x, x) {
    t.Errorf("CompareRunes('%c', '%c') = true; want false", x, x)
  }

  x = 'w'
  if CompareRunes(x, x) {
    t.Errorf("CompareRunes('%c', '%c') = true; want false", x, x)
  }

  // less
  x, y := 'в', 'я'
  if !CompareRunes(x, y) {
    t.Errorf("CompareRunes('%c', '%c') = false; want true", x, y)
  }

  x, y = 'Е', 'Ё'
  if !CompareRunes(x, y) {
    t.Errorf("CompareRunes('%c', '%c') = false; want true", x, y)
  }

  x, y = 'е', 'ё'
  if !CompareRunes(x, y) {
    t.Errorf("CompareRunes('%c', '%c') = false; want true", x, y)
  }

  x, y = 'е', 'ж'
  if !CompareRunes(x, y) {
    t.Errorf("CompareRunes('%c', '%c') = false; want true", x, y)
  }

  x, y = 'д', 'е'
  if !CompareRunes(x, y) {
    t.Errorf("CompareRunes('%c', '%c') = false; want true", x, y)
  }

  x, y = 'б', 'ё'
  if !CompareRunes(x, y) {
    t.Errorf("CompareRunes('%c', '%c') = false; want true", x, y)
  }

  // false
  x, y = 'ё', 'г'
  if CompareRunes(x, y) {
    t.Errorf("CompareRunes('%c', '%c') = true; want false", x, y)
  }

  x, y = 'Ю', 'Г'
  if CompareRunes(x, y) {
    t.Errorf("CompareRunes('%c', '%c') = true; want false", x, y)
  }

  x, y = 'ё', 'е'
  if CompareRunes(x, y) {
    t.Errorf("CompareRunes('%c', '%c') = true; want false", x, y)
  }
}

func TestCompareStrings(t *testing.T) {
  x, y := "КОНЬ", "КОНЬЯК"
  if !CompareStrings(x, y) {
    t.Errorf("CompareStrings(\"%s\", \"%s\") = false; want true", x, y)
  }

  x, y = "конь", "коньяк"
  if !CompareStrings(x, y) {
    t.Errorf("CompareStrings(\"%s\", \"%s\") = false; want true", x, y)
  }

  x, y = "Кисель", "Киселёв"
  if CompareStrings(x, y) {
    t.Errorf("CompareStrings(\"%s\", \"%s\") = true; want false", x, y)
  }

  x, y = "АБВ", "АВБ"
  if !CompareStrings(x, y) {
    t.Errorf("CompareStrings(\"%s\", \"%s\") = false; want true", x, y)
  }
}

func TestSort(t *testing.T) {
  x := []string{"яблоко", "ананас", "Дыня", "посёлок", "поселок"}
  s := fmt.Sprintf("%v", x)
  expect := []string{"Дыня", "ананас", "поселок", "посёлок", "яблоко"}
  Sort(x)
  if !reflect.DeepEqual(x, expect) {
    t.Errorf("Sort(%s) = %v; want %v", s, x, expect)
  }
}

func TestSortFold(t *testing.T) {
  x := []string{"яблоко", "ананас", "Дыня", "посёлок", "поселок"}
  s := fmt.Sprintf("%v", x)
  expect := []string{"ананас", "Дыня", "поселок", "посёлок", "яблоко"}
  SortFold(x)
  if !reflect.DeepEqual(x, expect) {
    t.Errorf("SortFold(%s) = %v; want %v", s, x, expect)
  }
}

