package kata

func Solution(word string) string {
  hasil := ""
  for i := len(word)-1; i >= 0; i-- {
    hasil += string(word[i])
  }
  return hasil
}