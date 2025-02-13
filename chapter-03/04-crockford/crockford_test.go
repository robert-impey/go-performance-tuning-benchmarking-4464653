package crockford

import (
	"runtime"
	"testing"
)

func TestNewID(t *testing.T) {
	t.Logf("NewID() = %s", NewID())
}

func TestEncode(t *testing.T) {
	tests := map[string]struct {
		input    []byte
		expected string
	}{
		"empty string": {},
		"simple": {
			input:    []byte("The quick brown fox jumps over the lazy dog."),
			expected: "AHM6A83HENMP6TS0C9S6YXVE41K6YY10D9TPTW3K41QQCSBJ41T6GS90DHGQMY90CHQPEBG=",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			testInClosure := test
			t.Parallel()
			if got, expected := Encode(testInClosure.input), testInClosure.expected; got != expected {
				t.Fatalf("Expected Encode(%q) to return %q; got %q", testInClosure.input, expected, got)
			}
		})
	}
}

func BenchmarkNewID(b *testing.B) {
	b.ResetTimer()
	var id string
	for i := 0; i < b.N; i++ {
		id = NewID()
	}
	runtime.KeepAlive(&id)
}
