package kata

func Hero(bullets, dragons int) bool {
  // your code
  if dragons*2 <= bullets {
    return true
  }
  return false
}