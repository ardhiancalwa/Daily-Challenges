package kata

func CheckForFactor(base int, factor int) bool {
	// your code here
	if base%factor == 0 {
		return true
	} else {
		return false
	}
}
