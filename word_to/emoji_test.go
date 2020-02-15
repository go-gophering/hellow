package word_to

import "testing"

func TestEmoji(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello world", "👋🌎"},
		{"hello world", "👋🌎"},
		{"", ""},
	}
	for _, c := range cases {
		got := Emoji(c.in)
		if got != c.want {
			t.Errorf("Emoji(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
