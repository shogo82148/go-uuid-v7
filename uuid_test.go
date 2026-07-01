package gouuidv7

import (
	"testing"
)

func BenchmarkNewV7(b *testing.B) {
	for b.Loop() {
		NewV7()
	}
}
