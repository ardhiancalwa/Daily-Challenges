package kata

import "fmt"

func PrinterError(s string) string {
	errorCount := 0
	for _, ch := range s {
		if ch < 'a' || ch > 'm' {
			errorCount++
		}
	}
	return fmt.Sprintf("%d/%d", errorCount, len(s))
}
