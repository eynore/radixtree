package radixtree

import "testing"

func TestLookup(t *testing.T) {
	tree := New()

	tests := []struct {
		key   string
		value string
	}{
		{"tony", "1"},
		{"tonyx", "2"},
		{"tonyxx", "3"},
		{"tonyxy", "4"},
		{"to", "5"},
		{"tox", "6"},
		{"toy", "7"},
		{"xoy", "8"},
		{"abc", "9"},
		{"abd", "10"},
	}
	for _, test := range tests {
		tree.Insert(test.key, test.value)
	}

	for _, test := range tests {
		if got, ok := tree.Lookup(test.key); !ok || got.(string) != test.value {
			t.Errorf("Lookup(%q) = %v, %v", test.key, got, ok)
		}
	}
	others := []string{
		"ab",
		"xxx",
	}
	for _, test := range others {
		if got, ok := tree.Lookup(test); ok || got != nil {
			t.Errorf("Lookup(%q) = %v, %v", test, got, ok)
		}
	}
}
