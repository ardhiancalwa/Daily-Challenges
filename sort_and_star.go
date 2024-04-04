package kata

import (
  "sort"
  "strings"
)

func TwoSort(arr []string) string {
  sort.Strings(arr)
  firstWord := arr[0]
  result := strings.Join(strings.Split(firstWord, ""), "***")
  return result
}