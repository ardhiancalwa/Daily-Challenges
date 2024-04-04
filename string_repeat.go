package kata

func RepeatStr(repetitions int, value string) string {
  var result string
  for i := 1; i <= repetitions; i++ {
    result += value
  }
  return result
}

// package kata

// import "strings"

// func RepeatStr(repititions int, value string) string {
//   return strings.Repeat(value, repititions)
// }