package kata

import (
	"strconv"
)

func StringToNumber(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}
