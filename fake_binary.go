package kata

func FakeBin(x string) string {
      var biner string
      // your code here
      for i := 0; i < len(x);i++ {
        if x[i] < '5' {
          biner += "0"
        } else {
          biner += "1"
        }
      }
      return biner
}
