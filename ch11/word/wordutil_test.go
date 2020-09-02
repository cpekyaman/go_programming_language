package word

import (
	"math/rand"
	"testing"
	"time"
)

func randomPalindrome(r *rand.Rand) string {
	n := r.Intn(25)
	runes := make([]rune, n)

	for i := 0; i < (n+1)/2; i++ {
		r := rune(r.Intn(0x1000))
		runes[i] = r
		runes[n-i-1] = r
	}

	return string(runes)
}

func TestRandomPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 100; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestPalindrome(t *testing.T) {
	var testCases = []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}

	for _, test := range testCases {
		if got := IsPalindrome(test.input); got != test.expected {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
